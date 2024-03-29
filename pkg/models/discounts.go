// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// Discount is an object representing the database table.
type Discount struct {
	ID             uint              `boil:"id" json:"id" toml:"id" yaml:"id"`
	Naming         string            `boil:"naming" json:"naming" toml:"naming" yaml:"naming"`
	ManufacturerID uint              `boil:"manufacturer_id" json:"manufacturer_id" toml:"manufacturer_id" yaml:"manufacturer_id"`
	CategoryID     uint              `boil:"category_id" json:"category_id" toml:"category_id" yaml:"category_id"`
	SubcategoryID  null.Uint         `boil:"subcategory_id" json:"subcategory_id,omitempty" toml:"subcategory_id" yaml:"subcategory_id,omitempty"`
	IsFixed        bool              `boil:"is_fixed" json:"is_fixed" toml:"is_fixed" yaml:"is_fixed"`
	ValueFrom      types.Decimal     `boil:"value_from" json:"value_from" toml:"value_from" yaml:"value_from"`
	ValueTo        types.NullDecimal `boil:"value_to" json:"value_to,omitempty" toml:"value_to" yaml:"value_to,omitempty"`
	CreatedAt      null.Time         `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt      null.Time         `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *discountR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L discountL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DiscountColumns = struct {
	ID             string
	Naming         string
	ManufacturerID string
	CategoryID     string
	SubcategoryID  string
	IsFixed        string
	ValueFrom      string
	ValueTo        string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	Naming:         "naming",
	ManufacturerID: "manufacturer_id",
	CategoryID:     "category_id",
	SubcategoryID:  "subcategory_id",
	IsFixed:        "is_fixed",
	ValueFrom:      "value_from",
	ValueTo:        "value_to",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// DiscountRels is where relationship names are stored.
var DiscountRels = struct {
	Category     string
	Manufacturer string
	Subcategory  string
}{
	Category:     "Category",
	Manufacturer: "Manufacturer",
	Subcategory:  "Subcategory",
}

// discountR is where relationships are stored.
type discountR struct {
	Category     *Category
	Manufacturer *Manufacturer
	Subcategory  *Subcategory
}

// NewStruct creates a new relationship struct
func (*discountR) NewStruct() *discountR {
	return &discountR{}
}

// discountL is where Load methods for each relationship are stored.
type discountL struct{}

var (
	discountColumns               = []string{"id", "naming", "manufacturer_id", "category_id", "subcategory_id", "is_fixed", "value_from", "value_to", "created_at", "updated_at"}
	discountColumnsWithoutDefault = []string{"naming", "manufacturer_id", "category_id", "subcategory_id", "value_to", "created_at", "updated_at"}
	discountColumnsWithDefault    = []string{"id", "is_fixed", "value_from"}
	discountPrimaryKeyColumns     = []string{"id"}
)

type (
	// DiscountSlice is an alias for a slice of pointers to Discount.
	// This should generally be used opposed to []Discount.
	DiscountSlice []*Discount
	// DiscountHook is the signature for custom Discount hook methods
	DiscountHook func(context.Context, boil.ContextExecutor, *Discount) error

	discountQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	discountType                 = reflect.TypeOf(&Discount{})
	discountMapping              = queries.MakeStructMapping(discountType)
	discountPrimaryKeyMapping, _ = queries.BindMapping(discountType, discountMapping, discountPrimaryKeyColumns)
	discountInsertCacheMut       sync.RWMutex
	discountInsertCache          = make(map[string]insertCache)
	discountUpdateCacheMut       sync.RWMutex
	discountUpdateCache          = make(map[string]updateCache)
	discountUpsertCacheMut       sync.RWMutex
	discountUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var discountBeforeInsertHooks []DiscountHook
var discountBeforeUpdateHooks []DiscountHook
var discountBeforeDeleteHooks []DiscountHook
var discountBeforeUpsertHooks []DiscountHook

var discountAfterInsertHooks []DiscountHook
var discountAfterSelectHooks []DiscountHook
var discountAfterUpdateHooks []DiscountHook
var discountAfterDeleteHooks []DiscountHook
var discountAfterUpsertHooks []DiscountHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Discount) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Discount) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Discount) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Discount) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Discount) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Discount) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Discount) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Discount) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Discount) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range discountAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDiscountHook registers your hook function for all future operations.
func AddDiscountHook(hookPoint boil.HookPoint, discountHook DiscountHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		discountBeforeInsertHooks = append(discountBeforeInsertHooks, discountHook)
	case boil.BeforeUpdateHook:
		discountBeforeUpdateHooks = append(discountBeforeUpdateHooks, discountHook)
	case boil.BeforeDeleteHook:
		discountBeforeDeleteHooks = append(discountBeforeDeleteHooks, discountHook)
	case boil.BeforeUpsertHook:
		discountBeforeUpsertHooks = append(discountBeforeUpsertHooks, discountHook)
	case boil.AfterInsertHook:
		discountAfterInsertHooks = append(discountAfterInsertHooks, discountHook)
	case boil.AfterSelectHook:
		discountAfterSelectHooks = append(discountAfterSelectHooks, discountHook)
	case boil.AfterUpdateHook:
		discountAfterUpdateHooks = append(discountAfterUpdateHooks, discountHook)
	case boil.AfterDeleteHook:
		discountAfterDeleteHooks = append(discountAfterDeleteHooks, discountHook)
	case boil.AfterUpsertHook:
		discountAfterUpsertHooks = append(discountAfterUpsertHooks, discountHook)
	}
}

