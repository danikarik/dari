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

func testDescriptions(t *testing.T) {
	t.Parallel()

	query := Descriptions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDescriptionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
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

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDescriptionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Descriptions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDescriptionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DescriptionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDescriptionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DescriptionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Description exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DescriptionExists to return true, but got false.")
	}
}

func testDescriptionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	descriptionFound, err := FindDescription(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if descriptionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDescriptionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Descriptions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDescriptionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Descriptions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDescriptionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	descriptionOne := &Description{}
	descriptionTwo := &Description{}
	if err = randomize.Struct(seed, descriptionOne, descriptionDBTypes, false, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}
	if err = randomize.Struct(seed, descriptionTwo, descriptionDBTypes, false, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = descriptionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = descriptionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Descriptions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDescriptionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	descriptionOne := &Description{}
	descriptionTwo := &Description{}
	if err = randomize.Struct(seed, descriptionOne, descriptionDBTypes, false, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}
	if err = randomize.Struct(seed, descriptionTwo, descriptionDBTypes, false, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = descriptionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = descriptionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func descriptionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func descriptionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Description) error {
	*o = Description{}
	return nil
}

func testDescriptionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Description{}
	o := &Description{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, descriptionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Description object: %s", err)
	}

	AddDescriptionHook(boil.BeforeInsertHook, descriptionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	descriptionBeforeInsertHooks = []DescriptionHook{}

	AddDescriptionHook(boil.AfterInsertHook, descriptionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	descriptionAfterInsertHooks = []DescriptionHook{}

	AddDescriptionHook(boil.AfterSelectHook, descriptionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	descriptionAfterSelectHooks = []DescriptionHook{}

	AddDescriptionHook(boil.BeforeUpdateHook, descriptionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	descriptionBeforeUpdateHooks = []DescriptionHook{}

	AddDescriptionHook(boil.AfterUpdateHook, descriptionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	descriptionAfterUpdateHooks = []DescriptionHook{}

	AddDescriptionHook(boil.BeforeDeleteHook, descriptionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	descriptionBeforeDeleteHooks = []DescriptionHook{}

	AddDescriptionHook(boil.AfterDeleteHook, descriptionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	descriptionAfterDeleteHooks = []DescriptionHook{}

	AddDescriptionHook(boil.BeforeUpsertHook, descriptionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	descriptionBeforeUpsertHooks = []DescriptionHook{}

	AddDescriptionHook(boil.AfterUpsertHook, descriptionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	descriptionAfterUpsertHooks = []DescriptionHook{}
}

func testDescriptionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDescriptionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(descriptionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDescriptionToOneProductUsingProduct(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Description
	var foreign Product

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, descriptionDBTypes, false, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
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

	slice := DescriptionSlice{&local}
	if err = local.L.LoadProduct(ctx, tx, false, (*[]*Description)(&slice), nil); err != nil {
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

func testDescriptionToOneSetOpProductUsingProduct(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Description
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, descriptionDBTypes, false, strmangle.SetComplement(descriptionPrimaryKeyColumns, descriptionColumnsWithoutDefault)...); err != nil {
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

		if x.R.Descriptions[0] != &a {
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

func testDescriptionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
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

func testDescriptionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DescriptionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDescriptionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Descriptions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	descriptionDBTypes = map[string]string{`CreatedAt`: `timestamp`, `ID`: `int`, `ProductID`: `int`, `Text`: `longtext`, `UpdatedAt`: `timestamp`}
	_                  = bytes.MinRead
)

func testDescriptionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(descriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(descriptionColumns) == len(descriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDescriptionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(descriptionColumns) == len(descriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Description{}
	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, descriptionDBTypes, true, descriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(descriptionColumns, descriptionPrimaryKeyColumns) {
		fields = descriptionColumns
	} else {
		fields = strmangle.SetComplement(
			descriptionColumns,
			descriptionPrimaryKeyColumns,
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

	slice := DescriptionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDescriptionsUpsert(t *testing.T) {
	t.Parallel()

	if len(descriptionColumns) == len(descriptionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDescriptionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Description{}
	if err = randomize.Struct(seed, &o, descriptionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Description: %s", err)
	}

	count, err := Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, descriptionDBTypes, false, descriptionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Description struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Description: %s", err)
	}

	count, err = Descriptions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
