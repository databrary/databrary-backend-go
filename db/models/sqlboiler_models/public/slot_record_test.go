// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package public

import (
	"bytes"
	"github.com/databrary/databrary-backend-go/db/models/custom_types"
	"github.com/databrary/sqlboiler/boil"
	"github.com/databrary/sqlboiler/randomize"
	"github.com/databrary/sqlboiler/strmangle"
	"github.com/pmezard/go-difflib/difflib"
	"os"
	"os/exec"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func testSlotRecords(t *testing.T) {
	t.Parallel()

	query := SlotRecords(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testSlotRecordsLive(t *testing.T) {
	all, err := SlotRecords(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all SlotRecords err: ", err)
	}
	tx, err := dbMain.liveTestDbConn.Begin()
	if err != nil {
		t.Fatalf("failed to begin transaction: ", err)
	}
	for _, v := range all {
		err := v.Insert(tx)
		if err != nil {
			t.Fatalf("failed to failed to insert %s because of %s", v, err)
		}

	}
	err = tx.Commit()
	if err != nil {
		t.Fatalf("failed to commit transaction: ", err)
	}
	bf := &bytes.Buffer{}
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM slot_record) TO STDOUT" -d `, dbMain.DbName)
	dumpCmd.Env = append(os.Environ(), dbMain.pgEnv()...)
	dumpCmd.Stdout = bf
	err = dumpCmd.Start()
	if err != nil {
		t.Fatalf("failed to start dump from live db because of %s", err)
	}
	dumpCmd.Wait()
	if err != nil {
		t.Fatalf("failed to wait dump from live db because of %s", err)
	}
	bg := &bytes.Buffer{}
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM slot_record) TO STDOUT" -d `, dbMain.LiveTestDBName)
	dumpCmd.Env = append(os.Environ(), dbMain.pgEnv()...)
	dumpCmd.Stdout = bg
	err = dumpCmd.Start()
	if err != nil {
		t.Fatalf("failed to start dump from test db because of %s", err)
	}
	dumpCmd.Wait()
	if err != nil {
		t.Fatalf("failed to wait dump from test db because of %s", err)
	}
	bfslice := sort.StringSlice(difflib.SplitLines(bf.String()))
	gfslice := sort.StringSlice(difflib.SplitLines(bg.String()))
	bfslice.Sort()
	gfslice.Sort()
	diff := difflib.ContextDiff{
		A:        bfslice,
		B:        gfslice,
		FromFile: "databrary",
		ToFile:   "test",
		Context:  1,
	}
	result, _ := difflib.GetContextDiffString(diff)
	if len(result) > 0 {
		t.Fatalf("SlotRecordsLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testSlotRecordsDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = slotRecord.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlotRecordsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SlotRecords(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlotRecordsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SlotRecordSlice{slotRecord}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testSlotRecordsExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := SlotRecordExists(tx, slotRecord.Container, slotRecord.Segment, slotRecord.Record)
	if err != nil {
		t.Errorf("Unable to check if SlotRecord exists: %s", err)
	}
	if !e {
		t.Errorf("Expected SlotRecordExistsG to return true, but got false.")
	}
}

func testSlotRecordsFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	slotRecordFound, err := FindSlotRecord(tx, slotRecord.Container, slotRecord.Segment, slotRecord.Record)
	if err != nil {
		t.Error(err)
	}

	if slotRecordFound == nil {
		t.Error("want a record, got nil")
	}
}

func testSlotRecordsBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = SlotRecords(tx).Bind(slotRecord); err != nil {
		t.Error(err)
	}
}

func testSlotRecordsOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := SlotRecords(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testSlotRecordsAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecordOne := &SlotRecord{}
	slotRecordTwo := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecordOne, slotRecordDBTypes, false, slotRecordColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}
	if err = randomize.Struct(seed, slotRecordTwo, slotRecordDBTypes, false, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecordOne.Segment = custom_types.SegmentRandom()
	slotRecordTwo.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecordOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = slotRecordTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SlotRecords(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testSlotRecordsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecordOne := &SlotRecord{}
	slotRecordTwo := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecordOne, slotRecordDBTypes, false, slotRecordColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}
	if err = randomize.Struct(seed, slotRecordTwo, slotRecordDBTypes, false, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecordOne.Segment = custom_types.SegmentRandom()
	slotRecordTwo.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecordOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = slotRecordTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func slotRecordBeforeInsertHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordAfterInsertHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordAfterSelectHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordBeforeUpdateHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordAfterUpdateHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordBeforeDeleteHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordAfterDeleteHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordBeforeUpsertHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func slotRecordAfterUpsertHook(e boil.Executor, o *SlotRecord) error {
	*o = SlotRecord{}
	return nil
}

func testSlotRecordsHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	empty := &SlotRecord{}

	AddSlotRecordHook(boil.BeforeInsertHook, slotRecordBeforeInsertHook)
	if err = slotRecord.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordBeforeInsertHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.AfterInsertHook, slotRecordAfterInsertHook)
	if err = slotRecord.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordAfterInsertHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.AfterSelectHook, slotRecordAfterSelectHook)
	if err = slotRecord.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordAfterSelectHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.BeforeUpdateHook, slotRecordBeforeUpdateHook)
	if err = slotRecord.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordBeforeUpdateHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.AfterUpdateHook, slotRecordAfterUpdateHook)
	if err = slotRecord.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordAfterUpdateHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.BeforeDeleteHook, slotRecordBeforeDeleteHook)
	if err = slotRecord.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordBeforeDeleteHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.AfterDeleteHook, slotRecordAfterDeleteHook)
	if err = slotRecord.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordAfterDeleteHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.BeforeUpsertHook, slotRecordBeforeUpsertHook)
	if err = slotRecord.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordBeforeUpsertHooks = []SlotRecordHook{}

	AddSlotRecordHook(boil.AfterUpsertHook, slotRecordAfterUpsertHook)
	if err = slotRecord.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(slotRecord, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", slotRecord)
	}
	slotRecordAfterUpsertHooks = []SlotRecordHook{}
}
func testSlotRecordsInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlotRecordsInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx, slotRecordColumns...); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testSlotRecordToOneContainerUsingContainer(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Container
	var local SlotRecord

	foreignBlacklist := containerColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, containerDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	localBlacklist := slotRecordColumnsWithDefault
	localBlacklist = append(localBlacklist, slotRecordColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, slotRecordDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}
	local.Segment = custom_types.SegmentRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Container = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.ContainerByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SlotRecordSlice{&local}
	if err = local.L.LoadContainer(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Container == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Container = nil
	if err = local.L.LoadContainer(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Container == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSlotRecordToOneRecordUsingRecord(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Record
	var local SlotRecord

	foreignBlacklist := recordColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, recordDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := slotRecordColumnsWithDefault
	localBlacklist = append(localBlacklist, slotRecordColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, slotRecordDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}
	local.Segment = custom_types.SegmentRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Record = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.RecordByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := SlotRecordSlice{&local}
	if err = local.L.LoadRecord(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Record == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Record = nil
	if err = local.L.LoadRecord(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Record == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testSlotRecordToOneSetOpContainerUsingContainer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a SlotRecord
	var b, c Container

	foreignBlacklist := strmangle.SetComplement(containerPrimaryKeyColumns, containerColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, containerDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, containerDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(slotRecordPrimaryKeyColumns, slotRecordColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, slotRecordColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, slotRecordDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}
	a.Segment = custom_types.SegmentRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Container{&b, &c} {
		err = a.SetContainer(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Container != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SlotRecords[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Container != x.ID {
			t.Error("foreign key was wrong value", a.Container)
		}

		if exists, err := SlotRecordExists(tx, a.Container, a.Segment, a.Record); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testSlotRecordToOneSetOpRecordUsingRecord(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a SlotRecord
	var b, c Record

	foreignBlacklist := strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(slotRecordPrimaryKeyColumns, slotRecordColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, slotRecordColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, slotRecordDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}
	a.Segment = custom_types.SegmentRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Record{&b, &c} {
		err = a.SetRecord(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Record != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.SlotRecords[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Record != x.ID {
			t.Error("foreign key was wrong value", a.Record)
		}

		if exists, err := SlotRecordExists(tx, a.Container, a.Segment, a.Record); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testSlotRecordsReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = slotRecord.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testSlotRecordsReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := SlotRecordSlice{slotRecord}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testSlotRecordsSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := SlotRecords(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	slotRecordDBTypes = map[string]string{`Container`: `integer`, `Record`: `integer`, `Segment`: `USER-DEFINED`}
	_                 = bytes.MinRead
)

func testSlotRecordsUpdate(t *testing.T) {
	t.Parallel()

	if len(slotRecordColumns) == len(slotRecordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := slotRecordColumnsWithDefault
	blacklist = append(blacklist, slotRecordColumnsWithCustom...)

	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	if err = slotRecord.Update(tx); err != nil {
		t.Error(err)
	}
}

func testSlotRecordsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(slotRecordColumns) == len(slotRecordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := slotRecordPrimaryKeyColumns
	blacklist = append(blacklist, slotRecordColumnsWithCustom...)

	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(slotRecordColumns, slotRecordPrimaryKeyColumns) {
		fields = slotRecordColumns
	} else {
		fields = strmangle.SetComplement(
			slotRecordColumns,
			slotRecordPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(slotRecord))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := SlotRecordSlice{slotRecord}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testSlotRecordsUpsert(t *testing.T) {
	t.Parallel()

	if len(slotRecordColumns) == len(slotRecordPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	slotRecord := &SlotRecord{}
	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, true, slotRecordColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = slotRecord.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert SlotRecord: %s", err)
	}

	count, err := SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := slotRecordPrimaryKeyColumns

	blacklist = append(blacklist, slotRecordColumnsWithCustom...)

	if err = randomize.Struct(seed, slotRecord, slotRecordDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize SlotRecord struct: %s", err)
	}

	slotRecord.Segment = custom_types.SegmentRandom()

	if err = slotRecord.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert SlotRecord: %s", err)
	}

	count, err = SlotRecords(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
