// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package public

import (
	"bytes"
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

func testFunders(t *testing.T) {
	t.Parallel()

	query := Funders(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFundersLive(t *testing.T) {
	all, err := Funders(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all Funders err: ", err)
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
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM funder) TO STDOUT" -d `, dbMain.DbName)
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
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM funder) TO STDOUT" -d `, dbMain.LiveTestDBName)
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
		t.Fatalf("FundersLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testFundersDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = funder.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFundersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Funders(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFundersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FunderSlice{funder}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFundersExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := FunderExists(tx, funder.FundrefID)
	if err != nil {
		t.Errorf("Unable to check if Funder exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FunderExistsG to return true, but got false.")
	}
}

func testFundersFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	funderFound, err := FindFunder(tx, funder.FundrefID)
	if err != nil {
		t.Error(err)
	}

	if funderFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFundersBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Funders(tx).Bind(funder); err != nil {
		t.Error(err)
	}
}

func testFundersOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Funders(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFundersAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funderOne := &Funder{}
	funderTwo := &Funder{}
	if err = randomize.Struct(seed, funderOne, funderDBTypes, false, funderColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Funder struct: %s", err)
	}
	if err = randomize.Struct(seed, funderTwo, funderDBTypes, false, funderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = funderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Funders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFundersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funderOne := &Funder{}
	funderTwo := &Funder{}
	if err = randomize.Struct(seed, funderOne, funderDBTypes, false, funderColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Funder struct: %s", err)
	}
	if err = randomize.Struct(seed, funderTwo, funderDBTypes, false, funderColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funderOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = funderTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func funderBeforeInsertHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderAfterInsertHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderAfterSelectHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderBeforeUpdateHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderAfterUpdateHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderBeforeDeleteHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderAfterDeleteHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderBeforeUpsertHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func funderAfterUpsertHook(e boil.Executor, o *Funder) error {
	*o = Funder{}
	return nil
}

func testFundersHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	empty := &Funder{}

	AddFunderHook(boil.BeforeInsertHook, funderBeforeInsertHook)
	if err = funder.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", funder)
	}
	funderBeforeInsertHooks = []FunderHook{}

	AddFunderHook(boil.AfterInsertHook, funderAfterInsertHook)
	if err = funder.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", funder)
	}
	funderAfterInsertHooks = []FunderHook{}

	AddFunderHook(boil.AfterSelectHook, funderAfterSelectHook)
	if err = funder.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", funder)
	}
	funderAfterSelectHooks = []FunderHook{}

	AddFunderHook(boil.BeforeUpdateHook, funderBeforeUpdateHook)
	if err = funder.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", funder)
	}
	funderBeforeUpdateHooks = []FunderHook{}

	AddFunderHook(boil.AfterUpdateHook, funderAfterUpdateHook)
	if err = funder.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", funder)
	}
	funderAfterUpdateHooks = []FunderHook{}

	AddFunderHook(boil.BeforeDeleteHook, funderBeforeDeleteHook)
	if err = funder.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", funder)
	}
	funderBeforeDeleteHooks = []FunderHook{}

	AddFunderHook(boil.AfterDeleteHook, funderAfterDeleteHook)
	if err = funder.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", funder)
	}
	funderAfterDeleteHooks = []FunderHook{}

	AddFunderHook(boil.BeforeUpsertHook, funderBeforeUpsertHook)
	if err = funder.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", funder)
	}
	funderBeforeUpsertHooks = []FunderHook{}

	AddFunderHook(boil.AfterUpsertHook, funderAfterUpsertHook)
	if err = funder.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(funder, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", funder)
	}
	funderAfterUpsertHooks = []FunderHook{}
}
func testFundersInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFundersInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx, funderColumns...); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFunderToManyVolumeFundings(t *testing.T) {
	var err error
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a Funder
	var b, c VolumeFunding

	foreignBlacklist := volumeFundingColumnsWithDefault
	if err := randomize.Struct(seed, &b, volumeFundingDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeFunding struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, volumeFundingDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeFunding struct: %s", err)
	}
	localBlacklist := funderColumnsWithDefault
	if err := randomize.Struct(seed, &a, funderDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	b.Funder = a.FundrefID
	c.Funder = a.FundrefID
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	volumeFunding, err := a.VolumeFundingsByFk(tx).All()
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range volumeFunding {
		if v.Funder == b.Funder {
			bFound = true
		}
		if v.Funder == c.Funder {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FunderSlice{&a}
	if err = a.L.LoadVolumeFundings(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.VolumeFundings); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.VolumeFundings = nil
	if err = a.L.LoadVolumeFundings(tx, true, &a); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.VolumeFundings); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", volumeFunding)
	}
}

func testFunderToManyAddOpVolumeFundings(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	var a Funder
	var b, c, d, e VolumeFunding

	seed := randomize.NewSeed()
	localComplelementList := strmangle.SetComplement(funderPrimaryKeyColumns, funderColumnsWithoutDefault)
	if err = randomize.Struct(seed, &a, funderDBTypes, false, localComplelementList...); err != nil {
		t.Fatal(err)
	}

	foreignComplementList := strmangle.SetComplement(volumeFundingPrimaryKeyColumns, volumeFundingColumnsWithoutDefault)

	foreigners := []*VolumeFunding{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, volumeFundingDBTypes, false, foreignComplementList...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(tx); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*VolumeFunding{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddVolumeFundings(tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.FundrefID != first.Funder {
			t.Error("foreign key was wrong value", a.FundrefID, first.Funder)
		}
		if a.FundrefID != second.Funder {
			t.Error("foreign key was wrong value", a.FundrefID, second.Funder)
		}

		if first.R.Funder != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Funder != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.VolumeFundings[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.VolumeFundings[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.VolumeFundingsByFk(tx).Count()
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testFundersReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = funder.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testFundersReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := FunderSlice{funder}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testFundersSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Funders(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	funderDBTypes = map[string]string{`FundrefID`: `bigint`, `Name`: `text`}
	_             = bytes.MinRead
)

func testFundersUpdate(t *testing.T) {
	t.Parallel()

	if len(funderColumns) == len(funderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := funderColumnsWithDefault

	if err = randomize.Struct(seed, funder, funderDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	if err = funder.Update(tx); err != nil {
		t.Error(err)
	}
}

func testFundersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(funderColumns) == len(funderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := funderPrimaryKeyColumns

	if err = randomize.Struct(seed, funder, funderDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(funderColumns, funderPrimaryKeyColumns) {
		fields = funderColumns
	} else {
		fields = strmangle.SetComplement(
			funderColumns,
			funderPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(funder))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := FunderSlice{funder}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testFundersUpsert(t *testing.T) {
	t.Parallel()

	if len(funderColumns) == len(funderPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	funder := &Funder{}
	if err = randomize.Struct(seed, funder, funderDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = funder.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Funder: %s", err)
	}

	count, err := Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := funderPrimaryKeyColumns

	if err = randomize.Struct(seed, funder, funderDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize Funder struct: %s", err)
	}

	if err = funder.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Funder: %s", err)
	}

	count, err = Funders(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
