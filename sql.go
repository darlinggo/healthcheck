package healthcheck

import (
	"database/sql"

	"golang.org/x/net/context"
)

type SQL struct {
	DB *sql.DB
	ID string
}

func NewSQL(db *sql.DB, id string) SQL {
	return SQL{
		DB: db,
		ID: id,
	}
}

func (s SQL) Check(ctx context.Context) error {
	return s.DB.Ping()
}

func (s SQL) LogInfo(ctx context.Context) string {
	return s.ID
}
