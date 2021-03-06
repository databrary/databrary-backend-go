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

// VolumeAccessView is an object representing the database view.
type VolumeAccessView struct {
	Volume null.Int                    `db:"volume" json:"volumeAccessView_volume,omitempty"`
	Party  null.Int                    `db:"party" json:"volumeAccessView_party,omitempty"`
	Access custom_types.NullPermission `db:"access" json:"volumeAccessView_access,omitempty"`

	R *volumeAccessViewR `db:"-" json:"-"`
	L volumeAccessViewL  `db:"-" json:"-"`
}

// volumeAccessViewR is where relationships are stored.
type volumeAccessViewR struct {
}

// volumeAccessViewL is where Load methods for each relationship are stored.
type volumeAccessViewL struct{}

var (
	volumeAccessViewColumns               = []string{"volume", "party", "access"}
	volumeAccessViewColumnsWithoutDefault = []string{"volume", "party", "access"}
	volumeAccessViewColumnsWithDefault    = []string{}
	volumeAccessViewColumnsWithCustom     = []string{"access"}
)

type (
	// VolumeAccessViewSlice is an alias for a slice of pointers to VolumeAccessView.
	// This should generally be used opposed to []VolumeAccessView.
	VolumeAccessViewSlice []*VolumeAccessView
	// VolumeAccessViewHook is the signature for custom VolumeAccessView hook methods
	VolumeAccessViewHook func(boil.Executor, *VolumeAccessView) error

	volumeAccessViewQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	volumeAccessViewType    = reflect.TypeOf(&VolumeAccessView{})
	volumeAccessViewMapping = queries.MakeStructMapping(volumeAccessViewType)

	volumeAccessViewInsertCacheMut sync.RWMutex
	volumeAccessViewInsertCache    = make(map[string]insertCache)
	volumeAccessViewUpdateCacheMut sync.RWMutex
	volumeAccessViewUpdateCache    = make(map[string]updateCache)
	volumeAccessViewUpsertCacheMut sync.RWMutex
	volumeAccessViewUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var volumeAccessViewBeforeInsertHooks []VolumeAccessViewHook
var volumeAccessViewBeforeUpdateHooks []VolumeAccessViewHook
var volumeAccessViewBeforeDeleteHooks []VolumeAccessViewHook
var volumeAccessViewBeforeUpsertHooks []VolumeAccessViewHook

var volumeAccessViewAfterInsertHooks []VolumeAccessViewHook
var volumeAccessViewAfterSelectHooks []VolumeAccessViewHook
var volumeAccessViewAfterUpdateHooks []VolumeAccessViewHook
var volumeAccessViewAfterDeleteHooks []VolumeAccessViewHook
var volumeAccessViewAfterUpsertHooks []VolumeAccessViewHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VolumeAccessView) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VolumeAccessView) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VolumeAccessView) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VolumeAccessView) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VolumeAccessView) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VolumeAccessView) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VolumeAccessView) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VolumeAccessView) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VolumeAccessView) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAccessViewAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVolumeAccessViewHook registers your hook function for all future operations.
func AddVolumeAccessViewHook(hookPoint boil.HookPoint, volumeAccessViewHook VolumeAccessViewHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		volumeAccessViewBeforeInsertHooks = append(volumeAccessViewBeforeInsertHooks, volumeAccessViewHook)
	case boil.BeforeUpdateHook:
		volumeAccessViewBeforeUpdateHooks = append(volumeAccessViewBeforeUpdateHooks, volumeAccessViewHook)
	case boil.BeforeDeleteHook:
		volumeAccessViewBeforeDeleteHooks = append(volumeAccessViewBeforeDeleteHooks, volumeAccessViewHook)
	case boil.BeforeUpsertHook:
		volumeAccessViewBeforeUpsertHooks = append(volumeAccessViewBeforeUpsertHooks, volumeAccessViewHook)
	case boil.AfterInsertHook:
		volumeAccessViewAfterInsertHooks = append(volumeAccessViewAfterInsertHooks, volumeAccessViewHook)
	case boil.AfterSelectHook:
		volumeAccessViewAfterSelectHooks = append(volumeAccessViewAfterSelectHooks, volumeAccessViewHook)
	case boil.AfterUpdateHook:
		volumeAccessViewAfterUpdateHooks = append(volumeAccessViewAfterUpdateHooks, volumeAccessViewHook)
	case boil.AfterDeleteHook:
		volumeAccessViewAfterDeleteHooks = append(volumeAccessViewAfterDeleteHooks, volumeAccessViewHook)
	case boil.AfterUpsertHook:
		volumeAccessViewAfterUpsertHooks = append(volumeAccessViewAfterUpsertHooks, volumeAccessViewHook)
	}
}

// OneP returns a single volumeAccessView record from the query, and panics on error.
func (q volumeAccessViewQuery) OneP() *VolumeAccessView {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single volumeAccessView record from the query.
func (q volumeAccessViewQuery) One() (*VolumeAccessView, error) {
	o := &VolumeAccessView{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "public: failed to execute a one query for volume_access_view")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all VolumeAccessView records from the query, and panics on error.
func (q volumeAccessViewQuery) AllP() VolumeAccessViewSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all VolumeAccessView records from the query.
func (q volumeAccessViewQuery) All() (VolumeAccessViewSlice, error) {
	var o VolumeAccessViewSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "public: failed to assign all query results to VolumeAccessView slice")
	}

	if len(volumeAccessViewAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all VolumeAccessView records in the query, and panics on error.
func (q volumeAccessViewQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all VolumeAccessView records in the query.
func (q volumeAccessViewQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "public: failed to count volume_access_view rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q volumeAccessViewQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q volumeAccessViewQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "public: failed to check if volume_access_view exists")
	}

	return count > 0, nil
}

// VolumeAccessViewsG retrieves all records.
func VolumeAccessViewsG(mods ...qm.QueryMod) volumeAccessViewQuery {
	return VolumeAccessViews(boil.GetDB(), mods...)
}

// VolumeAccessViews retrieves all the records using an executor.
func VolumeAccessViews(exec boil.Executor, mods ...qm.QueryMod) volumeAccessViewQuery {
	mods = append(mods, qm.From("\"volume_access_view\""))
	return volumeAccessViewQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VolumeAccessView) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *VolumeAccessView) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *VolumeAccessView) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *VolumeAccessView) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("public: no volume_access_view provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeAccessViewColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	volumeAccessViewInsertCacheMut.RLock()
	cache, cached := volumeAccessViewInsertCache[key]
	volumeAccessViewInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			volumeAccessViewColumns,
			volumeAccessViewColumnsWithDefault,
			volumeAccessViewColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(volumeAccessViewType, volumeAccessViewMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(volumeAccessViewType, volumeAccessViewMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"volume_access_view\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"volume_access_view\" DEFAULT VALUES"
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
		return errors.Wrap(err, "public: unable to insert into volume_access_view")
	}

	if !cached {
		volumeAccessViewInsertCacheMut.Lock()
		volumeAccessViewInsertCache[key] = cache
		volumeAccessViewInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}
