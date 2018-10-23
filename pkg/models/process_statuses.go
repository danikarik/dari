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
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// ProcessStatus is an object representing the database table.
type ProcessStatus struct {
	ID    uint   `boil:"id" json:"id" toml:"id" yaml:"id"`
	Title string `boil:"title" json:"title" toml:"title" yaml:"title"`

	R *processStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L processStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ProcessStatusColumns = struct {
	ID    string
	Title string
}{
	ID:    "id",
	Title: "title",
}

// ProcessStatusRels is where relationship names are stored.
var ProcessStatusRels = struct {
	RegistryJournals string
}{
	RegistryJournals: "RegistryJournals",
}

// processStatusR is where relationships are stored.
type processStatusR struct {
	RegistryJournals RegistryJournalSlice
}

// NewStruct creates a new relationship struct
func (*processStatusR) NewStruct() *processStatusR {
	return &processStatusR{}
}

// processStatusL is where Load methods for each relationship are stored.
type processStatusL struct{}

var (
	processStatusColumns               = []string{"id", "title"}
	processStatusColumnsWithoutDefault = []string{"id", "title"}
	processStatusColumnsWithDefault    = []string{}
	processStatusPrimaryKeyColumns     = []string{"id"}
)

type (
	// ProcessStatusSlice is an alias for a slice of pointers to ProcessStatus.
	// This should generally be used opposed to []ProcessStatus.
	ProcessStatusSlice []*ProcessStatus
	// ProcessStatusHook is the signature for custom ProcessStatus hook methods
	ProcessStatusHook func(context.Context, boil.ContextExecutor, *ProcessStatus) error

	processStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	processStatusType                 = reflect.TypeOf(&ProcessStatus{})
	processStatusMapping              = queries.MakeStructMapping(processStatusType)
	processStatusPrimaryKeyMapping, _ = queries.BindMapping(processStatusType, processStatusMapping, processStatusPrimaryKeyColumns)
	processStatusInsertCacheMut       sync.RWMutex
	processStatusInsertCache          = make(map[string]insertCache)
	processStatusUpdateCacheMut       sync.RWMutex
	processStatusUpdateCache          = make(map[string]updateCache)
	processStatusUpsertCacheMut       sync.RWMutex
	processStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
)

var processStatusBeforeInsertHooks []ProcessStatusHook
var processStatusBeforeUpdateHooks []ProcessStatusHook
var processStatusBeforeDeleteHooks []ProcessStatusHook
var processStatusBeforeUpsertHooks []ProcessStatusHook

var processStatusAfterInsertHooks []ProcessStatusHook
var processStatusAfterSelectHooks []ProcessStatusHook
var processStatusAfterUpdateHooks []ProcessStatusHook
var processStatusAfterDeleteHooks []ProcessStatusHook
var processStatusAfterUpsertHooks []ProcessStatusHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *ProcessStatus) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *ProcessStatus) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *ProcessStatus) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *ProcessStatus) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *ProcessStatus) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *ProcessStatus) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *ProcessStatus) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *ProcessStatus) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *ProcessStatus) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	for _, hook := range processStatusAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddProcessStatusHook registers your hook function for all future operations.
func AddProcessStatusHook(hookPoint boil.HookPoint, processStatusHook ProcessStatusHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		processStatusBeforeInsertHooks = append(processStatusBeforeInsertHooks, processStatusHook)
	case boil.BeforeUpdateHook:
		processStatusBeforeUpdateHooks = append(processStatusBeforeUpdateHooks, processStatusHook)
	case boil.BeforeDeleteHook:
		processStatusBeforeDeleteHooks = append(processStatusBeforeDeleteHooks, processStatusHook)
	case boil.BeforeUpsertHook:
		processStatusBeforeUpsertHooks = append(processStatusBeforeUpsertHooks, processStatusHook)
	case boil.AfterInsertHook:
		processStatusAfterInsertHooks = append(processStatusAfterInsertHooks, processStatusHook)
	case boil.AfterSelectHook:
		processStatusAfterSelectHooks = append(processStatusAfterSelectHooks, processStatusHook)
	case boil.AfterUpdateHook:
		processStatusAfterUpdateHooks = append(processStatusAfterUpdateHooks, processStatusHook)
	case boil.AfterDeleteHook:
		processStatusAfterDeleteHooks = append(processStatusAfterDeleteHooks, processStatusHook)
	case boil.AfterUpsertHook:
		processStatusAfterUpsertHooks = append(processStatusAfterUpsertHooks, processStatusHook)
	}
}

// One returns a single processStatus record from the query.
func (q processStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*ProcessStatus, error) {
	o := &ProcessStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for process_statuses")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all ProcessStatus records from the query.
func (q processStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (ProcessStatusSlice, error) {
	var o []*ProcessStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to ProcessStatus slice")
	}

	if len(processStatusAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all ProcessStatus records in the query.
func (q processStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count process_statuses rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q processStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if process_statuses exists")
	}

	return count > 0, nil
}

// RegistryJournals retrieves all the registry_journal's RegistryJournals with an executor.
func (o *ProcessStatus) RegistryJournals(mods ...qm.QueryMod) registryJournalQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`registry_journals`.`process_status_id`=?", o.ID),
	)

	query := RegistryJournals(queryMods...)
	queries.SetFrom(query.Query, "`registry_journals`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`registry_journals`.*"})
	}

	return query
}

