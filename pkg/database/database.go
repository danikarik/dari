package database

import (
	"context"
	"database/sql"
	"time"

	"bitbucket.org/kit-systems/dari/pkg/dict"
	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/hashicorp/go-multierror"
	"github.com/oklog/run"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// DB is used for database operations.
type DB struct {
	ctx  context.Context
	conn *sql.DB
}

// New returns new DB container.
func New(ctx context.Context, conn *sql.DB) *DB {
	return &DB{
		ctx:  ctx,
		conn: conn,
	}
}

// GetRegistries returns all registry records.
func (db *DB) GetRegistries() (models.RegistrySlice, error) {
	return models.Registries().All(db.ctx, db.conn)
}

// GetRegistryManufacturers returns all manufacturers of given registry record.
func (db *DB) GetRegistryManufacturers(r *models.Registry) (models.RegistryManufacturerSlice, error) {
	return r.RegistryManufacturers().All(db.ctx, db.conn)
}

// GetRegistryBuild returns all builds of given registry record.
func (db *DB) GetRegistryBuild(r *models.Registry) (models.RegistryBuildSlice, error) {
	return r.RegistryBuilds().All(db.ctx, db.conn)
}

// MarkAsProcessing updates status field for all records
// except those are marked as parsing failed.
func (db *DB) MarkAsProcessing(regs models.RegistrySlice) error {
	tx, err := db.conn.BeginTx(db.ctx, nil)
	if err != nil {
		return err
	}
	_, err = models.Registries().UpdateAll(db.ctx, db.conn, models.M{
		"registry_status_id": uint(dict.Processing),
		"updated_at":         time.Now(),
	})
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

// MarkAsDeleted updates status field for all records.
func (db *DB) MarkAsDeleted(regs models.RegistrySlice) error {
	tx, err := db.conn.BeginTx(db.ctx, nil)
	if err != nil {
		return err
	}
	_, err = models.Registries(qm.Where("registry_status_id = ?", uint(dict.Processing))).UpdateAll(db.ctx, db.conn, models.M{
		"registry_status_id": uint(dict.Deleted),
		"updated_at":         time.Now(),
	})
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}

// AddRegistry inserts registry into database.
func (db *DB) AddRegistry(r *models.Registry) error {
	tx, err := db.conn.BeginTx(db.ctx, nil)
	if err != nil {
		return err
	}
	err = r.Upsert(db.ctx, db.conn, boil.Blacklist("created_at"), boil.Infer())
	if err != nil {
		return err
	}
	var g run.Group
	{
		g.Add(func() error {
			_, err := r.RegistryManufacturers().DeleteAll(db.ctx, db.conn)
			if err != nil {
				return err
			}
			for _, m := range r.R.RegistryManufacturers {
				m.RegistryID = r.ID
				terr := m.Insert(db.ctx, db.conn, boil.Infer())
				if terr != nil {
					err = multierror.Append(err, terr)
					break
				}
			}
			return err
		}, func(error) {})
		g.Add(func() error {
			_, err := r.RegistryBuilds().DeleteAll(db.ctx, db.conn)
			if err != nil {
				return err
			}
			for _, b := range r.R.RegistryBuilds {
				b.RegistryID = r.ID
				terr := b.Insert(db.ctx, db.conn, boil.Infer())
				if terr != nil {
					err = multierror.Append(err, terr)
					break
				}
			}
			return err
		}, func(error) {})
	}
	if err = g.Run(); err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
