// Package xo_models contains the types for schema 'public'.
package xo_models

// GENERATED BY XO. DO NOT EDIT.

import (
	"time"

	"github.com/databrary/databrary-backend-go/db/models/custom_types"
	"github.com/databrary/sqlboiler/boil"
)

// Segment calls the stored procedure 'public.segment(interval, interval, interval, interval, interval, text) segment' on db.
func Segment(exec boil.Executor, v0 *time.Duration, v1 *time.Duration, v2 *time.Duration, v3 *time.Duration, v4 *time.Duration, v5 string) (custom_types.Segment, error) {
	var err error

	// sql query
	const query = `SELECT public.segment($1, $2, $3, $4, $5, $6)`

	// run query
	var ret custom_types.Segment

	err = exec.QueryRow(query, v0, v1, v2, v3, v4, v5).Scan(&ret)
	if err != nil {
		return custom_types.Segment{}, err
	}

	return ret, nil
}
