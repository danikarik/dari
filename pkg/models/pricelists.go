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
)

// Pricelist is an object representing the database table.
type Pricelist struct {
	ID             uint      `boil:"id" json:"id" toml:"id" yaml:"id"`
	ListName       string    `boil:"list_name" json:"list_name" toml:"list_name" yaml:"list_name"`
	ManufacturerID uint      `boil:"manufacturer_id" json:"manufacturer_id" toml:"manufacturer_id" yaml:"manufacturer_id"`
	Year           int       `boil:"year" json:"year" toml:"year" yaml:"year"`
	IsActive       bool      `boil:"is_active" json:"is_active" toml:"is_active" yaml:"is_active"`
	IsArchive      bool      `boil:"is_archive" json:"is_archive" toml:"is_archive" yaml:"is_archive"`
	UserID         uint      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	CreatedAt      null.Time `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt      null.Time `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`

	R *pricelistR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pricelistL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PricelistColumns = struct {
	ID             string
	ListName       string
	ManufacturerID string
	Year           string
	IsActive       string
	IsArchive      string
	UserID         string
	CreatedAt      string
	UpdatedAt      string
}{
	ID:             "id",
	ListName:       "list_name",
	ManufacturerID: "manufacturer_id",
	Year:           "year",
	IsActive:       "is_active",
	IsArchive:      "is_archive",
	UserID:         "user_id",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// PricelistRels is where relationship names are stored.
var PricelistRels = struct {
	Manufacturer string
	User         string
	Products     string
}{
	Manufacturer: "Manufacturer",
	User:         "User",
	Products:     "Products",
}

// pricelistR is where relationships are stored.
type pricelistR struct {
	Manufacturer *Manufacturer
	User         *User
	Products     ProductSlice
}

// NewStruct creates a new relationship struct
func (*pricelistR) NewStruct() *pricelistR {
	return &pricelistR{}
}

// pricelistL is where Load methods for each relationship are stored.
type pricelistL struct{}

var (
	pricelistColumns               = []string{"id", "list_name", "manufacturer_id", "year", "is_active", "is_archive", "user_id", "created_at", "updated_at"}
	pricelistColumnsWithoutDefault = []string{"list_name", "manufacturer_id", "year", "user_id", "created_at", "updated_at"}
	pricelistColumnsWithDefault    = []string{"id", "is_active", "is_archive"}
	pricelistPrimaryKeyColumns     = []string{"id"}
)

type (
	// PricelistSlice is an alias for a slice of pointers to Pricelist.
	// This should generally be used opposed to []Pricelist.
	PricelistSlice []*Pricelist
	// PricelistHook is the signature for custom Pricelist hook methods
	PricelistHook func(context.Context, boil.ContextExecutor, *Pricelist) error

	pricelistQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pricelistType                 = reflect.TypeOf(&Pricelist{})
	pricelistMapping              = queries.MakeStructMapping(pricelistType)
	pricelistPrimaryKeyMapping, _ = queries.BindMapping(pricelistType, pricelistMapping, pricelistPrimaryKeyColumns)
	pricelistInsertCacheMut       sync.RWMutex
	pricelistInsertCache          = make(map[string]insertCache)
	pricelistUpdateCacheMut       sync.RWMutex
	pricelistUpdateCache          = make(map[string]updateCache)
	pricelistUpsertCacheMut       sync.RWMutex
	pricelistUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var pricelistBeforeInsertHooks []PricelistHook
var pricelistBeforeUpdateHooks []PricelistHook
var pricelistBeforeDeleteHooks []PricelistHook
var pricelistBeforeUpsertHooks []PricelistHook

var pricelistAfterInsertHooks []PricelistHook
var pricelistAfterSelectHooks []PricelistHook
var pricelistAfterUpdateHooks []PricelistHook
var pricelistAfterDeleteHooks []PricelistHook
var pricelistAfterUpsertHooks []PricelistHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Pricelist) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Pricelist) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Pricelist) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Pricelist) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Pricelist) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Pricelist) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Pricelist) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Pricelist) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Pricelist) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range pricelistAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPricelistHook registers your hook function for all future operations.
func AddPricelistHook(hookPoint boil.HookPoint, pricelistHook PricelistHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		pricelistBeforeInsertHooks = append(pricelistBeforeInsertHooks, pricelistHook)
	case boil.BeforeUpdateHook:
		pricelistBeforeUpdateHooks = append(pricelistBeforeUpdateHooks, pricelistHook)
	case boil.BeforeDeleteHook:
		pricelistBeforeDeleteHooks = append(pricelistBeforeDeleteHooks, pricelistHook)
	case boil.BeforeUpsertHook:
		pricelistBeforeUpsertHooks = append(pricelistBeforeUpsertHooks, pricelistHook)
	case boil.AfterInsertHook:
		pricelistAfterInsertHooks = append(pricelistAfterInsertHooks, pricelistHook)
	case boil.AfterSelectHook:
		pricelistAfterSelectHooks = append(pricelistAfterSelectHooks, pricelistHook)
	case boil.AfterUpdateHook:
		pricelistAfterUpdateHooks = append(pricelistAfterUpdateHooks, pricelistHook)
	case boil.AfterDeleteHook:
		pricelistAfterDeleteHooks = append(pricelistAfterDeleteHooks, pricelistHook)
	case boil.AfterUpsertHook:
		pricelistAfterUpsertHooks = append(pricelistAfterUpsertHooks, pricelistHook)
	}
}

// One returns a single pricelist record from the query.
func (q pricelistQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Pricelist, error) {
	o := &Pricelist{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pricelists")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Pricelist records from the query.
func (q pricelistQuery) All(ctx context.Context, exec boil.ContextExecutor) (PricelistSlice, error) {
	var o []*Pricelist

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pricelist slice")
	}

	if len(pricelistAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Pricelist records in the query.
func (q pricelistQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pricelists rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q pricelistQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pricelists exists")
	}

	return count > 0, nil
}

// Manufacturer pointed to by the foreign key.
func (o *Pricelist) Manufacturer(mods ...qm.QueryMod) manufacturerQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.ManufacturerID),
	}

	queryMods = append(queryMods, mods...)

	query := Manufacturers(queryMods...)
	queries.SetFrom(query.Query, "`manufacturers`")

	return query
}

// User pointed to by the foreign key.
func (o *Pricelist) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("id=?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// Products retrieves all the product's Products with an executor.
func (o *Pricelist) Products(mods ...qm.QueryMod) productQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`products`.`pricelist_id`=?", o.ID),
	)

	query := Products(queryMods...)
	queries.SetFrom(query.Query, "`products`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`products`.*"})
	}

	return query
}

// LoadManufacturer allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pricelistL) LoadManufacturer(ctx context.Context, e boil.ContextExecutor, singular bool, maybePricelist interface{}, mods queries.Applicator) error {
	var slice []*Pricelist
	var object *Pricelist

	if singular {
		object = maybePricelist.(*Pricelist)
	} else {
		slice = *maybePricelist.(*[]*Pricelist)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pricelistR{}
		}
		args = append(args, object.ManufacturerID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pricelistR{}
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

	if len(pricelistAfterSelectHooks) != 0 {
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
		foreign.R.Pricelists = append(foreign.R.Pricelists, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ManufacturerID == foreign.ID {
				local.R.Manufacturer = foreign
				if foreign.R == nil {
					foreign.R = &manufacturerR{}
				}
				foreign.R.Pricelists = append(foreign.R.Pricelists, local)
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (pricelistL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybePricelist interface{}, mods queries.Applicator) error {
	var slice []*Pricelist
	var object *Pricelist

	if singular {
		object = maybePricelist.(*Pricelist)
	} else {
		slice = *maybePricelist.(*[]*Pricelist)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pricelistR{}
		}
		args = append(args, object.UserID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pricelistR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)
		}
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(pricelistAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Pricelists = append(foreign.R.Pricelists, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Pricelists = append(foreign.R.Pricelists, local)
				break
			}
		}
	}

	return nil
}

// LoadProducts allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pricelistL) LoadProducts(ctx context.Context, e boil.ContextExecutor, singular bool, maybePricelist interface{}, mods queries.Applicator) error {
	var slice []*Pricelist
	var object *Pricelist

	if singular {
		object = maybePricelist.(*Pricelist)
	} else {
		slice = *maybePricelist.(*[]*Pricelist)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pricelistR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pricelistR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`products`), qm.WhereIn(`pricelist_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load products")
	}

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice products")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on products")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for products")
	}

	if len(productAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Products = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &productR{}
			}
			foreign.R.Pricelist = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PricelistID {
				local.R.Products = append(local.R.Products, foreign)
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.Pricelist = local
				break
			}
		}
	}

	return nil
}

