package database

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/hashicorp/go-multierror"
	"github.com/satori/go.uuid"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// FindProducts gets all products without registry id.
func (db *DB) FindProducts() (models.ProductSlice, error) {
	return models.Products(qm.Where("registry_id IS NULL")).All(db.ctx, db.conn)
}

// FindRegistry searches for registry with similar title, model, etc.
func (db *DB) FindRegistry(prod *models.Product) error {
	var (
		// Транзакция
		tx *sql.Tx
		// Ошибка
		err error
		// Производитель
		mfc *models.Manufacturer
		// Список РУ
		regs models.RegistrySlice
		//
	)
	tx, err = db.conn.BeginTx(db.ctx, nil)
	if err != nil {
		return err
	}
	// Читаем производителя
	mfc, err = prod.Manufacturer().One(db.ctx, db.conn)
	if err != nil {
		terr := tx.Rollback()
		return multierror.Append(terr, err)
	}
	// Получаем РУ, где производитель сходится по наименованию
	regs, err = models.Registries(
		qm.InnerJoin("registry_manufacturers rm on rm.registry_id = registries.id"),
		qm.InnerJoin("registry_builds rb on rb.registry_id = registries.id"),
		qm.Where("lower(rm.title) like ?", `%`+strings.ToLower(mfc.BrandName)+`%`),
		qm.Where(
			"lower(registries.title) like ? or lower(rb.title) like ?",
			`%`+strings.ToLower(prod.BaseName)+`%`,
			`%`+strings.ToLower(prod.BaseName)+`%`,
		),
	).All(db.ctx, db.conn)
	// Записываем рекомендации
	for _, reg := range regs {
		rem := &models.RegistryRecommendation{
			RegistryID: reg.ID,
			ProductID:  prod.ID,
			IsApplied:  false,
		}
		var ok bool
		ok, err = models.RegistryRecommendations(
			qm.Where("registry_id = ?", reg.ID),
			qm.And("product_id = ?", prod.ID),
		).Exists(db.ctx, db.conn)
		if ok {
			_, err = rem.Update(db.ctx, db.conn, boil.Whitelist("is_applied"))
		} else {
			err = rem.Insert(db.ctx, db.conn, boil.Infer())
		}
		if err != nil {
			terr := tx.Rollback()
			return multierror.Append(terr, err)
		}
	}
	return tx.Commit()
}

// CreateNotifications creates notification of matched registries.
func (db *DB) CreateNotifications(prods models.ProductSlice) error {
	var ids []interface{}
	var rds []interface{}
	for _, p := range prods {
		ids = append(ids, p.ID)
	}
	rems, err := models.RegistryRecommendations(
		qm.Select("distinct(registry_id)"),
		qm.WhereIn("product_id in ?", ids...),
	).All(db.ctx, db.conn)
	if err != nil {
		return err
	}
	if len(rems) > 0 {
		for _, r := range rems {
			rds = append(rds, r.RegistryID)
		}
		var regs models.RegistrySlice
		regs, err = models.Registries(qm.WhereIn("id in ?", rds...), qm.Load("RegistryRecommendations")).All(db.ctx, db.conn)
		if err != nil {
			return err
		}
		for _, r := range regs {
			_, err = models.Notifications(
				qm.Where("type like ?", `%RecommendationFound`),
				qm.And("notifiable_type like ?", `%Registry`),
				qm.And("notifiable_id = ?", r.ID),
			).DeleteAll(db.ctx, db.conn)
			if err != nil {
				continue
			}
			note := models.Notification{
				ID:             uuid.NewV4().String(),
				Type:           `App\Notifications\RecommendationFound`,
				NotifiableType: `App\Registry`,
				NotifiableID:   uint64(r.ID),
				Data:           fmt.Sprintf(`{"count":%d}`, len(r.R.RegistryRecommendations)),
				CreatedAt:      null.TimeFrom(time.Now()),
				UpdatedAt:      null.TimeFrom(time.Now()),
			}
			err = note.Insert(db.ctx, db.conn, boil.Infer())
			if err != nil {
				continue
			}
		}
	}
	return nil
}
