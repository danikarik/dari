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

func testRegistryRecommendations(t *testing.T) {
	t.Parallel()

	query := RegistryRecommendations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRegistryRecommendationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
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

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryRecommendationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RegistryRecommendations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryRecommendationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegistryRecommendationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryRecommendationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RegistryRecommendationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RegistryRecommendation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RegistryRecommendationExists to return true, but got false.")
	}
}

func testRegistryRecommendationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	registryRecommendationFound, err := FindRegistryRecommendation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if registryRecommendationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRegistryRecommendationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RegistryRecommendations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRegistryRecommendationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RegistryRecommendations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRegistryRecommendationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	registryRecommendationOne := &RegistryRecommendation{}
	registryRecommendationTwo := &RegistryRecommendation{}
	if err = randomize.Struct(seed, registryRecommendationOne, registryRecommendationDBTypes, false, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}
	if err = randomize.Struct(seed, registryRecommendationTwo, registryRecommendationDBTypes, false, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registryRecommendationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registryRecommendationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegistryRecommendations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRegistryRecommendationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	registryRecommendationOne := &RegistryRecommendation{}
	registryRecommendationTwo := &RegistryRecommendation{}
	if err = randomize.Struct(seed, registryRecommendationOne, registryRecommendationDBTypes, false, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}
	if err = randomize.Struct(seed, registryRecommendationTwo, registryRecommendationDBTypes, false, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registryRecommendationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registryRecommendationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func registryRecommendationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func registryRecommendationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryRecommendation) error {
	*o = RegistryRecommendation{}
	return nil
}

func testRegistryRecommendationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RegistryRecommendation{}
	o := &RegistryRecommendation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation object: %s", err)
	}

	AddRegistryRecommendationHook(boil.BeforeInsertHook, registryRecommendationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	registryRecommendationBeforeInsertHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.AfterInsertHook, registryRecommendationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	registryRecommendationAfterInsertHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.AfterSelectHook, registryRecommendationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	registryRecommendationAfterSelectHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.BeforeUpdateHook, registryRecommendationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	registryRecommendationBeforeUpdateHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.AfterUpdateHook, registryRecommendationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	registryRecommendationAfterUpdateHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.BeforeDeleteHook, registryRecommendationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	registryRecommendationBeforeDeleteHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.AfterDeleteHook, registryRecommendationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	registryRecommendationAfterDeleteHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.BeforeUpsertHook, registryRecommendationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	registryRecommendationBeforeUpsertHooks = []RegistryRecommendationHook{}

	AddRegistryRecommendationHook(boil.AfterUpsertHook, registryRecommendationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	registryRecommendationAfterUpsertHooks = []RegistryRecommendationHook{}
}

func testRegistryRecommendationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegistryRecommendationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(registryRecommendationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegistryRecommendationToOneProductUsingProduct(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RegistryRecommendation
	var foreign Product

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, registryRecommendationDBTypes, false, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Product struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ProductID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Product().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RegistryRecommendationSlice{&local}
	if err = local.L.LoadProduct(ctx, tx, false, (*[]*RegistryRecommendation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Product = nil
	if err = local.L.LoadProduct(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Product == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRegistryRecommendationToOneRegistryUsingRegistry(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RegistryRecommendation
	var foreign Registry

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, registryRecommendationDBTypes, false, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
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

	slice := RegistryRecommendationSlice{&local}
	if err = local.L.LoadRegistry(ctx, tx, false, (*[]*RegistryRecommendation)(&slice), nil); err != nil {
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

func testRegistryRecommendationToOneSetOpProductUsingProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RegistryRecommendation
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, registryRecommendationDBTypes, false, strmangle.SetComplement(registryRecommendationPrimaryKeyColumns, registryRecommendationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Product{&b, &c} {
		err = a.SetProduct(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Product != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RegistryRecommendations[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ProductID != x.ID {
			t.Error("foreign key was wrong value", a.ProductID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ProductID))
		reflect.Indirect(reflect.ValueOf(&a.ProductID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ProductID != x.ID {
			t.Error("foreign key was wrong value", a.ProductID, x.ID)
		}
	}
}
func testRegistryRecommendationToOneSetOpRegistryUsingRegistry(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RegistryRecommendation
	var b, c Registry

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, registryRecommendationDBTypes, false, strmangle.SetComplement(registryRecommendationPrimaryKeyColumns, registryRecommendationColumnsWithoutDefault)...); err != nil {
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

		if x.R.RegistryRecommendations[0] != &a {
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

func testRegistryRecommendationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
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

func testRegistryRecommendationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegistryRecommendationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRegistryRecommendationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegistryRecommendations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	registryRecommendationDBTypes = map[string]string{`ID`: `int`, `IsApplied`: `tinyint`, `ProductID`: `int`, `RegistryID`: `int`}
	_                             = bytes.MinRead
)

func testRegistryRecommendationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(registryRecommendationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(registryRecommendationColumns) == len(registryRecommendationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRegistryRecommendationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(registryRecommendationColumns) == len(registryRecommendationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegistryRecommendation{}
	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registryRecommendationDBTypes, true, registryRecommendationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(registryRecommendationColumns, registryRecommendationPrimaryKeyColumns) {
		fields = registryRecommendationColumns
	} else {
		fields = strmangle.SetComplement(
			registryRecommendationColumns,
			registryRecommendationPrimaryKeyColumns,
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

	slice := RegistryRecommendationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRegistryRecommendationsUpsert(t *testing.T) {
	t.Parallel()

	if len(registryRecommendationColumns) == len(registryRecommendationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRegistryRecommendationUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RegistryRecommendation{}
	if err = randomize.Struct(seed, &o, registryRecommendationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegistryRecommendation: %s", err)
	}

	count, err := RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, registryRecommendationDBTypes, false, registryRecommendationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryRecommendation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegistryRecommendation: %s", err)
	}

	count, err = RegistryRecommendations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