// SetManufacturer of the pricelist to the related item.
// Sets o.R.Manufacturer to related.
// Adds o to related.R.Pricelists.
func (o *Pricelist) SetManufacturer(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Manufacturer) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `pricelists` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"manufacturer_id"}),
		strmangle.WhereClause("`", "`", 0, pricelistPrimaryKeyColumns),
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
		o.R = &pricelistR{
			Manufacturer: related,
		}
	} else {
		o.R.Manufacturer = related
	}

	if related.R == nil {
		related.R = &manufacturerR{
			Pricelists: PricelistSlice{o},
		}
	} else {
		related.R.Pricelists = append(related.R.Pricelists, o)
	}

	return nil
}

// SetUser of the pricelist to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Pricelists.
func (o *Pricelist) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `pricelists` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, pricelistPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &pricelistR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Pricelists: PricelistSlice{o},
		}
	} else {
		related.R.Pricelists = append(related.R.Pricelists, o)
	}

	return nil
}

// AddProducts adds the given related objects to the existing relationships
// of the pricelist, optionally inserting them as new records.
// Appends related to o.R.Products.
// Sets related.R.Pricelist appropriately.
func (o *Pricelist) AddProducts(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Product) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PricelistID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `products` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"pricelist_id"}),
				strmangle.WhereClause("`", "`", 0, productPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PricelistID = o.ID
		}
	}

	if o.R == nil {
		o.R = &pricelistR{
			Products: related,
		}
	} else {
		o.R.Products = append(o.R.Products, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &productR{
				Pricelist: o,
			}
		} else {
			rel.R.Pricelist = o
		}
	}
	return nil
}

