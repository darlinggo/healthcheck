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

// Check verifies that a connection to the database
// can be obtained, using the Ping method of the
// underlying *sql.DB. It then makes sure the database
// is still reachable, by sending a simple query.
func (s SQL) Check(ctx context.Context) error {
	err := s.DB.Ping()
	if err != nil {
		return err
	}
	var one int
	err = s.DB.QueryRow("SELECT 1;").Scan(&one)
	return err
}

// LogInfo returns the ID string associated with the
// SQL it is called on. It should usually be set to a
// connection string, a database name, or anything else
// that would be useful to have in the log output.
func (s SQL) LogInfo(ctx context.Context) string {
	return s.ID
}
