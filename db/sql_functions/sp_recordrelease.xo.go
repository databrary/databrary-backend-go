// Package xo_models contains the types for schema 'public'.
package xo_models

import "github.com/databrary/sqlboiler/boil"

// GENERATED BY XO. DO NOT EDIT.

// RecordRelease calls the stored procedure 'public.record_release(integer) release' on db.
func RecordRelease(exec boil.Executor, v0 int) (Release, error) {
	var err error

	// sql query
	const query = `SELECT public.record_release($1)`

	// run query
	var ret Release

	err = exec.QueryRow(query, v0).Scan(&ret)
	if err != nil {
		return Release{}, err
	}

	return ret, nil
}
