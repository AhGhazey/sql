package sql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Connection interface {
	Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row
	Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Prepare(ctx context.Context, query string) (*sqlx.Stmt, error)
	Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}
