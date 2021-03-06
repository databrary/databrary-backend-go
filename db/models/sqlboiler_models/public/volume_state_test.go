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

func testVolumeStates(t *testing.T) {
	t.Parallel()

	query := VolumeStates(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testVolumeStatesLive(t *testing.T) {
	all, err := VolumeStates(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all VolumeStates err: ", err)
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
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM volume_state) TO STDOUT" -d `, dbMain.DbName)
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
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM volume_state) TO STDOUT" -d `, dbMain.LiveTestDBName)
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
		t.Fatalf("VolumeStatesLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testVolumeStatesDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = volumeState.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testVolumeStatesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = VolumeStates(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testVolumeStatesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := VolumeStateSlice{volumeState}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testVolumeStatesExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := VolumeStateExists(tx, volumeState.Volume, volumeState.Key)
	if err != nil {
		t.Errorf("Unable to check if VolumeState exists: %s", err)
	}
	if !e {
		t.Errorf("Expected VolumeStateExistsG to return true, but got false.")
	}
}

func testVolumeStatesFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	volumeStateFound, err := FindVolumeState(tx, volumeState.Volume, volumeState.Key)
	if err != nil {
		t.Error(err)
	}

	if volumeStateFound == nil {
		t.Error("want a record, got nil")
	}
}

func testVolumeStatesBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = VolumeStates(tx).Bind(volumeState); err != nil {
		t.Error(err)
	}
}

func testVolumeStatesOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := VolumeStates(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testVolumeStatesAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeStateOne := &VolumeState{}
	volumeStateTwo := &VolumeState{}
	if err = randomize.Struct(seed, volumeStateOne, volumeStateDBTypes, false, volumeStateColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}
	if err = randomize.Struct(seed, volumeStateTwo, volumeStateDBTypes, false, volumeStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeStateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = volumeStateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := VolumeStates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testVolumeStatesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeStateOne := &VolumeState{}
	volumeStateTwo := &VolumeState{}
	if err = randomize.Struct(seed, volumeStateOne, volumeStateDBTypes, false, volumeStateColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}
	if err = randomize.Struct(seed, volumeStateTwo, volumeStateDBTypes, false, volumeStateColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeStateOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = volumeStateTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func volumeStateBeforeInsertHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateAfterInsertHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateAfterSelectHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateBeforeUpdateHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateAfterUpdateHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateBeforeDeleteHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateAfterDeleteHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateBeforeUpsertHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func volumeStateAfterUpsertHook(e boil.Executor, o *VolumeState) error {
	*o = VolumeState{}
	return nil
}

func testVolumeStatesHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	empty := &VolumeState{}

	AddVolumeStateHook(boil.BeforeInsertHook, volumeStateBeforeInsertHook)
	if err = volumeState.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateBeforeInsertHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.AfterInsertHook, volumeStateAfterInsertHook)
	if err = volumeState.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateAfterInsertHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.AfterSelectHook, volumeStateAfterSelectHook)
	if err = volumeState.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateAfterSelectHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.BeforeUpdateHook, volumeStateBeforeUpdateHook)
	if err = volumeState.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateBeforeUpdateHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.AfterUpdateHook, volumeStateAfterUpdateHook)
	if err = volumeState.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateAfterUpdateHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.BeforeDeleteHook, volumeStateBeforeDeleteHook)
	if err = volumeState.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateBeforeDeleteHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.AfterDeleteHook, volumeStateAfterDeleteHook)
	if err = volumeState.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateAfterDeleteHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.BeforeUpsertHook, volumeStateBeforeUpsertHook)
	if err = volumeState.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateBeforeUpsertHooks = []VolumeStateHook{}

	AddVolumeStateHook(boil.AfterUpsertHook, volumeStateAfterUpsertHook)
	if err = volumeState.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(volumeState, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", volumeState)
	}
	volumeStateAfterUpsertHooks = []VolumeStateHook{}
}
func testVolumeStatesInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testVolumeStatesInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx, volumeStateColumns...); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testVolumeStateToOneVolumeUsingVolume(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Volume
	var local VolumeState

	foreignBlacklist := volumeColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, volumeDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	localBlacklist := volumeStateColumnsWithDefault
	if err := randomize.Struct(seed, &local, volumeStateDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Volume = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.VolumeByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := VolumeStateSlice{&local}
	if err = local.L.LoadVolume(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Volume = nil
	if err = local.L.LoadVolume(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Volume == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testVolumeStateToOneSetOpVolumeUsingVolume(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a VolumeState
	var b, c Volume

	foreignBlacklist := strmangle.SetComplement(volumePrimaryKeyColumns, volumeColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, volumeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, volumeDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Volume struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(volumeStatePrimaryKeyColumns, volumeStateColumnsWithoutDefault)
	if err := randomize.Struct(seed, &a, volumeStateDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Volume{&b, &c} {
		err = a.SetVolume(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Volume != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.VolumeStates[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Volume != x.ID {
			t.Error("foreign key was wrong value", a.Volume)
		}

		if exists, err := VolumeStateExists(tx, a.Volume, a.Key); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testVolumeStatesReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = volumeState.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testVolumeStatesReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := VolumeStateSlice{volumeState}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testVolumeStatesSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := VolumeStates(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	volumeStateDBTypes = map[string]string{`Key`: `character varying`, `Public`: `boolean`, `Value`: `json`, `Volume`: `integer`}
	_                  = bytes.MinRead
)

func testVolumeStatesUpdate(t *testing.T) {
	t.Parallel()

	if len(volumeStateColumns) == len(volumeStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := volumeStateColumnsWithDefault

	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	if err = volumeState.Update(tx); err != nil {
		t.Error(err)
	}
}

func testVolumeStatesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(volumeStateColumns) == len(volumeStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := volumeStatePrimaryKeyColumns

	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(volumeStateColumns, volumeStatePrimaryKeyColumns) {
		fields = volumeStateColumns
	} else {
		fields = strmangle.SetComplement(
			volumeStateColumns,
			volumeStatePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(volumeState))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := VolumeStateSlice{volumeState}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testVolumeStatesUpsert(t *testing.T) {
	t.Parallel()

	if len(volumeStateColumns) == len(volumeStatePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	volumeState := &VolumeState{}
	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, true); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = volumeState.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert VolumeState: %s", err)
	}

	count, err := VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := volumeStatePrimaryKeyColumns

	if err = randomize.Struct(seed, volumeState, volumeStateDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize VolumeState struct: %s", err)
	}

	if err = volumeState.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert VolumeState: %s", err)
	}

	count, err = VolumeStates(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
