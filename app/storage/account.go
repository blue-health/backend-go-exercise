package storage

import "github.com/jmoiron/sqlx"

// We follow the CQRS pattern when
// interacting with the database
type (
	//nolint:structcheck  //to-do
	AccountReader struct{ db *sqlx.DB }
	//nolint:structcheck  //to-do
	AccountWriter struct{ db *sqlx.DB }
)
