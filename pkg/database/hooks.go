package database

import (
	"context"

	"bitbucket.org/kit-systems/dari/pkg/dict"
	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func beforeRegistryUpsert(ctx context.Context, exec boil.ContextExecutor, r *models.Registry) error {
	_, err := models.FindRegistry(ctx, exec, r.ID)
	if err != nil {
		r.RegistryStatusID = uint(dict.Inserted)
	} else {
		r.RegistryStatusID = uint(dict.Updated)
	}
	return nil
}

// AddRegistryHooks append hook to registry model.
func (db *DB) AddRegistryHooks() error {
	models.AddRegistryHook(boil.BeforeUpsertHook, beforeRegistryUpsert)
	return nil
}
