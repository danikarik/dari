package database

import (
	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/volatiletech/sqlboiler/boil"
)

// StartProcess launch a parsing process.
func (db *DB) StartProcess(j *models.RegistryJournal) error {
	return j.Insert(db.ctx, db.conn, boil.Infer())
}

// UpdateCounter updates database related counters.
func (db *DB) UpdateCounter(j *models.RegistryJournal) error {
	_, err := j.Update(db.ctx, db.conn, boil.Blacklist(
		"started_at",
		"finished_at",
	))
	return err
}

// FinishProcess summarize parsing process.
func (db *DB) FinishProcess(j *models.RegistryJournal) error {
	_, err := j.Update(db.ctx, db.conn, boil.Infer())
	return err
}
