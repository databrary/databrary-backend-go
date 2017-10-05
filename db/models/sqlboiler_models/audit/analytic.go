// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package audit

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/databrary/databrary/db/models/custom_types"
	"github.com/databrary/sqlboiler/boil"
	"github.com/databrary/sqlboiler/queries"
	"github.com/databrary/sqlboiler/queries/qm"
	"github.com/databrary/sqlboiler/strmangle"
	"github.com/pkg/errors"
	"gopkg.in/nullbio/null.v6"
	"reflect"
	"strings"
	"sync"
	"time"
)

// Analytic is an object representing the database table.
type Analytic struct {
	AuditTime   time.Time           `db:"audit_time" json:"analytic_audit_time"`
	AuditUser   int                 `db:"audit_user" json:"analytic_audit_user"`
	AuditIP     custom_types.Inet   `db:"audit_ip" json:"analytic_audit_ip"`
	AuditAction custom_types.Action `db:"audit_action" json:"analytic_audit_action"`
	Route       string              `db:"route" json:"analytic_route"`
	Data        null.JSON           `db:"data" json:"analytic_data,omitempty"`

	R *analyticR `db:"-" json:"-"`
	L analyticL  `db:"-" json:"-"`
}

// analyticR is where relationships are stored.
type analyticR struct {
}

// analyticL is where Load methods for each relationship are stored.
type analyticL struct{}

var (
	analyticColumns               = []string{"audit_time", "audit_user", "audit_ip", "audit_action", "route", "data"}
	analyticColumnsWithoutDefault = []string{"audit_user", "audit_ip", "audit_action", "route", "data"}
	analyticColumnsWithDefault    = []string{"audit_time"}
	analyticColumnsWithCustom     = []string{"audit_ip", "audit_action"}
)

type (
	// AnalyticSlice is an alias for a slice of pointers to Analytic.
	// This should generally be used opposed to []Analytic.
	AnalyticSlice []*Analytic
	// AnalyticHook is the signature for custom Analytic hook methods
	AnalyticHook func(boil.Executor, *Analytic) error

	analyticQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	analyticType    = reflect.TypeOf(&Analytic{})
	analyticMapping = queries.MakeStructMapping(analyticType)

	analyticInsertCacheMut sync.RWMutex
	analyticInsertCache    = make(map[string]insertCache)
	analyticUpdateCacheMut sync.RWMutex
	analyticUpdateCache    = make(map[string]updateCache)
	analyticUpsertCacheMut sync.RWMutex
	analyticUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var analyticBeforeInsertHooks []AnalyticHook
var analyticBeforeUpdateHooks []AnalyticHook
var analyticBeforeDeleteHooks []AnalyticHook
var analyticBeforeUpsertHooks []AnalyticHook

var analyticAfterInsertHooks []AnalyticHook
var analyticAfterSelectHooks []AnalyticHook
var analyticAfterUpdateHooks []AnalyticHook
var analyticAfterDeleteHooks []AnalyticHook
var analyticAfterUpsertHooks []AnalyticHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Analytic) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Analytic) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Analytic) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Analytic) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Analytic) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Analytic) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Analytic) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Analytic) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Analytic) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range analyticAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnalyticHook registers your hook function for all future operations.
func AddAnalyticHook(hookPoint boil.HookPoint, analyticHook AnalyticHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		analyticBeforeInsertHooks = append(analyticBeforeInsertHooks, analyticHook)
	case boil.BeforeUpdateHook:
		analyticBeforeUpdateHooks = append(analyticBeforeUpdateHooks, analyticHook)
	case boil.BeforeDeleteHook:
		analyticBeforeDeleteHooks = append(analyticBeforeDeleteHooks, analyticHook)
	case boil.BeforeUpsertHook:
		analyticBeforeUpsertHooks = append(analyticBeforeUpsertHooks, analyticHook)
	case boil.AfterInsertHook:
		analyticAfterInsertHooks = append(analyticAfterInsertHooks, analyticHook)
	case boil.AfterSelectHook:
		analyticAfterSelectHooks = append(analyticAfterSelectHooks, analyticHook)
	case boil.AfterUpdateHook:
		analyticAfterUpdateHooks = append(analyticAfterUpdateHooks, analyticHook)
	case boil.AfterDeleteHook:
		analyticAfterDeleteHooks = append(analyticAfterDeleteHooks, analyticHook)
	case boil.AfterUpsertHook:
		analyticAfterUpsertHooks = append(analyticAfterUpsertHooks, analyticHook)
	}
}

// OneP returns a single analytic record from the query, and panics on error.
func (q analyticQuery) OneP() *Analytic {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single analytic record from the query.
func (q analyticQuery) One() (*Analytic, error) {
	o := &Analytic{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for analytic")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Analytic records from the query, and panics on error.
func (q analyticQuery) AllP() AnalyticSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Analytic records from the query.
func (q analyticQuery) All() (AnalyticSlice, error) {
	var o AnalyticSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Analytic slice")
	}

	if len(analyticAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Analytic records in the query, and panics on error.
func (q analyticQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Analytic records in the query.
func (q analyticQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count analytic rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q analyticQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q analyticQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if analytic exists")
	}

	return count > 0, nil
}

// AnalyticsG retrieves all records.
func AnalyticsG(mods ...qm.QueryMod) analyticQuery {
	return Analytics(boil.GetDB(), mods...)
}

// Analytics retrieves all the records using an executor.
func Analytics(exec boil.Executor, mods ...qm.QueryMod) analyticQuery {
	mods = append(mods, qm.From("\"audit\".\"analytic\""))
	return analyticQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Analytic) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Analytic) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Analytic) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Analytic) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no analytic provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(analyticColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	analyticInsertCacheMut.RLock()
	cache, cached := analyticInsertCache[key]
	analyticInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			analyticColumns,
			analyticColumnsWithDefault,
			analyticColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(analyticType, analyticMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(analyticType, analyticMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit\".\"analytic\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit\".\"analytic\" DEFAULT VALUES"
		}

		if len(cache.retMapping) != 0 {
			cache.query += fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into analytic")
	}

	if !cached {
		analyticInsertCacheMut.Lock()
		analyticInsertCache[key] = cache
		analyticInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}
