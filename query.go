package sql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type Query struct {
	db *sqlx.DB
}

func NewQuery(db *sqlx.DB) *Query {
	return &Query{
		db: db,
	}
}

func (q *Query) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return q.db.SelectContext(ctx, dest, query, args...)
}

func (q *Query) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return q.db.GetContext(ctx, dest, query, args...)
}

func (q *Query) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return q.db.ExecContext(ctx, query, args...)
}

func (q *Query) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return q.db.QueryxContext(ctx, query, args...)
}

func (q *Query) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return q.db.QueryRowxContext(ctx, query, args...)
}

func (q *Query) Prepare(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return q.db.PreparexContext(ctx, query)
}

func (q *Query) Close() error {
	return q.db.Close()
}