// One returns a single discount record from the query.
func (q discountQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Discount, error) {
	o := &Discount{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for discounts")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Discount records from the query.
func (q discountQuery) All(ctx context.Context, exec boil.ContextExecutor) (DiscountSlice, error) {
	var o []*Discount

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Discount slice")
	}

	if len(discountAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Discount records in the query.
func (q discountQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count discounts rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q discountQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if discounts exists")
	}

	return count > 0, nil
}

// Category pointed to by the foreign key.
func (o *Discount) Category(mods ...qm.QueryMod) categoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.CategoryID),
	}

	queryMods = append(queryMods, mods...)

	query := Categories(queryMods...)
	queries.SetFrom(query.Query, "`categories`")

	return query
}

// Manufacturer pointed to by the foreign key.
func (o *Discount) Manufacturer(mods ...qm.QueryMod) manufacturerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ManufacturerID),
	}

	queryMods = append(queryMods, mods...)

	query := Manufacturers(queryMods...)
	queries.SetFrom(query.Query, "`manufacturers`")

	return query
}

// Subcategory pointed to by the foreign key.
func (o *Discount) Subcategory(mods ...qm.QueryMod) subcategoryQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.SubcategoryID),
	}

	queryMods = append(queryMods, mods...)

	query := Subcategories(queryMods...)
	queries.SetFrom(query.Query, "`subcategories`")

	return query
}

// LoadCategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (discountL) LoadCategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDiscount interface{}, mods queries.Applicator) error {
	var slice []*Discount
	var object *Discount

	if singular {
		object = maybeDiscount.(*Discount)
	} else {
		slice = *maybeDiscount.(*[]*Discount)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &discountR{}
		}
		args = append(args, object.CategoryID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &discountR{}
			}

			for _, a := range args {
				if a == obj.CategoryID {
					continue Outer
				}
			}

			args = append(args, obj.CategoryID)
		}
	}

	query := NewQuery(qm.From(`categories`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Category")
	}

	var resultSlice []*Category
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Category")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for categories")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for categories")
	}

	if len(discountAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Category = foreign
		if foreign.R == nil {
			foreign.R = &categoryR{}
		}
		foreign.R.Discounts = append(foreign.R.Discounts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.CategoryID == foreign.ID {
				local.R.Category = foreign
				if foreign.R == nil {
					foreign.R = &categoryR{}
				}
				foreign.R.Discounts = append(foreign.R.Discounts, local)
				break
			}
		}
	}

	return nil
}

