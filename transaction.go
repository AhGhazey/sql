package sql

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"log"
)

type TX interface {
	Begin(ctx context.Context) error
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}

type Transaction struct {
	db *sqlx.DB
	tx *sqlx.Tx
}

func NewTransaction(db *sqlx.DB) *Transaction {
	return &Transaction{
		db: db,
	}
}

func (t *Transaction) Begin(ctx context.Context) error {
	tx, err := t.db.BeginTxx(ctx, nil)
	if err != nil {
		log.Print("failed to start transaction")
		return err
	}
	t.tx = tx
	return nil
}

func (t *Transaction) Commit(ctx context.Context) error {
	err := t.tx.Commit()
	if err != nil {
		log.Print("failed to commit transaction")
	}
	return err
}

func (t *Transaction) Rollback(ctx context.Context) error {
	err := t.tx.Rollback()
	if err != nil {
		log.Print("failed to rollback transaction")
	}
	return err
}

func (t *Transaction) Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error) {
	return t.tx.QueryxContext(ctx, query, args...)
}

func (t *Transaction) QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row {
	return t.tx.QueryRowxContext(ctx, query, args...)
}

func (t *Transaction) Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return t.tx.ExecContext(ctx, query, args...)
}

func (t *Transaction) Prepare(ctx context.Context, query string) (*sqlx.Stmt, error) {
	return t.tx.PreparexContext(ctx, query)
}

func (t *Transaction) Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.tx.SelectContext(ctx, dest, query, args...)
}

func (t *Transaction) Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	return t.tx.GetContext(ctx, dest, query, args...)
}

func (t *Transaction) Close() error {
	return t.db.Close()
}
