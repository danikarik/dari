package database

import (
	"context"
	"database/sql"

	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/hashicorp/go-multierror"
	"github.com/volatiletech/sqlboiler/boil"
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

// AddRegistry inserts registry into database.
func (db *DB) AddRegistry(r *models.Registry) error {
	tx, err := db.conn.BeginTx(db.ctx, nil)
	if err != nil {
		return err
	}
	err = r.Insert(db.ctx, db.conn, boil.Infer())
	if err != nil {
		return err
	}
	for _, m := range r.R.RegistryManufacturers {
		m.RegistryID = r.ID
		terr := m.Insert(db.ctx, db.conn, boil.Infer())
		if terr != nil {
			err = multierror.Append(err, terr)
		}
	}
	for _, b := range r.R.RegistryBuilds {
		b.RegistryID = r.ID
		terr := b.Insert(db.ctx, db.conn, boil.Infer())
		if terr != nil {
			err = multierror.Append(err, terr)
		}
	}
	if err != nil {
		return tx.Rollback()
	}
	return tx.Commit()
}
