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
	"reflect"
	"strings"
	"sync"
	"time"
)

// SlotRecord is an object representing the database table.
type SlotRecord struct {
	AuditTime   time.Time            `db:"audit_time" json:"slotRecord_audit_time"`
	AuditUser   int                  `db:"audit_user" json:"slotRecord_audit_user"`
	AuditIP     custom_types.Inet    `db:"audit_ip" json:"slotRecord_audit_ip"`
	AuditAction custom_types.Action  `db:"audit_action" json:"slotRecord_audit_action"`
	Container   int                  `db:"container" json:"slotRecord_container"`
	Segment     custom_types.Segment `db:"segment" json:"slotRecord_segment"`
	Record      int                  `db:"record" json:"slotRecord_record"`

	R *slotRecordR `db:"-" json:"-"`
	L slotRecordL  `db:"-" json:"-"`
}

// slotRecordR is where relationships are stored.
type slotRecordR struct {
}

// slotRecordL is where Load methods for each relationship are stored.
type slotRecordL struct{}

var (
	slotRecordColumns               = []string{"audit_time", "audit_user", "audit_ip", "audit_action", "container", "segment", "record"}
	slotRecordColumnsWithoutDefault = []string{"audit_user", "audit_ip", "audit_action", "container", "segment", "record"}
	slotRecordColumnsWithDefault    = []string{"audit_time"}
	slotRecordColumnsWithCustom     = []string{"audit_ip", "audit_action", "segment"}
)

type (
	// SlotRecordSlice is an alias for a slice of pointers to SlotRecord.
	// This should generally be used opposed to []SlotRecord.
	SlotRecordSlice []*SlotRecord
	// SlotRecordHook is the signature for custom SlotRecord hook methods
	SlotRecordHook func(boil.Executor, *SlotRecord) error

	slotRecordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	slotRecordType    = reflect.TypeOf(&SlotRecord{})
	slotRecordMapping = queries.MakeStructMapping(slotRecordType)

	slotRecordInsertCacheMut sync.RWMutex
	slotRecordInsertCache    = make(map[string]insertCache)
	slotRecordUpdateCacheMut sync.RWMutex
	slotRecordUpdateCache    = make(map[string]updateCache)
	slotRecordUpsertCacheMut sync.RWMutex
	slotRecordUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var slotRecordBeforeInsertHooks []SlotRecordHook
var slotRecordBeforeUpdateHooks []SlotRecordHook
var slotRecordBeforeDeleteHooks []SlotRecordHook
var slotRecordBeforeUpsertHooks []SlotRecordHook

var slotRecordAfterInsertHooks []SlotRecordHook
var slotRecordAfterSelectHooks []SlotRecordHook
var slotRecordAfterUpdateHooks []SlotRecordHook
var slotRecordAfterDeleteHooks []SlotRecordHook
var slotRecordAfterUpsertHooks []SlotRecordHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *SlotRecord) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *SlotRecord) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *SlotRecord) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *SlotRecord) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *SlotRecord) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *SlotRecord) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *SlotRecord) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *SlotRecord) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *SlotRecord) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotRecordAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSlotRecordHook registers your hook function for all future operations.
func AddSlotRecordHook(hookPoint boil.HookPoint, slotRecordHook SlotRecordHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		slotRecordBeforeInsertHooks = append(slotRecordBeforeInsertHooks, slotRecordHook)
	case boil.BeforeUpdateHook:
		slotRecordBeforeUpdateHooks = append(slotRecordBeforeUpdateHooks, slotRecordHook)
	case boil.BeforeDeleteHook:
		slotRecordBeforeDeleteHooks = append(slotRecordBeforeDeleteHooks, slotRecordHook)
	case boil.BeforeUpsertHook:
		slotRecordBeforeUpsertHooks = append(slotRecordBeforeUpsertHooks, slotRecordHook)
	case boil.AfterInsertHook:
		slotRecordAfterInsertHooks = append(slotRecordAfterInsertHooks, slotRecordHook)
	case boil.AfterSelectHook:
		slotRecordAfterSelectHooks = append(slotRecordAfterSelectHooks, slotRecordHook)
	case boil.AfterUpdateHook:
		slotRecordAfterUpdateHooks = append(slotRecordAfterUpdateHooks, slotRecordHook)
	case boil.AfterDeleteHook:
		slotRecordAfterDeleteHooks = append(slotRecordAfterDeleteHooks, slotRecordHook)
	case boil.AfterUpsertHook:
		slotRecordAfterUpsertHooks = append(slotRecordAfterUpsertHooks, slotRecordHook)
	}
}

// OneP returns a single slotRecord record from the query, and panics on error.
func (q slotRecordQuery) OneP() *SlotRecord {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single slotRecord record from the query.
func (q slotRecordQuery) One() (*SlotRecord, error) {
	o := &SlotRecord{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for slot_record")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all SlotRecord records from the query, and panics on error.
func (q slotRecordQuery) AllP() SlotRecordSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all SlotRecord records from the query.
func (q slotRecordQuery) All() (SlotRecordSlice, error) {
	var o SlotRecordSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to SlotRecord slice")
	}

	if len(slotRecordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all SlotRecord records in the query, and panics on error.
func (q slotRecordQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all SlotRecord records in the query.
func (q slotRecordQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count slot_record rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q slotRecordQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q slotRecordQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if slot_record exists")
	}

	return count > 0, nil
}

// SlotRecordsG retrieves all records.
func SlotRecordsG(mods ...qm.QueryMod) slotRecordQuery {
	return SlotRecords(boil.GetDB(), mods...)
}

// SlotRecords retrieves all the records using an executor.
func SlotRecords(exec boil.Executor, mods ...qm.QueryMod) slotRecordQuery {
	mods = append(mods, qm.From("\"audit\".\"slot_record\""))
	return slotRecordQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *SlotRecord) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *SlotRecord) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *SlotRecord) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *SlotRecord) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no slot_record provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(slotRecordColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	slotRecordInsertCacheMut.RLock()
	cache, cached := slotRecordInsertCache[key]
	slotRecordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			slotRecordColumns,
			slotRecordColumnsWithDefault,
			slotRecordColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(slotRecordType, slotRecordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(slotRecordType, slotRecordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit\".\"slot_record\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit\".\"slot_record\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into slot_record")
	}

	if !cached {
		slotRecordInsertCacheMut.Lock()
		slotRecordInsertCache[key] = cache
		slotRecordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}
