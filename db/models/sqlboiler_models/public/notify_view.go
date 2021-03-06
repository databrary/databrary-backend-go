// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package public

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/databrary/databrary-backend-go/db/models/custom_types"
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

// NotifyView is an object representing the database view.
type NotifyView struct {
	Target   null.Int                        `db:"target" json:"notifyView_target,omitempty"`
	Notice   null.Int16                      `db:"notice" json:"notifyView_notice,omitempty"`
	Delivery custom_types.NullNoticeDelivery `db:"delivery" json:"notifyView_delivery,omitempty"`

	R *notifyViewR `db:"-" json:"-"`
	L notifyViewL  `db:"-" json:"-"`
}

// notifyViewR is where relationships are stored.
type notifyViewR struct {
}

// notifyViewL is where Load methods for each relationship are stored.
type notifyViewL struct{}

var (
	notifyViewColumns               = []string{"target", "notice", "delivery"}
	notifyViewColumnsWithoutDefault = []string{"target", "notice", "delivery"}
	notifyViewColumnsWithDefault    = []string{}
	notifyViewColumnsWithCustom     = []string{"delivery"}
)

type (
	// NotifyViewSlice is an alias for a slice of pointers to NotifyView.
	// This should generally be used opposed to []NotifyView.
	NotifyViewSlice []*NotifyView
	// NotifyViewHook is the signature for custom NotifyView hook methods
	NotifyViewHook func(boil.Executor, *NotifyView) error

	notifyViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	notifyViewType    = reflect.TypeOf(&NotifyView{})
	notifyViewMapping = queries.MakeStructMapping(notifyViewType)

	notifyViewInsertCacheMut sync.RWMutex
	notifyViewInsertCache    = make(map[string]insertCache)
	notifyViewUpdateCacheMut sync.RWMutex
	notifyViewUpdateCache    = make(map[string]updateCache)
	notifyViewUpsertCacheMut sync.RWMutex
	notifyViewUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var notifyViewBeforeInsertHooks []NotifyViewHook
var notifyViewBeforeUpdateHooks []NotifyViewHook
var notifyViewBeforeDeleteHooks []NotifyViewHook
var notifyViewBeforeUpsertHooks []NotifyViewHook

var notifyViewAfterInsertHooks []NotifyViewHook
var notifyViewAfterSelectHooks []NotifyViewHook
var notifyViewAfterUpdateHooks []NotifyViewHook
var notifyViewAfterDeleteHooks []NotifyViewHook
var notifyViewAfterUpsertHooks []NotifyViewHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *NotifyView) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *NotifyView) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *NotifyView) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *NotifyView) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *NotifyView) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *NotifyView) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *NotifyView) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *NotifyView) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *NotifyView) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range notifyViewAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNotifyViewHook registers your hook function for all future operations.
func AddNotifyViewHook(hookPoint boil.HookPoint, notifyViewHook NotifyViewHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		notifyViewBeforeInsertHooks = append(notifyViewBeforeInsertHooks, notifyViewHook)
	case boil.BeforeUpdateHook:
		notifyViewBeforeUpdateHooks = append(notifyViewBeforeUpdateHooks, notifyViewHook)
	case boil.BeforeDeleteHook:
		notifyViewBeforeDeleteHooks = append(notifyViewBeforeDeleteHooks, notifyViewHook)
	case boil.BeforeUpsertHook:
		notifyViewBeforeUpsertHooks = append(notifyViewBeforeUpsertHooks, notifyViewHook)
	case boil.AfterInsertHook:
		notifyViewAfterInsertHooks = append(notifyViewAfterInsertHooks, notifyViewHook)
	case boil.AfterSelectHook:
		notifyViewAfterSelectHooks = append(notifyViewAfterSelectHooks, notifyViewHook)
	case boil.AfterUpdateHook:
		notifyViewAfterUpdateHooks = append(notifyViewAfterUpdateHooks, notifyViewHook)
	case boil.AfterDeleteHook:
		notifyViewAfterDeleteHooks = append(notifyViewAfterDeleteHooks, notifyViewHook)
	case boil.AfterUpsertHook:
		notifyViewAfterUpsertHooks = append(notifyViewAfterUpsertHooks, notifyViewHook)
	}
}

// OneP returns a single notifyView record from the query, and panics on error.
func (q notifyViewQuery) OneP() *NotifyView {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single notifyView record from the query.
func (q notifyViewQuery) One() (*NotifyView, error) {
	o := &NotifyView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "public: failed to execute a one query for notify_view")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all NotifyView records from the query, and panics on error.
func (q notifyViewQuery) AllP() NotifyViewSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all NotifyView records from the query.
func (q notifyViewQuery) All() (NotifyViewSlice, error) {
	var o NotifyViewSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "public: failed to assign all query results to NotifyView slice")
	}

	if len(notifyViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all NotifyView records in the query, and panics on error.
func (q notifyViewQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all NotifyView records in the query.
func (q notifyViewQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "public: failed to count notify_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q notifyViewQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q notifyViewQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "public: failed to check if notify_view exists")
	}

	return count > 0, nil
}

// NotifyViewsG retrieves all records.
func NotifyViewsG(mods ...qm.QueryMod) notifyViewQuery {
	return NotifyViews(boil.GetDB(), mods...)
}

// NotifyViews retrieves all the records using an executor.
func NotifyViews(exec boil.Executor, mods ...qm.QueryMod) notifyViewQuery {
	mods = append(mods, qm.From("\"notify_view\""))
	return notifyViewQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *NotifyView) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *NotifyView) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *NotifyView) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *NotifyView) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("public: no notify_view provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(notifyViewColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	notifyViewInsertCacheMut.RLock()
	cache, cached := notifyViewInsertCache[key]
	notifyViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			notifyViewColumns,
			notifyViewColumnsWithDefault,
			notifyViewColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(notifyViewType, notifyViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(notifyViewType, notifyViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"notify_view\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"notify_view\" DEFAULT VALUES"
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
		return errors.Wrap(err, "public: unable to insert into notify_view")
	}

	if !cached {
		notifyViewInsertCacheMut.Lock()
		notifyViewInsertCache[key] = cache
		notifyViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}
