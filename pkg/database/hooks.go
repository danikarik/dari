package database

import (
	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// AddRegistryHooks append hook to registry model.
func (db *DB) AddRegistryHooks(p boil.HookPoint, h models.RegistryHook) error {
	models.AddRegistryHook(p, h)
	return nil
}

// InsertFieldStat inserts new field changes.
func (db *DB) InsertFieldStat(fs *models.RegistryFieldStat) error {
	return fs.Insert(db.ctx, db.conn, boil.Infer())
}
