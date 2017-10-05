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

func testMeasureDates(t *testing.T) {
	t.Parallel()

	query := MeasureDates(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMeasureDatesLive(t *testing.T) {
	all, err := MeasureDates(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all MeasureDates err: ", err)
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
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM measure_date) TO STDOUT" -d `, dbMain.DbName)
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
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM measure_date) TO STDOUT" -d `, dbMain.LiveTestDBName)
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
		t.Fatalf("MeasureDatesLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testMeasureDatesDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = measureDate.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasureDatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = MeasureDates(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasureDatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := MeasureDateSlice{measureDate}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasureDatesExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := MeasureDateExists(tx, measureDate.Record, measureDate.Metric)
	if err != nil {
		t.Errorf("Unable to check if MeasureDate exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MeasureDateExistsG to return true, but got false.")
	}
}

func testMeasureDatesFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	measureDateFound, err := FindMeasureDate(tx, measureDate.Record, measureDate.Metric)
	if err != nil {
		t.Error(err)
	}

	if measureDateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMeasureDatesBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = MeasureDates(tx).Bind(measureDate); err != nil {
		t.Error(err)
	}
}

func testMeasureDatesOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := MeasureDates(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMeasureDatesAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDateOne := &MeasureDate{}
	measureDateTwo := &MeasureDate{}
	if err = randomize.Struct(seed, measureDateOne, measureDateDBTypes, false, measureDateColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}
	if err = randomize.Struct(seed, measureDateTwo, measureDateDBTypes, false, measureDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = measureDateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := MeasureDates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMeasureDatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDateOne := &MeasureDate{}
	measureDateTwo := &MeasureDate{}
	if err = randomize.Struct(seed, measureDateOne, measureDateDBTypes, false, measureDateColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}
	if err = randomize.Struct(seed, measureDateTwo, measureDateDBTypes, false, measureDateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = measureDateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func measureDateBeforeInsertHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateAfterInsertHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateAfterSelectHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateBeforeUpdateHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateAfterUpdateHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateBeforeDeleteHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateAfterDeleteHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateBeforeUpsertHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func measureDateAfterUpsertHook(e boil.Executor, o *MeasureDate) error {
	*o = MeasureDate{}
	return nil
}

func testMeasureDatesHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	empty := &MeasureDate{}

	AddMeasureDateHook(boil.BeforeInsertHook, measureDateBeforeInsertHook)
	if err = measureDate.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", measureDate)
	}
	measureDateBeforeInsertHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.AfterInsertHook, measureDateAfterInsertHook)
	if err = measureDate.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", measureDate)
	}
	measureDateAfterInsertHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.AfterSelectHook, measureDateAfterSelectHook)
	if err = measureDate.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", measureDate)
	}
	measureDateAfterSelectHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.BeforeUpdateHook, measureDateBeforeUpdateHook)
	if err = measureDate.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", measureDate)
	}
	measureDateBeforeUpdateHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.AfterUpdateHook, measureDateAfterUpdateHook)
	if err = measureDate.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", measureDate)
	}
	measureDateAfterUpdateHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.BeforeDeleteHook, measureDateBeforeDeleteHook)
	if err = measureDate.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", measureDate)
	}
	measureDateBeforeDeleteHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.AfterDeleteHook, measureDateAfterDeleteHook)
	if err = measureDate.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", measureDate)
	}
	measureDateAfterDeleteHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.BeforeUpsertHook, measureDateBeforeUpsertHook)
	if err = measureDate.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", measureDate)
	}
	measureDateBeforeUpsertHooks = []MeasureDateHook{}

	AddMeasureDateHook(boil.AfterUpsertHook, measureDateAfterUpsertHook)
	if err = measureDate.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureDate, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", measureDate)
	}
	measureDateAfterUpsertHooks = []MeasureDateHook{}
}
func testMeasureDatesInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeasureDatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx, measureDateColumns...); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeasureDateToOneMetricUsingMetric(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Metric
	var local MeasureDate

	foreignBlacklist := metricColumnsWithDefault
	foreignBlacklist = append(foreignBlacklist, metricColumnsWithCustom...)

	if err := randomize.Struct(seed, &foreign, metricDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Metric struct: %s", err)
	}
	foreign.Release = custom_types.NullReleaseRandom()
	foreign.Type = custom_types.DataTypeRandom()

	localBlacklist := measureDateColumnsWithDefault
	if err := randomize.Struct(seed, &local, measureDateDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Metric = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.MetricByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := MeasureDateSlice{&local}
	if err = local.L.LoadMetric(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Metric == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Metric = nil
	if err = local.L.LoadMetric(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Metric == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testMeasureDateToOneRecordUsingRecord(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Record
	var local MeasureDate

	foreignBlacklist := recordColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, recordDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := measureDateColumnsWithDefault
	if err := randomize.Struct(seed, &local, measureDateDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

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

	slice := MeasureDateSlice{&local}
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

func testMeasureDateToOneSetOpMetricUsingMetric(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a MeasureDate
	var b, c Metric

	foreignBlacklist := strmangle.SetComplement(metricPrimaryKeyColumns, metricColumnsWithoutDefault)
	foreignBlacklist = append(foreignBlacklist, metricColumnsWithCustom...)

	if err := randomize.Struct(seed, &b, metricDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Metric struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, metricDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Metric struct: %s", err)
	}
	b.Release = custom_types.NullReleaseRandom()
	c.Release = custom_types.NullReleaseRandom()
	b.Type = custom_types.DataTypeRandom()
	c.Type = custom_types.DataTypeRandom()

	localBlacklist := strmangle.SetComplement(measureDatePrimaryKeyColumns, measureDateColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, measureDateDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Metric{&b, &c} {
		err = a.SetMetric(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Metric != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.MeasureDates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Metric != x.ID {
			t.Error("foreign key was wrong value", a.Metric)
		}

		if exists, err := MeasureDateExists(tx, a.Record, a.Metric); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testMeasureDateToOneSetOpRecordUsingRecord(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a MeasureDate
	var b, c Record

	foreignBlacklist := strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(measureDatePrimaryKeyColumns, measureDateColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, measureDateDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

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

		if x.R.MeasureDates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Record != x.ID {
			t.Error("foreign key was wrong value", a.Record)
		}

		if exists, err := MeasureDateExists(tx, a.Record, a.Metric); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testMeasureDatesReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = measureDate.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testMeasureDatesReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := MeasureDateSlice{measureDate}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testMeasureDatesSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := MeasureDates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	measureDateDBTypes = map[string]string{`Datum`: `date`, `Metric`: `integer`, `Record`: `integer`}
	_                  = bytes.MinRead
)

func testMeasureDatesUpdate(t *testing.T) {
	t.Parallel()

	if len(measureDateColumns) == len(measureDatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := measureDateColumnsWithDefault

	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	if err = measureDate.Update(tx); err != nil {
		t.Error(err)
	}
}

func testMeasureDatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(measureDateColumns) == len(measureDatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := measureDatePrimaryKeyColumns

	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(measureDateColumns, measureDatePrimaryKeyColumns) {
		fields = measureDateColumns
	} else {
		fields = strmangle.SetComplement(
			measureDateColumns,
			measureDatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(measureDate))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := MeasureDateSlice{measureDate}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testMeasureDatesUpsert(t *testing.T) {
	t.Parallel()

	if len(measureDateColumns) == len(measureDatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	measureDate := &MeasureDate{}
	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureDate.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert MeasureDate: %s", err)
	}

	count, err := MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := measureDatePrimaryKeyColumns

	if err = randomize.Struct(seed, measureDate, measureDateDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureDate struct: %s", err)
	}

	if err = measureDate.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert MeasureDate: %s", err)
	}

	count, err = MeasureDates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
