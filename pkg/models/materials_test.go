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

func testMaterials(t *testing.T) {
	t.Parallel()

	query := Materials()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMaterialsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
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

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMaterialsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Materials().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMaterialsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MaterialSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMaterialsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := MaterialExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Material exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MaterialExists to return true, but got false.")
	}
}

func testMaterialsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	materialFound, err := FindMaterial(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if materialFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMaterialsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Materials().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testMaterialsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Materials().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMaterialsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	materialOne := &Material{}
	materialTwo := &Material{}
	if err = randomize.Struct(seed, materialOne, materialDBTypes, false, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}
	if err = randomize.Struct(seed, materialTwo, materialDBTypes, false, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = materialOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = materialTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Materials().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMaterialsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	materialOne := &Material{}
	materialTwo := &Material{}
	if err = randomize.Struct(seed, materialOne, materialDBTypes, false, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}
	if err = randomize.Struct(seed, materialTwo, materialDBTypes, false, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = materialOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = materialTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func materialBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func materialAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Material) error {
	*o = Material{}
	return nil
}

func testMaterialsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Material{}
	o := &Material{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, materialDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Material object: %s", err)
	}

	AddMaterialHook(boil.BeforeInsertHook, materialBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	materialBeforeInsertHooks = []MaterialHook{}

	AddMaterialHook(boil.AfterInsertHook, materialAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	materialAfterInsertHooks = []MaterialHook{}

	AddMaterialHook(boil.AfterSelectHook, materialAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	materialAfterSelectHooks = []MaterialHook{}

	AddMaterialHook(boil.BeforeUpdateHook, materialBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	materialBeforeUpdateHooks = []MaterialHook{}

	AddMaterialHook(boil.AfterUpdateHook, materialAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	materialAfterUpdateHooks = []MaterialHook{}

	AddMaterialHook(boil.BeforeDeleteHook, materialBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	materialBeforeDeleteHooks = []MaterialHook{}

	AddMaterialHook(boil.AfterDeleteHook, materialAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	materialAfterDeleteHooks = []MaterialHook{}

	AddMaterialHook(boil.BeforeUpsertHook, materialBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	materialBeforeUpsertHooks = []MaterialHook{}

	AddMaterialHook(boil.AfterUpsertHook, materialAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	materialAfterUpsertHooks = []MaterialHook{}
}

func testMaterialsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMaterialsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(materialColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMaterialToManyProducts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Material
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, productDBTypes, false, productColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&b.MaterialID, a.ID)
	queries.Assign(&c.MaterialID, a.ID)
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	product, err := a.Products().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range product {
		if queries.Equal(v.MaterialID, b.MaterialID) {
			bFound = true
		}
		if queries.Equal(v.MaterialID, c.MaterialID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := MaterialSlice{&a}
	if err = a.L.LoadProducts(ctx, tx, false, (*[]*Material)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Products); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.Products = nil
	if err = a.L.LoadProducts(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.Products); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", product)
	}
}

func testMaterialToManyAddOpProducts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Material
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, materialDBTypes, false, strmangle.SetComplement(materialPrimaryKeyColumns, materialColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Product{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Product{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddProducts(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if !queries.Equal(a.ID, first.MaterialID) {
			t.Error("foreign key was wrong value", a.ID, first.MaterialID)
		}
		if !queries.Equal(a.ID, second.MaterialID) {
			t.Error("foreign key was wrong value", a.ID, second.MaterialID)
		}

		if first.R.Material != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Material != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.Products[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.Products[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.Products().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testMaterialToManySetOpProducts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Material
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, materialDBTypes, false, strmangle.SetComplement(materialPrimaryKeyColumns, materialColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Product{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetProducts(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetProducts(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MaterialID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MaterialID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.MaterialID) {
		t.Error("foreign key was wrong value", a.ID, d.MaterialID)
	}
	if !queries.Equal(a.ID, e.MaterialID) {
		t.Error("foreign key was wrong value", a.ID, e.MaterialID)
	}

	if b.R.Material != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Material != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Material != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.Material != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.Products[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.Products[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testMaterialToManyRemoveOpProducts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Material
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, materialDBTypes, false, strmangle.SetComplement(materialPrimaryKeyColumns, materialColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Product{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, productDBTypes, false, strmangle.SetComplement(productPrimaryKeyColumns, productColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.AddProducts(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveProducts(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.Products().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.MaterialID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.MaterialID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.Material != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.Material != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.Material != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.Material != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.Products) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.Products[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.Products[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testMaterialsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
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

func testMaterialsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := MaterialSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testMaterialsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Materials().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	materialDBTypes = map[string]string{`ID`: `int`, `Title`: `varchar`}
	_               = bytes.MinRead
)

func testMaterialsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(materialPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(materialColumns) == len(materialPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, materialDBTypes, true, materialPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testMaterialsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(materialColumns) == len(materialPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Material{}
	if err = randomize.Struct(seed, o, materialDBTypes, true, materialColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, materialDBTypes, true, materialPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(materialColumns, materialPrimaryKeyColumns) {
		fields = materialColumns
	} else {
		fields = strmangle.SetComplement(
			materialColumns,
			materialPrimaryKeyColumns,
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

	slice := MaterialSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testMaterialsUpsert(t *testing.T) {
	t.Parallel()

	if len(materialColumns) == len(materialPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLMaterialUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Material{}
	if err = randomize.Struct(seed, &o, materialDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Material: %s", err)
	}

	count, err := Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, materialDBTypes, false, materialPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Material struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Material: %s", err)
	}

	count, err = Materials().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}