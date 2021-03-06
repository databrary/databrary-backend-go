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

func testTokens(t *testing.T) {
	t.Parallel()

	query := Tokens(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testTokensLive(t *testing.T) {
	all, err := Tokens(dbMain.liveDbConn).All()
	if err != nil {
		t.Fatalf("failed to get all Tokens err: ", err)
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
	dumpCmd := exec.Command("psql", `-c "COPY (SELECT * FROM token) TO STDOUT" -d `, dbMain.DbName)
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
	dumpCmd = exec.Command("psql", `-c "COPY (SELECT * FROM token) TO STDOUT" -d `, dbMain.LiveTestDBName)
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
		t.Fatalf("TokensLive failed but it's probably trivial: %s", strings.Replace(result, "\t", " ", -1))
	}

}

func testTokensDelete(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = token.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokensQueryDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tokens(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokensSliceDeleteAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TokenSlice{token}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTokensExists(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TokenExists(tx, token.Token)
	if err != nil {
		t.Errorf("Unable to check if Token exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TokenExistsG to return true, but got false.")
	}
}

func testTokensFind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	tokenFound, err := FindToken(tx, token.Token)
	if err != nil {
		t.Error(err)
	}

	if tokenFound == nil {
		t.Error("want a record, got nil")
	}
}

func testTokensBind(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tokens(tx).Bind(token); err != nil {
		t.Error(err)
	}
}

func testTokensOne(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Tokens(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTokensAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tokenOne := &Token{}
	tokenTwo := &Token{}
	if err = randomize.Struct(seed, tokenOne, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Token struct: %s", err)
	}
	if err = randomize.Struct(seed, tokenTwo, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tokenOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tokenTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tokens(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTokensCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	tokenOne := &Token{}
	tokenTwo := &Token{}
	if err = randomize.Struct(seed, tokenOne, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {

		t.Errorf("Unable to randomize Token struct: %s", err)
	}
	if err = randomize.Struct(seed, tokenTwo, tokenDBTypes, false, tokenColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = tokenOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = tokenTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func tokenBeforeInsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterInsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterSelectHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenBeforeUpdateHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterUpdateHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenBeforeDeleteHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterDeleteHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenBeforeUpsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func tokenAfterUpsertHook(e boil.Executor, o *Token) error {
	*o = Token{}
	return nil
}

func testTokensHooks(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	empty := &Token{}

	AddTokenHook(boil.BeforeInsertHook, tokenBeforeInsertHook)
	if err = token.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", token)
	}
	tokenBeforeInsertHooks = []TokenHook{}

	AddTokenHook(boil.AfterInsertHook, tokenAfterInsertHook)
	if err = token.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", token)
	}
	tokenAfterInsertHooks = []TokenHook{}

	AddTokenHook(boil.AfterSelectHook, tokenAfterSelectHook)
	if err = token.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", token)
	}
	tokenAfterSelectHooks = []TokenHook{}

	AddTokenHook(boil.BeforeUpdateHook, tokenBeforeUpdateHook)
	if err = token.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", token)
	}
	tokenBeforeUpdateHooks = []TokenHook{}

	AddTokenHook(boil.AfterUpdateHook, tokenAfterUpdateHook)
	if err = token.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", token)
	}
	tokenAfterUpdateHooks = []TokenHook{}

	AddTokenHook(boil.BeforeDeleteHook, tokenBeforeDeleteHook)
	if err = token.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", token)
	}
	tokenBeforeDeleteHooks = []TokenHook{}

	AddTokenHook(boil.AfterDeleteHook, tokenAfterDeleteHook)
	if err = token.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", token)
	}
	tokenAfterDeleteHooks = []TokenHook{}

	AddTokenHook(boil.BeforeUpsertHook, tokenBeforeUpsertHook)
	if err = token.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", token)
	}
	tokenBeforeUpsertHooks = []TokenHook{}

	AddTokenHook(boil.AfterUpsertHook, tokenAfterUpsertHook)
	if err = token.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(token, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", token)
	}
	tokenAfterUpsertHooks = []TokenHook{}
}
func testTokensInsert(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTokensInsertWhitelist(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx, tokenColumns...); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTokensReload(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = token.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTokensReloadAll(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TokenSlice{token}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}

func testTokensSelect(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tokens(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	tokenDBTypes = map[string]string{`Expires`: `timestamp with time zone`, `Token`: `character varying`}
	_            = bytes.MinRead
)

func testTokensUpdate(t *testing.T) {
	t.Parallel()

	if len(tokenColumns) == len(tokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := tokenColumnsWithDefault

	if err = randomize.Struct(seed, token, tokenDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	if err = token.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTokensSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(tokenColumns) == len(tokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	blacklist := tokenPrimaryKeyColumns

	if err = randomize.Struct(seed, token, tokenDBTypes, true, blacklist...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(tokenColumns, tokenPrimaryKeyColumns) {
		fields = tokenColumns
	} else {
		fields = strmangle.SetComplement(
			tokenColumns,
			tokenPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(token))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TokenSlice{token}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}

func testTokensUpsert(t *testing.T) {
	t.Parallel()

	if len(tokenColumns) == len(tokenPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	var err error
	seed := randomize.NewSeed()
	token := &Token{}
	if err = randomize.Struct(seed, token, tokenDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = token.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Token: %s", err)
	}

	count, err := Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	blacklist := tokenPrimaryKeyColumns

	if err = randomize.Struct(seed, token, tokenDBTypes, false, blacklist...); err != nil {
		t.Errorf("Unable to randomize Token struct: %s", err)
	}

	if err = token.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Token: %s", err)
	}

	count, err = Tokens(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
