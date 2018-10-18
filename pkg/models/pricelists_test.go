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

func testPricelists(t *testing.T) {
	t.Parallel()

	query := Pricelists()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testPricelistsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
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

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPricelistsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Pricelists().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPricelistsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PricelistSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testPricelistsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := PricelistExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Pricelist exists: %s", err)
	}
	if !e {
		t.Errorf("Expected PricelistExists to return true, but got false.")
	}
}

func testPricelistsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	pricelistFound, err := FindPricelist(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if pricelistFound == nil {
		t.Error("want a record, got nil")
	}
}

func testPricelistsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Pricelists().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testPricelistsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Pricelists().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testPricelistsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	pricelistOne := &Pricelist{}
	pricelistTwo := &Pricelist{}
	if err = randomize.Struct(seed, pricelistOne, pricelistDBTypes, false, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}
	if err = randomize.Struct(seed, pricelistTwo, pricelistDBTypes, false, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = pricelistOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = pricelistTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Pricelists().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testPricelistsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	pricelistOne := &Pricelist{}
	pricelistTwo := &Pricelist{}
	if err = randomize.Struct(seed, pricelistOne, pricelistDBTypes, false, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}
	if err = randomize.Struct(seed, pricelistTwo, pricelistDBTypes, false, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = pricelistOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = pricelistTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func pricelistBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func pricelistAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Pricelist) error {
	*o = Pricelist{}
	return nil
}

func testPricelistsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Pricelist{}
	o := &Pricelist{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, pricelistDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Pricelist object: %s", err)
	}

	AddPricelistHook(boil.BeforeInsertHook, pricelistBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	pricelistBeforeInsertHooks = []PricelistHook{}

	AddPricelistHook(boil.AfterInsertHook, pricelistAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	pricelistAfterInsertHooks = []PricelistHook{}

	AddPricelistHook(boil.AfterSelectHook, pricelistAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	pricelistAfterSelectHooks = []PricelistHook{}

	AddPricelistHook(boil.BeforeUpdateHook, pricelistBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	pricelistBeforeUpdateHooks = []PricelistHook{}

	AddPricelistHook(boil.AfterUpdateHook, pricelistAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	pricelistAfterUpdateHooks = []PricelistHook{}

	AddPricelistHook(boil.BeforeDeleteHook, pricelistBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	pricelistBeforeDeleteHooks = []PricelistHook{}

	AddPricelistHook(boil.AfterDeleteHook, pricelistAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	pricelistAfterDeleteHooks = []PricelistHook{}

	AddPricelistHook(boil.BeforeUpsertHook, pricelistBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	pricelistBeforeUpsertHooks = []PricelistHook{}

	AddPricelistHook(boil.AfterUpsertHook, pricelistAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	pricelistAfterUpsertHooks = []PricelistHook{}
}

func testPricelistsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPricelistsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(pricelistColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testPricelistToManyProducts(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pricelist
	var b, c Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
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

	b.PricelistID = a.ID
	c.PricelistID = a.ID

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
		if v.PricelistID == b.PricelistID {
			bFound = true
		}
		if v.PricelistID == c.PricelistID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := PricelistSlice{&a}
	if err = a.L.LoadProducts(ctx, tx, false, (*[]*Pricelist)(&slice), nil); err != nil {
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

func testPricelistToManyAddOpProducts(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pricelist
	var b, c, d, e Product

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pricelistDBTypes, false, strmangle.SetComplement(pricelistPrimaryKeyColumns, pricelistColumnsWithoutDefault)...); err != nil {
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

		if a.ID != first.PricelistID {
			t.Error("foreign key was wrong value", a.ID, first.PricelistID)
		}
		if a.ID != second.PricelistID {
			t.Error("foreign key was wrong value", a.ID, second.PricelistID)
		}

		if first.R.Pricelist != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Pricelist != &a {
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
func testPricelistToOneManufacturerUsingManufacturer(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Pricelist
	var foreign Manufacturer

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pricelistDBTypes, false, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, manufacturerDBTypes, false, manufacturerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Manufacturer struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ManufacturerID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Manufacturer().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PricelistSlice{&local}
	if err = local.L.LoadManufacturer(ctx, tx, false, (*[]*Pricelist)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Manufacturer == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Manufacturer = nil
	if err = local.L.LoadManufacturer(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Manufacturer == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPricelistToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Pricelist
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, pricelistDBTypes, false, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := PricelistSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Pricelist)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testPricelistToOneSetOpManufacturerUsingManufacturer(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pricelist
	var b, c Manufacturer

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pricelistDBTypes, false, strmangle.SetComplement(pricelistPrimaryKeyColumns, pricelistColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, manufacturerDBTypes, false, strmangle.SetComplement(manufacturerPrimaryKeyColumns, manufacturerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, manufacturerDBTypes, false, strmangle.SetComplement(manufacturerPrimaryKeyColumns, manufacturerColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Manufacturer{&b, &c} {
		err = a.SetManufacturer(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Manufacturer != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Pricelists[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ManufacturerID != x.ID {
			t.Error("foreign key was wrong value", a.ManufacturerID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ManufacturerID))
		reflect.Indirect(reflect.ValueOf(&a.ManufacturerID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ManufacturerID != x.ID {
			t.Error("foreign key was wrong value", a.ManufacturerID, x.ID)
		}
	}
}
func testPricelistToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Pricelist
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, pricelistDBTypes, false, strmangle.SetComplement(pricelistPrimaryKeyColumns, pricelistColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Pricelists[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testPricelistsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
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

func testPricelistsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := PricelistSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testPricelistsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Pricelists().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	pricelistDBTypes = map[string]string{`CreatedAt`: `timestamp`, `ID`: `int`, `IsActive`: `tinyint`, `IsArchive`: `tinyint`, `ListName`: `varchar`, `ManufacturerID`: `int`, `UpdatedAt`: `timestamp`, `UserID`: `int`, `Year`: `int`}
	_                = bytes.MinRead
)

func testPricelistsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(pricelistPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(pricelistColumns) == len(pricelistPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testPricelistsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(pricelistColumns) == len(pricelistPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Pricelist{}
	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, pricelistDBTypes, true, pricelistPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(pricelistColumns, pricelistPrimaryKeyColumns) {
		fields = pricelistColumns
	} else {
		fields = strmangle.SetComplement(
			pricelistColumns,
			pricelistPrimaryKeyColumns,
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

	slice := PricelistSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testPricelistsUpsert(t *testing.T) {
	t.Parallel()

	if len(pricelistColumns) == len(pricelistPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLPricelistUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Pricelist{}
	if err = randomize.Struct(seed, &o, pricelistDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Pricelist: %s", err)
	}

	count, err := Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, pricelistDBTypes, false, pricelistPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Pricelist struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Pricelist: %s", err)
	}

	count, err = Pricelists().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
