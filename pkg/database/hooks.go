package database

import (
	"bitbucket.org/kit-systems/dari/pkg/dict"
	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// AddRegistryHooks append hook to registry model.
func (db *DB) AddRegistryHooks(p boil.HookPoint, h models.RegistryHook) error {
	models.AddRegistryHook(p, h)
	return nil
}

// InsertFieldStat inserts new field changes.
func (db *DB) InsertFieldStat(fs *models.RegistryFieldStat, r *models.Registry) error {
	err := fs.Insert(db.ctx, db.conn, boil.Infer())
	if err != nil {
		return err
	}
	if r != nil && r.RegistryStatusID != dict.AlreadyVisited {
		_, err = r.Update(db.ctx, db.conn, boil.Whitelist("registry_status_id"))
		if err != nil {
			return err
		}
	}
	return nil
}
