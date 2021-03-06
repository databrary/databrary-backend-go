// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package audit

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/databrary/databrary-backend-go/db/models/custom_types"
	"github.com/databrary/sqlboiler/boil"
	"github.com/databrary/sqlboiler/queries"
	"github.com/databrary/sqlboiler/queries/qm"
	"github.com/databrary/sqlboiler/strmangle"
	"github.com/pkg/errors"
)

// Slot is an object representing the database table.
type Slot struct {
	AuditTime   time.Time            `db:"audit_time" json:"slot_audit_time"`
	AuditUser   int                  `db:"audit_user" json:"slot_audit_user"`
	AuditIP     custom_types.Inet    `db:"audit_ip" json:"slot_audit_ip"`
	AuditAction custom_types.Action  `db:"audit_action" json:"slot_audit_action"`
	Container   int                  `db:"container" json:"slot_container"`
	Segment     custom_types.Segment `db:"segment" json:"slot_segment"`

	R *slotR `db:"-" json:"-"`
	L slotL  `db:"-" json:"-"`
}

// slotR is where relationships are stored.
type slotR struct {
}

// slotL is where Load methods for each relationship are stored.
type slotL struct{}

var (
	slotColumns               = []string{"audit_time", "audit_user", "audit_ip", "audit_action", "container", "segment"}
	slotColumnsWithoutDefault = []string{"audit_user", "audit_ip", "audit_action", "container", "segment"}
	slotColumnsWithDefault    = []string{"audit_time"}
	slotColumnsWithCustom     = []string{"audit_ip", "audit_action", "segment"}
)

type (
	// SlotSlice is an alias for a slice of pointers to Slot.
	// This should generally be used opposed to []Slot.
	SlotSlice []*Slot
	// SlotHook is the signature for custom Slot hook methods
	SlotHook func(boil.Executor, *Slot) error

	slotQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	slotType    = reflect.TypeOf(&Slot{})
	slotMapping = queries.MakeStructMapping(slotType)

	slotInsertCacheMut sync.RWMutex
	slotInsertCache    = make(map[string]insertCache)
	slotUpdateCacheMut sync.RWMutex
	slotUpdateCache    = make(map[string]updateCache)
	slotUpsertCacheMut sync.RWMutex
	slotUpsertCache    = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var slotBeforeInsertHooks []SlotHook
var slotBeforeUpdateHooks []SlotHook
var slotBeforeDeleteHooks []SlotHook
var slotBeforeUpsertHooks []SlotHook

var slotAfterInsertHooks []SlotHook
var slotAfterSelectHooks []SlotHook
var slotAfterUpdateHooks []SlotHook
var slotAfterDeleteHooks []SlotHook
var slotAfterUpsertHooks []SlotHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Slot) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Slot) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range slotBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Slot) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range slotBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Slot) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Slot) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Slot) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range slotAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Slot) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range slotAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Slot) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range slotAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Slot) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range slotAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddSlotHook registers your hook function for all future operations.
func AddSlotHook(hookPoint boil.HookPoint, slotHook SlotHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		slotBeforeInsertHooks = append(slotBeforeInsertHooks, slotHook)
	case boil.BeforeUpdateHook:
		slotBeforeUpdateHooks = append(slotBeforeUpdateHooks, slotHook)
	case boil.BeforeDeleteHook:
		slotBeforeDeleteHooks = append(slotBeforeDeleteHooks, slotHook)
	case boil.BeforeUpsertHook:
		slotBeforeUpsertHooks = append(slotBeforeUpsertHooks, slotHook)
	case boil.AfterInsertHook:
		slotAfterInsertHooks = append(slotAfterInsertHooks, slotHook)
	case boil.AfterSelectHook:
		slotAfterSelectHooks = append(slotAfterSelectHooks, slotHook)
	case boil.AfterUpdateHook:
		slotAfterUpdateHooks = append(slotAfterUpdateHooks, slotHook)
	case boil.AfterDeleteHook:
		slotAfterDeleteHooks = append(slotAfterDeleteHooks, slotHook)
	case boil.AfterUpsertHook:
		slotAfterUpsertHooks = append(slotAfterUpsertHooks, slotHook)
	}
}

// OneP returns a single slot record from the query, and panics on error.
func (q slotQuery) OneP() *Slot {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single slot record from the query.
func (q slotQuery) One() (*Slot, error) {
	o := &Slot{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "audit: failed to execute a one query for slot")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Slot records from the query, and panics on error.
func (q slotQuery) AllP() SlotSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Slot records from the query.
func (q slotQuery) All() (SlotSlice, error) {
	var o SlotSlice

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "audit: failed to assign all query results to Slot slice")
	}

	if len(slotAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Slot records in the query, and panics on error.
func (q slotQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Slot records in the query.
func (q slotQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "audit: failed to count slot rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q slotQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q slotQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "audit: failed to check if slot exists")
	}

	return count > 0, nil
}

// SlotsG retrieves all records.
func SlotsG(mods ...qm.QueryMod) slotQuery {
	return Slots(boil.GetDB(), mods...)
}

// Slots retrieves all the records using an executor.
func Slots(exec boil.Executor, mods ...qm.QueryMod) slotQuery {
	mods = append(mods, qm.From("\"audit\".\"slot\""))
	return slotQuery{NewQuery(exec, mods...)}
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Slot) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Slot) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Slot) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Slot) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("audit: no slot provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(slotColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	slotInsertCacheMut.RLock()
	cache, cached := slotInsertCache[key]
	slotInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			slotColumns,
			slotColumnsWithDefault,
			slotColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(slotType, slotMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(slotType, slotMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"audit\".\"slot\" (\"%s\") VALUES (%s)", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"audit\".\"slot\" DEFAULT VALUES"
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
		return errors.Wrap(err, "audit: unable to insert into slot")
	}

	if !cached {
		slotInsertCacheMut.Lock()
		slotInsertCache[key] = cache
		slotInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}
