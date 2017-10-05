// This file is generated by SQLBoiler (https://github.com/databrary/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// EDIT AT YOUR OWN RISK

package public

import (
	"bytes"
	"github.com/databrary/databrary/db/models/custom_types"
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

func testTagUses(t *testing.T) {
	t.Parallel()

	query := TagUses(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTagUsesLive(t *testing.T) {
	all, err := TagUses(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all TagUses err: ", err)
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
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM tag_use) TO STDOUT" -d `, dbMain.DbName)
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
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM tag_use) TO STDOUT" -d `, dbMain.LiveTestDBName)
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
		t.Fatalf("TagUsesLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testTagUsesDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = tagUse.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTagUsesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TagUses(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTagUsesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TagUseSlice{tagUse}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTagUsesExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TagUseExists(tx, tagUse.Container, tagUse.Segment, tagUse.Tag, tagUse.Who)
	if err != nil {
		t.Errorf("Unable to check if TagUse exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TagUseExistsG to return true, but got false.")
	}
}

func testTagUsesFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	tagUseFound, err := FindTagUse(tx, tagUse.Container, tagUse.Segment, tagUse.Tag, tagUse.Who)
	if err != nil {
		t.Error(err)
	}

	if tagUseFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTagUsesBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = TagUses(tx).Bind(tagUse); err != nil {
		t.Error(err)
	}
}

func testTagUsesOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := TagUses(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTagUsesAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUseOne := &TagUse{}
	tagUseTwo := &TagUse{}
	if err = randomize.Struct(seed, tagUseOne, tagUseDBTypes, false, tagUseColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}
	if err = randomize.Struct(seed, tagUseTwo, tagUseDBTypes, false, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUseOne.Segment = custom_types.SegmentRandom()
	tagUseTwo.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUseOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tagUseTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TagUses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTagUsesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUseOne := &TagUse{}
	tagUseTwo := &TagUse{}
	if err = randomize.Struct(seed, tagUseOne, tagUseDBTypes, false, tagUseColumnsWithCustom...); err != nil {

		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}
	if err = randomize.Struct(seed, tagUseTwo, tagUseDBTypes, false, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUseOne.Segment = custom_types.SegmentRandom()
	tagUseTwo.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUseOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tagUseTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func tagUseBeforeInsertHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseAfterInsertHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseAfterSelectHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseBeforeUpdateHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseAfterUpdateHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseBeforeDeleteHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseAfterDeleteHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseBeforeUpsertHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func tagUseAfterUpsertHook(e boil.Executor, o *TagUse) error {
	*o = TagUse{}
	return nil
}

func testTagUsesHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	empty := &TagUse{}

	AddTagUseHook(boil.BeforeInsertHook, tagUseBeforeInsertHook)
	if err = tagUse.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", tagUse)
	}
	tagUseBeforeInsertHooks = []TagUseHook{}

	AddTagUseHook(boil.AfterInsertHook, tagUseAfterInsertHook)
	if err = tagUse.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", tagUse)
	}
	tagUseAfterInsertHooks = []TagUseHook{}

	AddTagUseHook(boil.AfterSelectHook, tagUseAfterSelectHook)
	if err = tagUse.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", tagUse)
	}
	tagUseAfterSelectHooks = []TagUseHook{}

	AddTagUseHook(boil.BeforeUpdateHook, tagUseBeforeUpdateHook)
	if err = tagUse.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", tagUse)
	}
	tagUseBeforeUpdateHooks = []TagUseHook{}

	AddTagUseHook(boil.AfterUpdateHook, tagUseAfterUpdateHook)
	if err = tagUse.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", tagUse)
	}
	tagUseAfterUpdateHooks = []TagUseHook{}

	AddTagUseHook(boil.BeforeDeleteHook, tagUseBeforeDeleteHook)
	if err = tagUse.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", tagUse)
	}
	tagUseBeforeDeleteHooks = []TagUseHook{}

	AddTagUseHook(boil.AfterDeleteHook, tagUseAfterDeleteHook)
	if err = tagUse.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", tagUse)
	}
	tagUseAfterDeleteHooks = []TagUseHook{}

	AddTagUseHook(boil.BeforeUpsertHook, tagUseBeforeUpsertHook)
	if err = tagUse.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", tagUse)
	}
	tagUseBeforeUpsertHooks = []TagUseHook{}

	AddTagUseHook(boil.AfterUpsertHook, tagUseAfterUpsertHook)
	if err = tagUse.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(tagUse, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", tagUse)
	}
	tagUseAfterUpsertHooks = []TagUseHook{}
}
func testTagUsesInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTagUsesInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx, tagUseColumns...); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTagUseToOneContainerUsingContainer(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Container
	var local TagUse

	foreignBlacklist := containerColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, containerDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	localBlacklist := tagUseColumnsWithDefault
	localBlacklist = append(localBlacklist, tagUseColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, tagUseDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
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

	slice := TagUseSlice{&local}
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

func testTagUseToOneTagUsingTag(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Tag
	var local TagUse

	foreignBlacklist := tagColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, tagDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}
	localBlacklist := tagUseColumnsWithDefault
	localBlacklist = append(localBlacklist, tagUseColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, tagUseDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}
	local.Segment = custom_types.SegmentRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Tag = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.TagByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TagUseSlice{&local}
	if err = local.L.LoadTag(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Tag == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Tag = nil
	if err = local.L.LoadTag(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Tag == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTagUseToOneAccountUsingWho(t *testing.T) {
	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var foreign Account
	var local TagUse

	foreignBlacklist := accountColumnsWithDefault
	if err := randomize.Struct(seed, &foreign, accountDBTypes, true, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	localBlacklist := tagUseColumnsWithDefault
	localBlacklist = append(localBlacklist, tagUseColumnsWithCustom...)

	if err := randomize.Struct(seed, &local, tagUseDBTypes, true, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}
	local.Segment = custom_types.SegmentRandom()

	if err := foreign.Insert(tx); err != nil {
		t.Fatal(err)
	}

	local.Who = foreign.ID
	if err := local.Insert(tx); err != nil {
		t.Fatal(err)
	}

	check, err := local.WhoByFk(tx).One()
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := TagUseSlice{&local}
	if err = local.L.LoadWho(tx, false, &slice); err != nil {
		t.Fatal(err)
	}
	if local.R.Who == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Who = nil
	if err = local.L.LoadWho(tx, true, &local); err != nil {
		t.Fatal(err)
	}
	if local.R.Who == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testTagUseToOneSetOpContainerUsingContainer(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a TagUse
	var b, c Container

	foreignBlacklist := strmangle.SetComplement(containerPrimaryKeyColumns, containerColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, containerDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, containerDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(tagUsePrimaryKeyColumns, tagUseColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, tagUseColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, tagUseDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
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

		if x.R.TagUses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Container != x.ID {
			t.Error("foreign key was wrong value", a.Container)
		}

		if exists, err := TagUseExists(tx, a.Container, a.Segment, a.Tag, a.Who); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testTagUseToOneSetOpTagUsingTag(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a TagUse
	var b, c Tag

	foreignBlacklist := strmangle.SetComplement(tagPrimaryKeyColumns, tagColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, tagDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, tagDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Tag struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(tagUsePrimaryKeyColumns, tagUseColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, tagUseColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, tagUseDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}
	a.Segment = custom_types.SegmentRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Tag{&b, &c} {
		err = a.SetTag(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Tag != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.TagUses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Tag != x.ID {
			t.Error("foreign key was wrong value", a.Tag)
		}

		if exists, err := TagUseExists(tx, a.Container, a.Segment, a.Tag, a.Who); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}
func testTagUseToOneSetOpAccountUsingWho(t *testing.T) {
	var err error

	tx := MustTx(boil.Begin())
	defer tx.Rollback()

	seed := randomize.NewSeed()

	var a TagUse
	var b, c Account

	foreignBlacklist := strmangle.SetComplement(accountPrimaryKeyColumns, accountColumnsWithoutDefault)
	if err := randomize.Struct(seed, &b, accountDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	if err := randomize.Struct(seed, &c, accountDBTypes, false, foreignBlacklist...); err != nil {
		t.Errorf("Unable to randomize Account struct: %s", err)
	}
	localBlacklist := strmangle.SetComplement(tagUsePrimaryKeyColumns, tagUseColumnsWithoutDefault)
	localBlacklist = append(localBlacklist, tagUseColumnsWithCustom...)

	if err := randomize.Struct(seed, &a, tagUseDBTypes, false, localBlacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}
	a.Segment = custom_types.SegmentRandom()

	if err := a.Insert(tx); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(tx); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Account{&b, &c} {
		err = a.SetWho(tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Who != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.WhoTagUses[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.Who != x.ID {
			t.Error("foreign key was wrong value", a.Who)
		}

		if exists, err := TagUseExists(tx, a.Container, a.Segment, a.Tag, a.Who); err != nil {
			t.Fatal(err)
		} else if !exists {
			t.Error("want 'a' to exist")
		}

	}
}

func testTagUsesReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = tagUse.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTagUsesReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TagUseSlice{tagUse}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testTagUsesSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := TagUses(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tagUseDBTypes = map[string]string{`Container`: `integer`, `Segment`: `USER-DEFINED`, `Tag`: `integer`, `Who`: `integer`}
	_             = bytes.MinRead
)

func testTagUsesUpdate(t *testing.T) {
	t.Parallel()

	if len(tagUseColumns) == len(tagUsePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := tagUseColumnsWithDefault
	blacklist = append(blacklist, tagUseColumnsWithCustom...)

	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	if err = tagUse.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTagUsesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tagUseColumns) == len(tagUsePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := tagUsePrimaryKeyColumns
	blacklist = append(blacklist, tagUseColumnsWithCustom...)

	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tagUseColumns, tagUsePrimaryKeyColumns) {
		fields = tagUseColumns
	} else {
		fields = strmangle.SetComplement(
			tagUseColumns,
			tagUsePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(tagUse))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TagUseSlice{tagUse}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testTagUsesUpsert(t *testing.T) {
	t.Parallel()

	if len(tagUseColumns) == len(tagUsePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	tagUse := &TagUse{}
	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, true, tagUseColumnsWithCustom...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tagUse.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert TagUse: %s", err)
	}

	count, err := TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := tagUsePrimaryKeyColumns

	blacklist = append(blacklist, tagUseColumnsWithCustom...)

	if err = randomize.Struct(seed, tagUse, tagUseDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize TagUse struct: %s", err)
	}

	tagUse.Segment = custom_types.SegmentRandom()

	if err = tagUse.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert TagUse: %s", err)
	}

	count, err = TagUses(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
