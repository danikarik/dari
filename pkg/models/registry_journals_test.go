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

func testRegistryJournals(t *testing.T) {
	t.Parallel()

	query := RegistryJournals()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRegistryJournalsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
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

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryJournalsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RegistryJournals().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryJournalsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegistryJournalSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRegistryJournalsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RegistryJournalExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RegistryJournal exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RegistryJournalExists to return true, but got false.")
	}
}

func testRegistryJournalsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	registryJournalFound, err := FindRegistryJournal(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if registryJournalFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRegistryJournalsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RegistryJournals().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRegistryJournalsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RegistryJournals().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRegistryJournalsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	registryJournalOne := &RegistryJournal{}
	registryJournalTwo := &RegistryJournal{}
	if err = randomize.Struct(seed, registryJournalOne, registryJournalDBTypes, false, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}
	if err = randomize.Struct(seed, registryJournalTwo, registryJournalDBTypes, false, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registryJournalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registryJournalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegistryJournals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRegistryJournalsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	registryJournalOne := &RegistryJournal{}
	registryJournalTwo := &RegistryJournal{}
	if err = randomize.Struct(seed, registryJournalOne, registryJournalDBTypes, false, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}
	if err = randomize.Struct(seed, registryJournalTwo, registryJournalDBTypes, false, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = registryJournalOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = registryJournalTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func registryJournalBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func registryJournalAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RegistryJournal) error {
	*o = RegistryJournal{}
	return nil
}

func testRegistryJournalsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RegistryJournal{}
	o := &RegistryJournal{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, registryJournalDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegistryJournal object: %s", err)
	}

	AddRegistryJournalHook(boil.BeforeInsertHook, registryJournalBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	registryJournalBeforeInsertHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.AfterInsertHook, registryJournalAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	registryJournalAfterInsertHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.AfterSelectHook, registryJournalAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	registryJournalAfterSelectHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.BeforeUpdateHook, registryJournalBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	registryJournalBeforeUpdateHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.AfterUpdateHook, registryJournalAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	registryJournalAfterUpdateHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.BeforeDeleteHook, registryJournalBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	registryJournalBeforeDeleteHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.AfterDeleteHook, registryJournalAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	registryJournalAfterDeleteHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.BeforeUpsertHook, registryJournalBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	registryJournalBeforeUpsertHooks = []RegistryJournalHook{}

	AddRegistryJournalHook(boil.AfterUpsertHook, registryJournalAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	registryJournalAfterUpsertHooks = []RegistryJournalHook{}
}

func testRegistryJournalsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegistryJournalsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(registryJournalColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRegistryJournalToManyRegistryFieldStats(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RegistryJournal
	var b, c RegistryFieldStat

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, registryFieldStatDBTypes, false, registryFieldStatColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, registryFieldStatDBTypes, false, registryFieldStatColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.RegistryJournalID = a.ID
	c.RegistryJournalID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	registryFieldStat, err := a.RegistryFieldStats().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range registryFieldStat {
		if v.RegistryJournalID == b.RegistryJournalID {
			bFound = true
		}
		if v.RegistryJournalID == c.RegistryJournalID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := RegistryJournalSlice{&a}
	if err = a.L.LoadRegistryFieldStats(ctx, tx, false, (*[]*RegistryJournal)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RegistryFieldStats); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.RegistryFieldStats = nil
	if err = a.L.LoadRegistryFieldStats(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.RegistryFieldStats); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", registryFieldStat)
	}
}

func testRegistryJournalToManyAddOpRegistryFieldStats(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RegistryJournal
	var b, c, d, e RegistryFieldStat

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, registryJournalDBTypes, false, strmangle.SetComplement(registryJournalPrimaryKeyColumns, registryJournalColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*RegistryFieldStat{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, registryFieldStatDBTypes, false, strmangle.SetComplement(registryFieldStatPrimaryKeyColumns, registryFieldStatColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*RegistryFieldStat{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddRegistryFieldStats(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.RegistryJournalID {
			t.Error("foreign key was wrong value", a.ID, first.RegistryJournalID)
		}
		if a.ID != second.RegistryJournalID {
			t.Error("foreign key was wrong value", a.ID, second.RegistryJournalID)
		}

		if first.R.RegistryJournal != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.RegistryJournal != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.RegistryFieldStats[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.RegistryFieldStats[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.RegistryFieldStats().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testRegistryJournalsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
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

func testRegistryJournalsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RegistryJournalSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRegistryJournalsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RegistryJournals().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	registryJournalDBTypes = map[string]string{`DeletedCount`: `int`, `FailedCount`: `int`, `FinishedAt`: `timestamp`, `ID`: `int`, `InsertedCount`: `int`, `StartedAt`: `timestamp`, `TotalCount`: `int`, `UpdatedCount`: `int`}
	_                      = bytes.MinRead
)

func testRegistryJournalsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(registryJournalPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(registryJournalColumns) == len(registryJournalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRegistryJournalsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(registryJournalColumns) == len(registryJournalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RegistryJournal{}
	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, registryJournalDBTypes, true, registryJournalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(registryJournalColumns, registryJournalPrimaryKeyColumns) {
		fields = registryJournalColumns
	} else {
		fields = strmangle.SetComplement(
			registryJournalColumns,
			registryJournalPrimaryKeyColumns,
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

	slice := RegistryJournalSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRegistryJournalsUpsert(t *testing.T) {
	t.Parallel()

	if len(registryJournalColumns) == len(registryJournalPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLRegistryJournalUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RegistryJournal{}
	if err = randomize.Struct(seed, &o, registryJournalDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegistryJournal: %s", err)
	}

	count, err := RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, registryJournalDBTypes, false, registryJournalPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RegistryJournal struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RegistryJournal: %s", err)
	}

	count, err = RegistryJournals().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