// Pricelists retrieves all the records using an executor.
func Pricelists(mods ...qm.QueryMod) pricelistQuery {
	mods = append(mods, qm.From("`pricelists`"))
	return pricelistQuery{NewQuery(mods...)}
}

// FindPricelist retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPricelist(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*Pricelist, error) {
	pricelistObj := &Pricelist{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `pricelists` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, pricelistObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pricelists")
	}

	return pricelistObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Pricelist) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pricelists provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(pricelistColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pricelistInsertCacheMut.RLock()
	cache, cached := pricelistInsertCache[key]
	pricelistInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pricelistColumns,
			pricelistColumnsWithDefault,
			pricelistColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pricelistType, pricelistMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pricelistType, pricelistMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `pricelists` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `pricelists` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `pricelists` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, pricelistPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into pricelists")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == pricelistMapping["ID"] {
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
		return errors.Wrap(err, "models: unable to populate default values for pricelists")
	}

CacheNoHooks:
	if !cached {
		pricelistInsertCacheMut.Lock()
		pricelistInsertCache[key] = cache
		pricelistInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Pricelist.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Pricelist) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	queries.SetScanner(&o.UpdatedAt, currTime)

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	pricelistUpdateCacheMut.RLock()
	cache, cached := pricelistUpdateCache[key]
	pricelistUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pricelistColumns,
			pricelistPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pricelists, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `pricelists` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, pricelistPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pricelistType, pricelistMapping, append(wl, pricelistPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update pricelists row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pricelists")
	}

	if !cached {
		pricelistUpdateCacheMut.Lock()
		pricelistUpdateCache[key] = cache
		pricelistUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q pricelistQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pricelists")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pricelists")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PricelistSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricelistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `pricelists` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pricelistPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pricelist slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pricelist")
	}
	return rowsAff, nil
}

var mySQLPricelistUniqueColumns = []string{
	"id",
	"list_name",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Pricelist) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pricelists provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if queries.MustTime(o.CreatedAt).IsZero() {
		queries.SetScanner(&o.CreatedAt, currTime)
	}
	queries.SetScanner(&o.UpdatedAt, currTime)

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(pricelistColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPricelistUniqueColumns, o)

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

	pricelistUpsertCacheMut.RLock()
	cache, cached := pricelistUpsertCache[key]
	pricelistUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pricelistColumns,
			pricelistColumnsWithDefault,
			pricelistColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pricelistColumns,
			pricelistPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert pricelists, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "pricelists", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `pricelists` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(pricelistType, pricelistMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pricelistType, pricelistMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for pricelists")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == pricelistMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(pricelistType, pricelistMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for pricelists")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for pricelists")
	}

CacheNoHooks:
	if !cached {
		pricelistUpsertCacheMut.Lock()
		pricelistUpsertCache[key] = cache
		pricelistUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Pricelist record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pricelist) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pricelist provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pricelistPrimaryKeyMapping)
	sql := "DELETE FROM `pricelists` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pricelists")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pricelists")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q pricelistQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pricelistQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pricelists")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pricelists")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PricelistSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pricelist slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(pricelistBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricelistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `pricelists` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pricelistPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pricelist slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pricelists")
	}

	if len(pricelistAfterDeleteHooks) != 0 {
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
func (o *Pricelist) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPricelist(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PricelistSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PricelistSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pricelistPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `pricelists`.* FROM `pricelists` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pricelistPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PricelistSlice")
	}

	*o = slice

	return nil
}

// PricelistExists checks if the Pricelist row exists.
func PricelistExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `pricelists` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pricelists exists")
	}

	return exists, nil
}
