package database

import (
	"context"
	"fmt"

	"bitbucket.org/kit-systems/dari/pkg/dict"
	"bitbucket.org/kit-systems/dari/pkg/models"
	"github.com/volatiletech/sqlboiler/boil"
)

func beforeRegistryInsert(ctx context.Context, exec boil.ContextExecutor, r *models.Registry) error {
	fmt.Println("insert")
	r.RegistryStatusID = uint(dict.Inserted)
	return nil
}

func beforeRegistryUpdate(ctx context.Context, exec boil.ContextExecutor, r *models.Registry) error {
	fmt.Println("update")
	r.RegistryStatusID = uint(dict.Updated)
	return nil
}

// AddRegistryHooks append hook to registry model.
func (db *DB) AddRegistryHooks() error {
	models.AddRegistryHook(boil.BeforeInsertHook, beforeRegistryInsert)
	models.AddRegistryHook(boil.BeforeUpdateHook, beforeRegistryUpdate)
	return nil
}