// LoadRegistryJournals allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (processStatusL) LoadRegistryJournals(ctx context.Context, e boil.ContextExecutor, singular bool, maybeProcessStatus interface{}, mods queries.Applicator) error {
	var slice []*ProcessStatus
	var object *ProcessStatus

	if singular {
		object = maybeProcessStatus.(*ProcessStatus)
	} else {
		slice = *maybeProcessStatus.(*[]*ProcessStatus)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &processStatusR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &processStatusR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	query := NewQuery(qm.From(`registry_journals`), qm.WhereIn(`process_status_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load registry_journals")
	}

	var resultSlice []*RegistryJournal
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice registry_journals")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on registry_journals")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for registry_journals")
	}

	if len(registryJournalAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RegistryJournals = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &registryJournalR{}
			}
			foreign.R.ProcessStatus = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ProcessStatusID {
				local.R.RegistryJournals = append(local.R.RegistryJournals, foreign)
				if foreign.R == nil {
					foreign.R = &registryJournalR{}
				}
				foreign.R.ProcessStatus = local
				break
			}
		}
	}

	return nil
}

// AddRegistryJournals adds the given related objects to the existing relationships
// of the process_status, optionally inserting them as new records.
// Appends related to o.R.RegistryJournals.
// Sets related.R.ProcessStatus appropriately.
func (o *ProcessStatus) AddRegistryJournals(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RegistryJournal) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ProcessStatusID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `registry_journals` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"process_status_id"}),
				strmangle.WhereClause("`", "`", 0, registryJournalPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ProcessStatusID = o.ID
		}
	}

	if o.R == nil {
		o.R = &processStatusR{
			RegistryJournals: related,
		}
	} else {
		o.R.RegistryJournals = append(o.R.RegistryJournals, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &registryJournalR{
				ProcessStatus: o,
			}
		} else {
			rel.R.ProcessStatus = o
		}
	}
	return nil
}

// ProcessStatuses retrieves all the records using an executor.
func ProcessStatuses(mods ...qm.QueryMod) processStatusQuery {
	mods = append(mods, qm.From("`process_statuses`"))
	return processStatusQuery{NewQuery(mods...)}
}

// FindProcessStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindProcessStatus(ctx context.Context, exec boil.ContextExecutor, iD uint, selectCols ...string) (*ProcessStatus, error) {
	processStatusObj := &ProcessStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `process_statuses` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, processStatusObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from process_statuses")
	}

	return processStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *ProcessStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no process_statuses provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(processStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	processStatusInsertCacheMut.RLock()
	cache, cached := processStatusInsertCache[key]
	processStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			processStatusColumns,
			processStatusColumnsWithDefault,
			processStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(processStatusType, processStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(processStatusType, processStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `process_statuses` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `process_statuses` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `process_statuses` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, processStatusPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into process_statuses")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
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
		return errors.Wrap(err, "models: unable to populate default values for process_statuses")
	}

CacheNoHooks:
	if !cached {
		processStatusInsertCacheMut.Lock()
		processStatusInsertCache[key] = cache
		processStatusInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the ProcessStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *ProcessStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	processStatusUpdateCacheMut.RLock()
	cache, cached := processStatusUpdateCache[key]
	processStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			processStatusColumns,
			processStatusPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update process_statuses, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `process_statuses` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, processStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(processStatusType, processStatusMapping, append(wl, processStatusPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update process_statuses row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for process_statuses")
	}

	if !cached {
		processStatusUpdateCacheMut.Lock()
		processStatusUpdateCache[key] = cache
		processStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q processStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for process_statuses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for process_statuses")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ProcessStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), processStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `process_statuses` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, processStatusPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in processStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all processStatus")
	}
	return rowsAff, nil
}

var mySQLProcessStatusUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *ProcessStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no process_statuses provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(processStatusColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLProcessStatusUniqueColumns, o)

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

	processStatusUpsertCacheMut.RLock()
	cache, cached := processStatusUpsertCache[key]
	processStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			processStatusColumns,
			processStatusColumnsWithDefault,
			processStatusColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			processStatusColumns,
			processStatusPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert process_statuses, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "process_statuses", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `process_statuses` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(processStatusType, processStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(processStatusType, processStatusMapping, ret)
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

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for process_statuses")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(processStatusType, processStatusMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for process_statuses")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for process_statuses")
	}

CacheNoHooks:
	if !cached {
		processStatusUpsertCacheMut.Lock()
		processStatusUpsertCache[key] = cache
		processStatusUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single ProcessStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *ProcessStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProcessStatus provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), processStatusPrimaryKeyMapping)
	sql := "DELETE FROM `process_statuses` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from process_statuses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for process_statuses")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q processStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no processStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from process_statuses")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for process_statuses")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ProcessStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no ProcessStatus slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(processStatusBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), processStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `process_statuses` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, processStatusPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from processStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for process_statuses")
	}

	if len(processStatusAfterDeleteHooks) != 0 {
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
func (o *ProcessStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindProcessStatus(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ProcessStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ProcessStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), processStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `process_statuses`.* FROM `process_statuses` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, processStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ProcessStatusSlice")
	}

	*o = slice

	return nil
}

// ProcessStatusExists checks if the ProcessStatus row exists.
func ProcessStatusExists(ctx context.Context, exec boil.ContextExecutor, iD uint) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `process_statuses` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}

	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if process_statuses exists")
	}

	return exists, nil
}