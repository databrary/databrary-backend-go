// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package audit

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

// VolumeCitation is an object representing the database table.
type VolumeCitation struct {
	AuditTime   time.Time           `db:"audit_time" json:"volumeCitation_audit_time"`
	AuditUser   int                 `db:"audit_user" json:"volumeCitation_audit_user"`
	AuditIP     custom_types.Inet   `db:"audit_ip" json:"volumeCitation_audit_ip"`
	AuditAction custom_types.Action `db:"audit_action" json:"volumeCitation_audit_action"`
	Volume      int                 `db:"volume" json:"volumeCitation_volume"`
	Head        string              `db:"head" json:"volumeCitation_head"`
	URL         null.String         `db:"url" json:"volumeCitation_url,omitempty"`
	Year        null.Int16          `db:"year" json:"volumeCitation_year,omitempty"`

	R *volumeCitationR `db:"-" json:"-"`
	L volumeCitationL  `db:"-" json:"-"`
}

// volumeCitationR is where relationships are stored.
type volumeCitationR struct {
}

// volumeCitationL is where Load methods for each relationship are stored.
type volumeCitationL struct{}

var (
	volumeCitationColumns               = []string{"audit_time", "audit_user", "audit_ip", "audit_action", "volume", "head", "url", "year"}
	volumeCitationColumnsWithoutDefault = []string{"audit_user", "audit_ip", "audit_action", "volume", "head", "url", "year"}
	volumeCitationColumnsWithDefault    = []string{"audit_time"}
	volumeCitationColumnsWithCustom     = []string{"audit_ip", "audit_action"}
)

type (
	// VolumeCitationSlice is an alias for a slice of pointers to VolumeCitation.
	// This should generally be used opposed to []VolumeCitation.
	VolumeCitationSlice []*VolumeCitation
	// VolumeCitationHook is the signature for custom VolumeCitation hook methods
	VolumeCitationHook func(boil.Executor, *VolumeCitation) error

	volumeCitationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	volumeCitationType    = reflect.TypeOf(&VolumeCitation{})
	volumeCitationMapping = queries.MakeStructMapping(volumeCitationType)

	volumeCitationInsertCacheMut sync.RWMutex
	volumeCitationInsertCache    = make(map[string]insertCache)
	volumeCitationUpdateCacheMut sync.RWMutex
	volumeCitationUpdateCache    = make(map[string]updateCache)
	volumeCitationUpsertCacheMut sync.RWMutex
	volumeCitationUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var volumeCitationBeforeInsertHooks []VolumeCitationHook
var volumeCitationBeforeUpdateHooks []VolumeCitationHook
var volumeCitationBeforeDeleteHooks []VolumeCitationHook
var volumeCitationBeforeUpsertHooks []VolumeCitationHook

var volumeCitationAfterInsertHooks []VolumeCitationHook
var volumeCitationAfterSelectHooks []VolumeCitationHook
var volumeCitationAfterUpdateHooks []VolumeCitationHook
var volumeCitationAfterDeleteHooks []VolumeCitationHook
var volumeCitationAfterUpsertHooks []VolumeCitationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *VolumeCitation) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *VolumeCitation) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *VolumeCitation) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *VolumeCitation) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *VolumeCitation) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *VolumeCitation) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *VolumeCitation) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *VolumeCitation) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *VolumeCitation) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range volumeCitationAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddVolumeCitationHook registers your hook function for all future operations.
func AddVolumeCitationHook(hookPoint boil.HookPoint, volumeCitationHook VolumeCitationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		volumeCitationBeforeInsertHooks = append(volumeCitationBeforeInsertHooks, volumeCitationHook)
	case boil.BeforeUpdateHook:
		volumeCitationBeforeUpdateHooks = append(volumeCitationBeforeUpdateHooks, volumeCitationHook)
	case boil.BeforeDeleteHook:
		volumeCitationBeforeDeleteHooks = append(volumeCitationBeforeDeleteHooks, volumeCitationHook)
	case boil.BeforeUpsertHook:
		volumeCitationBeforeUpsertHooks = append(volumeCitationBeforeUpsertHooks, volumeCitationHook)
	case boil.AfterInsertHook:
		volumeCitationAfterInsertHooks = append(volumeCitationAfterInsertHooks, volumeCitationHook)
	case boil.AfterSelectHook:
		volumeCitationAfterSelectHooks = append(volumeCitationAfterSelectHooks, volumeCitationHook)
	case boil.AfterUpdateHook:
		volumeCitationAfterUpdateHooks = append(volumeCitationAfterUpdateHooks, volumeCitationHook)
	case boil.AfterDeleteHook:
		volumeCitationAfterDeleteHooks = append(volumeCitationAfterDeleteHooks, volumeCitationHook)
	case boil.AfterUpsertHook:
		volumeCitationAfterUpsertHooks = append(volumeCitationAfterUpsertHooks, volumeCitationHook)
	}
}

// OneP returns a single volumeCitation record from the query, and panics on error.
func (q volumeCitationQuery) OneP() *VolumeCitation {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single volumeCitation record from the query.
func (q volumeCitationQuery) One() (*VolumeCitation, error) {
	o := &VolumeCitation{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for volume_citation")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all VolumeCitation records from the query, and panics on error.
func (q volumeCitationQuery) AllP() VolumeCitationSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all VolumeCitation records from the query.
func (q volumeCitationQuery) All() (VolumeCitationSlice, error) {
	var o VolumeCitationSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to VolumeCitation slice")
	}

	if len(volumeCitationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all VolumeCitation records in the query, and panics on error.
func (q volumeCitationQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all VolumeCitation records in the query.
func (q volumeCitationQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count volume_citation rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q volumeCitationQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q volumeCitationQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if volume_citation exists")
	}

	return count > 0, nil
}

// VolumeCitationsG retrieves all records.
func VolumeCitationsG(mods ...qm.QueryMod) volumeCitationQuery {
	return VolumeCitations(boil.GetDB(), mods...)
}

// VolumeCitations retrieves all the records using an executor.
func VolumeCitations(exec boil.Executor, mods ...qm.QueryMod) volumeCitationQuery {
	mods = append(mods, qm.From("\"audit\".\"volume_citation\""))
	return volumeCitationQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *VolumeCitation) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *VolumeCitation) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *VolumeCitation) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *VolumeCitation) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no volume_citation provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(volumeCitationColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	volumeCitationInsertCacheMut.RLock()
	cache, cached := volumeCitationInsertCache[key]
	volumeCitationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			volumeCitationColumns,
			volumeCitationColumnsWithDefault,
			volumeCitationColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(volumeCitationType, volumeCitationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(volumeCitationType, volumeCitationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit\".\"volume_citation\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit\".\"volume_citation\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into volume_citation")
	}

	if !cached {
		volumeCitationInsertCacheMut.Lock()
		volumeCitationInsertCache[key] = cache
		volumeCitationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}