// LoadManufacturer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (discountL) LoadManufacturer(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDiscount interface{}, mods queries.Applicator) error {
	var slice []*Discount
	var object *Discount

	if singular {
		object = maybeDiscount.(*Discount)
	} else {
		slice = *maybeDiscount.(*[]*Discount)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &discountR{}
		}
		args = append(args, object.ManufacturerID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &discountR{}
			}

			for _, a := range args {
				if a == obj.ManufacturerID {
					continue Outer
				}
			}

			args = append(args, obj.ManufacturerID)
		}
	}

	query := NewQuery(qm.From(`manufacturers`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Manufacturer")
	}

	var resultSlice []*Manufacturer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Manufacturer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for manufacturers")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for manufacturers")
	}

	if len(discountAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Manufacturer = foreign
		if foreign.R == nil {
			foreign.R = &manufacturerR{}
		}
		foreign.R.Discounts = append(foreign.R.Discounts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ManufacturerID == foreign.ID {
				local.R.Manufacturer = foreign
				if foreign.R == nil {
					foreign.R = &manufacturerR{}
				}
				foreign.R.Discounts = append(foreign.R.Discounts, local)
				break
			}
		}
	}

	return nil
}

// LoadSubcategory allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (discountL) LoadSubcategory(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDiscount interface{}, mods queries.Applicator) error {
	var slice []*Discount
	var object *Discount

	if singular {
		object = maybeDiscount.(*Discount)
	} else {
		slice = *maybeDiscount.(*[]*Discount)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &discountR{}
		}
		args = append(args, object.SubcategoryID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &discountR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.SubcategoryID) {
					continue Outer
				}
			}

			args = append(args, obj.SubcategoryID)
		}
	}

	query := NewQuery(qm.From(`subcategories`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Subcategory")
	}

	var resultSlice []*Subcategory
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Subcategory")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for subcategories")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for subcategories")
	}

	if len(discountAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Subcategory = foreign
		if foreign.R == nil {
			foreign.R = &subcategoryR{}
		}
		foreign.R.Discounts = append(foreign.R.Discounts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.SubcategoryID, foreign.ID) {
				local.R.Subcategory = foreign
				if foreign.R == nil {
					foreign.R = &subcategoryR{}
				}
				foreign.R.Discounts = append(foreign.R.Discounts, local)
				break
			}
		}
	}

	return nil
}

// SetCategory of the discount to the related item.
// Sets o.R.Category to related.
// Adds o to related.R.Discounts.
func (o *Discount) SetCategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Category) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `discounts` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"category_id"}),
		strmangle.WhereClause("`", "`", 0, discountPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.CategoryID = related.ID
	if o.R == nil {
		o.R = &discountR{
			Category: related,
		}
	} else {
		o.R.Category = related
	}

	if related.R == nil {
		related.R = &categoryR{
			Discounts: DiscountSlice{o},
		}
	} else {
		related.R.Discounts = append(related.R.Discounts, o)
	}

	return nil
}

// SetManufacturer of the discount to the related item.
// Sets o.R.Manufacturer to related.
// Adds o to related.R.Discounts.
func (o *Discount) SetManufacturer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Manufacturer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `discounts` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"manufacturer_id"}),
		strmangle.WhereClause("`", "`", 0, discountPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ManufacturerID = related.ID
	if o.R == nil {
		o.R = &discountR{
			Manufacturer: related,
		}
	} else {
		o.R.Manufacturer = related
	}

	if related.R == nil {
		related.R = &manufacturerR{
			Discounts: DiscountSlice{o},
		}
	} else {
		related.R.Discounts = append(related.R.Discounts, o)
	}

	return nil
}

