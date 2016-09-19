package healthcheck

import (
	"database/sql"

	"context"
)

// SQL is a type that fills the Checker interface
// for an *sql.DB. ID is used as the identifying
// information returned by LogInfo, which is useful
// if more than one *sql.DB is needed for a service's
// health check.
type SQL struct {
	DB *sql.DB
	ID string
}

// NewSQL returns an SQL instance using the passed
// *sql.DB and identifier.
func NewSQL(db *sql.DB, id string) SQL {
	return SQL{
		DB: db,
		ID: id,
	}
}

// Check returns the output of the Ping method for
// the *sql.DB.
func (s SQL) Check(ctx context.Context) error {
	return s.DB.Ping()
}

// LogInfo returns the ID string associated with the
// SQL it is called on. It should usually be set to a
// connection string, a database name, or anything else
// that would be useful to have in the log output.
func (s SQL) LogInfo(ctx context.Context) string {
	return s.ID
}
