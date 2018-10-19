// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRegistryManufacturers(t *testing.T) {
	t.Parallel()

	query := RegistryManufacturers()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRegistryManufacturersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryManufacturersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RegistryManufacturers().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryManufacturersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegistryManufacturerSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryManufacturersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RegistryManufacturerExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RegistryManufacturer exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RegistryManufacturerExists to return true, but got false.")
	}
}

func testRegistryManufacturersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	registryManufacturerFound, err := FindRegistryManufacturer(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if registryManufacturerFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRegistryManufacturersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RegistryManufacturers().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRegistryManufacturersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RegistryManufacturers().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRegistryManufacturersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	registryManufacturerOne := &RegistryManufacturer{}
	registryManufacturerTwo := &RegistryManufacturer{}
	if err = randomize.Struct(seed, registryManufacturerOne, registryManufacturerDBTypes, false, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}
	if err = randomize.Struct(seed, registryManufacturerTwo, registryManufacturerDBTypes, false, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registryManufacturerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registryManufacturerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegistryManufacturers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRegistryManufacturersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	registryManufacturerOne := &RegistryManufacturer{}
	registryManufacturerTwo := &RegistryManufacturer{}
	if err = randomize.Struct(seed, registryManufacturerOne, registryManufacturerDBTypes, false, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}
	if err = randomize.Struct(seed, registryManufacturerTwo, registryManufacturerDBTypes, false, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registryManufacturerOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registryManufacturerTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func registryManufacturerBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func registryManufacturerAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryManufacturer) error {
	*o = RegistryManufacturer{}
	return nil
}

func testRegistryManufacturersHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RegistryManufacturer{}
	o := &RegistryManufacturer{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer object: %s", err)
	}

	AddRegistryManufacturerHook(boil.BeforeInsertHook, registryManufacturerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	registryManufacturerBeforeInsertHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.AfterInsertHook, registryManufacturerAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	registryManufacturerAfterInsertHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.AfterSelectHook, registryManufacturerAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	registryManufacturerAfterSelectHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.BeforeUpdateHook, registryManufacturerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	registryManufacturerBeforeUpdateHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.AfterUpdateHook, registryManufacturerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	registryManufacturerAfterUpdateHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.BeforeDeleteHook, registryManufacturerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	registryManufacturerBeforeDeleteHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.AfterDeleteHook, registryManufacturerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	registryManufacturerAfterDeleteHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.BeforeUpsertHook, registryManufacturerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	registryManufacturerBeforeUpsertHooks = []RegistryManufacturerHook{}

	AddRegistryManufacturerHook(boil.AfterUpsertHook, registryManufacturerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	registryManufacturerAfterUpsertHooks = []RegistryManufacturerHook{}
}

func testRegistryManufacturersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegistryManufacturersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(registryManufacturerColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegistryManufacturerToOneRegistryUsingRegistry(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RegistryManufacturer
	var foreign Registry

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, registryManufacturerDBTypes, false, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, registryDBTypes, false, registryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Registry struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RegistryID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Registry().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RegistryManufacturerSlice{&local}
	if err = local.L.LoadRegistry(ctx, tx, false, (*[]*RegistryManufacturer)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Registry == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Registry = nil
	if err = local.L.LoadRegistry(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Registry == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRegistryManufacturerToOneSetOpRegistryUsingRegistry(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RegistryManufacturer
	var b, c Registry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, registryManufacturerDBTypes, false, strmangle.SetComplement(registryManufacturerPrimaryKeyColumns, registryManufacturerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, registryDBTypes, false, strmangle.SetComplement(registryPrimaryKeyColumns, registryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, registryDBTypes, false, strmangle.SetComplement(registryPrimaryKeyColumns, registryColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Registry{&b, &c} {
		err = a.SetRegistry(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Registry != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RegistryManufacturers[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RegistryID != x.ID {
			t.Error("foreign key was wrong value", a.RegistryID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RegistryID))
		reflect.Indirect(reflect.ValueOf(&a.RegistryID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RegistryID != x.ID {
			t.Error("foreign key was wrong value", a.RegistryID, x.ID)
		}
	}
}

func testRegistryManufacturersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRegistryManufacturersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegistryManufacturerSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRegistryManufacturersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegistryManufacturers().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	registryManufacturerDBTypes = map[string]string{`Country`: `varchar`, `CreatedAt`: `timestamp`, `ID`: `int`, `RegistryID`: `int`, `Title`: `varchar`, `UpdatedAt`: `timestamp`}
	_                           = bytes.MinRead
)

func testRegistryManufacturersUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(registryManufacturerPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(registryManufacturerColumns) == len(registryManufacturerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRegistryManufacturersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(registryManufacturerColumns) == len(registryManufacturerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegistryManufacturer{}
	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registryManufacturerDBTypes, true, registryManufacturerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(registryManufacturerColumns, registryManufacturerPrimaryKeyColumns) {
		fields = registryManufacturerColumns
	} else {
		fields = strmangle.SetComplement(
			registryManufacturerColumns,
			registryManufacturerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RegistryManufacturerSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRegistryManufacturersUpsert(t *testing.T) {
	t.Parallel()

	if len(registryManufacturerColumns) == len(registryManufacturerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRegistryManufacturerUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RegistryManufacturer{}
	if err = randomize.Struct(seed, &o, registryManufacturerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegistryManufacturer: %s", err)
	}

	count, err := RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, registryManufacturerDBTypes, false, registryManufacturerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryManufacturer struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegistryManufacturer: %s", err)
	}

	count, err = RegistryManufacturers().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}