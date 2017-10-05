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

func testMeasureNumerics(t *testing.T) {
	t.Parallel()

	query := MeasureNumerics(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testMeasureNumericsLive(t *testing.T) {
	all, err := MeasureNumerics(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all MeasureNumerics err: ", err)
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
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM measure_numeric) TO STDOUT" -d `, dbMain.DbName)
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
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM measure_numeric) TO STDOUT" -d `, dbMain.LiveTestDBName)
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
		t.Fatalf("MeasureNumericsLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testMeasureNumericsDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = measureNumeric.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasureNumericsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = MeasureNumerics(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasureNumericsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := MeasureNumericSlice{measureNumeric}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testMeasureNumericsExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := MeasureNumericExists(tx, measureNumeric.Record, measureNumeric.Metric)
	if err != nil {
		t.Errorf("Unable to check if MeasureNumeric exists: %s", err)
	}
	if !e {
		t.Errorf("Expected MeasureNumericExistsG to return true, but got false.")
	}
}

func testMeasureNumericsFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	measureNumericFound, err := FindMeasureNumeric(tx, measureNumeric.Record, measureNumeric.Metric)
	if err != nil {
		t.Error(err)
	}

	if measureNumericFound == nil {
		t.Error("want a record, got nil")
	}
}

func testMeasureNumericsBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = MeasureNumerics(tx).Bind(measureNumeric); err != nil {
		t.Error(err)
	}
}

func testMeasureNumericsOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := MeasureNumerics(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testMeasureNumericsAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumericOne := &MeasureNumeric{}
	measureNumericTwo := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumericOne, measureNumericDBTypes, false, measureNumericColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}
	if err = randomize.Struct(seed, measureNumericTwo, measureNumericDBTypes, false, measureNumericColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumericOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = measureNumericTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := MeasureNumerics(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testMeasureNumericsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumericOne := &MeasureNumeric{}
	measureNumericTwo := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumericOne, measureNumericDBTypes, false, measureNumericColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}
	if err = randomize.Struct(seed, measureNumericTwo, measureNumericDBTypes, false, measureNumericColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumericOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = measureNumericTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func measureNumericBeforeInsertHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericAfterInsertHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericAfterSelectHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericBeforeUpdateHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericAfterUpdateHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericBeforeDeleteHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericAfterDeleteHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericBeforeUpsertHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func measureNumericAfterUpsertHook(e boil.Executor, o *MeasureNumeric) error {
	*o = MeasureNumeric{}
	return nil
}

func testMeasureNumericsHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	empty := &MeasureNumeric{}

	AddMeasureNumericHook(boil.BeforeInsertHook, measureNumericBeforeInsertHook)
	if err = measureNumeric.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericBeforeInsertHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.AfterInsertHook, measureNumericAfterInsertHook)
	if err = measureNumeric.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericAfterInsertHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.AfterSelectHook, measureNumericAfterSelectHook)
	if err = measureNumeric.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericAfterSelectHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.BeforeUpdateHook, measureNumericBeforeUpdateHook)
	if err = measureNumeric.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericBeforeUpdateHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.AfterUpdateHook, measureNumericAfterUpdateHook)
	if err = measureNumeric.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericAfterUpdateHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.BeforeDeleteHook, measureNumericBeforeDeleteHook)
	if err = measureNumeric.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericBeforeDeleteHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.AfterDeleteHook, measureNumericAfterDeleteHook)
	if err = measureNumeric.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericAfterDeleteHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.BeforeUpsertHook, measureNumericBeforeUpsertHook)
	if err = measureNumeric.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericBeforeUpsertHooks = []MeasureNumericHook{}

	AddMeasureNumericHook(boil.AfterUpsertHook, measureNumericAfterUpsertHook)
	if err = measureNumeric.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(measureNumeric, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", measureNumeric)
	}
	measureNumericAfterUpsertHooks = []MeasureNumericHook{}
}
func testMeasureNumericsInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeasureNumericsInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx, measureNumericColumns...); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testMeasureNumericToOneMetricUsingMetric(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Metric
	var local MeasureNumeric

	foreignBlacklist := metricColumnsWithDefault
	foreignBlacklist = append(foreignBlacklist, metricColumnsWithCustom...)

	if err := randomize.Struct(seed, &foreign, metricDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Metric struct: %s", err)
	}
	foreign.Release = custom_types.NullReleaseRandom()
	foreign.Type = custom_types.DataTypeRandom()

	localBlacklist := measureNumericColumnsWithDefault
	if err := randomize.Struct(seed, &local, measureNumericDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
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

	slice := MeasureNumericSlice{&local}
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

func testMeasureNumericToOneRecordUsingRecord(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Record
	var local MeasureNumeric

	foreignBlacklist := recordColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, recordDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := measureNumericColumnsWithDefault
	if err := randomize.Struct(seed, &local, measureNumericDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
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

	slice := MeasureNumericSlice{&local}
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

func testMeasureNumericToOneSetOpMetricUsingMetric(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a MeasureNumeric
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

	localBlacklist := strmangle.SetComplement(measureNumericPrimaryKeyColumns, measureNumericColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, measureNumericDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
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

		if x.R.MeasureNumerics[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Metric != x.ID {
			t.Error("foreign key was wrong value", a.Metric)
		}

		if exists, err := MeasureNumericExists(tx, a.Record, a.Metric); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testMeasureNumericToOneSetOpRecordUsingRecord(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a MeasureNumeric
	var b, c Record

	foreignBlacklist := strmangle.SetComplement(recordPrimaryKeyColumns, recordColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, recordDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Record struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(measureNumericPrimaryKeyColumns, measureNumericColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, measureNumericDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
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

		if x.R.MeasureNumerics[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Record != x.ID {
			t.Error("foreign key was wrong value", a.Record)
		}

		if exists, err := MeasureNumericExists(tx, a.Record, a.Metric); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testMeasureNumericsReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = measureNumeric.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testMeasureNumericsReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := MeasureNumericSlice{measureNumeric}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testMeasureNumericsSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := MeasureNumerics(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	measureNumericDBTypes = map[string]string{`Datum`: `numeric`, `Metric`: `integer`, `Record`: `integer`}
	_                     = bytes.MinRead
)

func testMeasureNumericsUpdate(t *testing.T) {
	t.Parallel()

	if len(measureNumericColumns) == len(measureNumericPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := measureNumericColumnsWithDefault

	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	if err = measureNumeric.Update(tx); err != nil {
		t.Error(err)
	}
}

func testMeasureNumericsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(measureNumericColumns) == len(measureNumericPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := measureNumericPrimaryKeyColumns

	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(measureNumericColumns, measureNumericPrimaryKeyColumns) {
		fields = measureNumericColumns
	} else {
		fields = strmangle.SetComplement(
			measureNumericColumns,
			measureNumericPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(measureNumeric))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := MeasureNumericSlice{measureNumeric}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testMeasureNumericsUpsert(t *testing.T) {
	t.Parallel()

	if len(measureNumericColumns) == len(measureNumericPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	measureNumeric := &MeasureNumeric{}
	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, true); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = measureNumeric.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert MeasureNumeric: %s", err)
	}

	count, err := MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := measureNumericPrimaryKeyColumns

	if err = randomize.Struct(seed, measureNumeric, measureNumericDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize MeasureNumeric struct: %s", err)
	}

	if err = measureNumeric.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert MeasureNumeric: %s", err)
	}

	count, err = MeasureNumerics(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
