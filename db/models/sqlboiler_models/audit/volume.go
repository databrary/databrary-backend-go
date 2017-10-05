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

// Volume is an object representing the database table.
type Volume struct {
	AuditTime   time.Time           `db:"audit_time" json:"volume_audit_time"`
	AuditUser   int                 `db:"audit_user" json:"volume_audit_user"`
	AuditIP     custom_types.Inet   `db:"audit_ip" json:"volume_audit_ip"`
	AuditAction custom_types.Action `db:"audit_action" json:"volume_audit_action"`
	ID          int                 `db:"id" json:"volume_id"`
	Name        null.String         `db:"name" json:"volume_name,omitempty"`
	Body        null.String         `db:"body" json:"volume_body,omitempty"`
	Alias       null.String         `db:"alias" json:"volume_alias,omitempty"`
	Doi         null.String         `db:"doi" json:"volume_doi,omitempty"`

	R *volumeR `db:"-" json:"-"`
	L volumeL  `db:"-" json:"-"`
}

// volumeR is where relationships are stored.
type volumeR struct {
}

// volumeL is where Load methods for each relationship are stored.
type volumeL struct{}

var (
	volumeColumns               = []string{"audit_time", "audit_user", "audit_ip", "audit_action", "id", "name", "body", "alias", "doi"}
	volumeColumnsWithoutDefault = []string{"audit_user", "audit_ip", "audit_action", "id", "name", "body", "alias", "doi"}
	volumeColumnsWithDefault    = []string{"audit_time"}
	volumeColumnsWithCustom     = []string{"audit_ip", "audit_action"}
)

type (
	// VolumeSlice is an alias for a slice of pointers to Volume.
	// This should generally be used opposed to []Volume.
	VolumeSlice []*Volume
	// VolumeHook is the signature for custom Volume hook methods
	VolumeHook func(boil.Executor, *Volume) error

	volumeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	volumeType    = reflect.TypeOf(&Volume{})
	volumeMapping = queries.MakeStructMapping(volumeType)

	volumeInsertCacheMut sync.RWMutex
	volumeInsertCache    = make(map[string]insertCache)
	volumeUpdateCacheMut sync.RWMutex
	volumeUpdateCache    = make(map[string]updateCache)
	volumeUpsertCacheMut sync.RWMutex
	volumeUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var volumeBeforeInsertHooks []VolumeHook
var volumeBeforeUpdateHooks []VolumeHook
var volumeBeforeDeleteHooks []VolumeHook
var volumeBeforeUpsertHooks []VolumeHook

var volumeAfterInsertHooks []VolumeHook
var volumeAfterSelectHooks []VolumeHook
var volumeAfterUpdateHooks []VolumeHook
var volumeAfterDeleteHooks []VolumeHook
var volumeAfterUpsertHooks []VolumeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Volume) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Volume) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Volume) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Volume) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Volume) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Volume) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Volume) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Volume) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Volume) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVolumeHook registers your hook function for all future operations.
func AddVolumeHook(hookPoint boil.HookPoint, volumeHook VolumeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		volumeBeforeInsertHooks = append(volumeBeforeInsertHooks, volumeHook)
	case boil.BeforeUpdateHook:
		volumeBeforeUpdateHooks = append(volumeBeforeUpdateHooks, volumeHook)
	case boil.BeforeDeleteHook:
		volumeBeforeDeleteHooks = append(volumeBeforeDeleteHooks, volumeHook)
	case boil.BeforeUpsertHook:
		volumeBeforeUpsertHooks = append(volumeBeforeUpsertHooks, volumeHook)
	case boil.AfterInsertHook:
		volumeAfterInsertHooks = append(volumeAfterInsertHooks, volumeHook)
	case boil.AfterSelectHook:
		volumeAfterSelectHooks = append(volumeAfterSelectHooks, volumeHook)
	case boil.AfterUpdateHook:
		volumeAfterUpdateHooks = append(volumeAfterUpdateHooks, volumeHook)
	case boil.AfterDeleteHook:
		volumeAfterDeleteHooks = append(volumeAfterDeleteHooks, volumeHook)
	case boil.AfterUpsertHook:
		volumeAfterUpsertHooks = append(volumeAfterUpsertHooks, volumeHook)
	}
}

// OneP returns a single volume record from the query, and panics on error.
func (q volumeQuery) OneP() *Volume {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single volume record from the query.
func (q volumeQuery) One() (*Volume, error) {
	o := &Volume{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for volume")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Volume records from the query, and panics on error.
func (q volumeQuery) AllP() VolumeSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Volume records from the query.
func (q volumeQuery) All() (VolumeSlice, error) {
	var o VolumeSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Volume slice")
	}

	if len(volumeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Volume records in the query, and panics on error.
func (q volumeQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Volume records in the query.
func (q volumeQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count volume rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q volumeQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q volumeQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if volume exists")
	}

	return count > 0, nil
}

// VolumesG retrieves all records.
func VolumesG(mods ...qm.QueryMod) volumeQuery {
	return Volumes(boil.GetDB(), mods...)
}

// Volumes retrieves all the records using an executor.
func Volumes(exec boil.Executor, mods ...qm.QueryMod) volumeQuery {
	mods = append(mods, qm.From("\"audit\".\"volume\""))
	return volumeQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Volume) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Volume) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Volume) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Volume) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no volume provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	volumeInsertCacheMut.RLock()
	cache, cached := volumeInsertCache[key]
	volumeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			volumeColumns,
			volumeColumnsWithDefault,
			volumeColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(volumeType, volumeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(volumeType, volumeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit\".\"volume\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit\".\"volume\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into volume")
	}

	if !cached {
		volumeInsertCacheMut.Lock()
		volumeInsertCache[key] = cache
		volumeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}