// SetSubcategory of the discount to the related item.
// Sets o.R.Subcategory to related.
// Adds o to related.R.Discounts.
func (o *Discount) SetSubcategory(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Subcategory) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `discounts` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"subcategory_id"}),
		strmangle.WhereClause("`", "`", 0, discountPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.SubcategoryID, related.ID)
	if o.R == nil {
		o.R = &discountR{
			Subcategory: related,
		}
	} else {
		o.R.Subcategory = related
	}

	if related.R == nil {
		related.R = &subcategoryR{
			Discounts: DiscountSlice{o},
		}
	} else {
		related.R.Discounts = append(related.R.Discounts, o)
	}

	return nil
}

// RemoveSubcategory relationship.
// Sets o.R.Subcategory to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Discount) RemoveSubcategory(ctx context.Context, exec boil.ContextExecutor, related *Subcategory) error {
	var err error

	queries.SetScanner(&o.SubcategoryID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("subcategory_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Subcategory = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.Discounts {
		if queries.Equal(o.SubcategoryID, ri.SubcategoryID) {
			continue
		}

		ln := len(related.R.Discounts)
		if ln > 1 && i < ln-1 {
			related.R.Discounts[i] = related.R.Discounts[ln-1]
		}
		related.R.Discounts = related.R.Discounts[:ln-1]
		break
	}
	return nil
}

// Discounts retrieves all the records using an executor.
func Discounts(mods ...qm.QueryMod) discountQuery {
	mods = append(mods, qm.From("`discounts`"))
	return discountQuery{NewQuery(mods...)}
}

// FindDiscount retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDiscount(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Discount, error) {
	discountObj := &Discount{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `discounts` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, discountObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from discounts")
	}

	return discountObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Discount) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no discounts provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	if queries.MustTime(o.UpdatedAt).IsZero() {
		queries.SetScanner(&o.UpdatedAt, currTime)
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(discountColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	discountInsertCacheMut.RLock()
	cache, cached := discountInsertCache[key]
	discountInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			discountColumns,
			discountColumnsWithDefault,
			discountColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(discountType, discountMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(discountType, discountMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `discounts` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `discounts` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `discounts` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, discountPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into discounts")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == discountMapping["ID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for discounts")
	}

CacheNoHooks:
	if !cached {
		discountInsertCacheMut.Lock()
		discountInsertCache[key] = cache
		discountInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Discount.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Discount) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	discountUpdateCacheMut.RLock()
	cache, cached := discountUpdateCache[key]
	discountUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			discountColumns,
			discountPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update discounts, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `discounts` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, discountPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(discountType, discountMapping, append(wl, discountPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update discounts row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for discounts")
	}

	if !cached {
		discountUpdateCacheMut.Lock()
		discountUpdateCache[key] = cache
		discountUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q discountQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for discounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for discounts")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DiscountSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `discounts` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, discountPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in discount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all discount")
	}
	return rowsAff, nil
}

var mySQLDiscountUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Discount) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no discounts provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(discountColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDiscountUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	discountUpsertCacheMut.RLock()
	cache, cached := discountUpsertCache[key]
	discountUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			discountColumns,
			discountColumnsWithDefault,
			discountColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			discountColumns,
			discountPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert discounts, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "discounts", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `discounts` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(discountType, discountMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(discountType, discountMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for discounts")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == discountMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(discountType, discountMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for discounts")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for discounts")
	}

CacheNoHooks:
	if !cached {
		discountUpsertCacheMut.Lock()
		discountUpsertCache[key] = cache
		discountUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Discount record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Discount) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Discount provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), discountPrimaryKeyMapping)
	sql := "DELETE FROM `discounts` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from discounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for discounts")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q discountQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no discountQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from discounts")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for discounts")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DiscountSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Discount slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(discountBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `discounts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, discountPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from discount slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for discounts")
	}

	if len(discountAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Discount) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDiscount(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DiscountSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DiscountSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), discountPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `discounts`.* FROM `discounts` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, discountPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DiscountSlice")
	}

	*o = slice

	return nil
}

// DiscountExists checks if the Discount row exists.
func DiscountExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `discounts` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if discounts exists")
	}

	return exists, nil
}
