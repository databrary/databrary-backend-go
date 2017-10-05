// Package xo_models contains the types for schema 'public'.
package xo_models

// GENERATED BY XO. DO NOT EDIT.

import (
	"time"

	"github.com/databrary/databrary-backend-go/db/models/custom_types"
	"github.com/databrary/sqlboiler/boil"
)

// Duration calls the stored procedure 'public.duration(segment) interval' on db.
func Duration(exec boil.Executor, v0 custom_types.Segment) (*time.Duration, error) {
	var err error

	// sql query
	const query = `SELECT public.duration($1)`

	// run query
	var ret *time.Duration

	err = exec.QueryRow(query, v0).Scan(&ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}
