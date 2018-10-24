package database

import (
	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

// FindProducts gets all products without registry id.
func (db *DB) FindProducts() (models.ProductSlice, error) {
	return models.Products(qm.Where("registry_id IS NULL")).All(db.ctx, db.conn)
}

// FindRegistry searches for registry with similar title, model, etc.
func (db *DB) FindRegistry(prod *models.Product) error {
	// regs, err := models.Registries(qm.Where("title like ?")).All(db.ctx, db.conn)
	return nil
}
