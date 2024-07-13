package sql

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

const (
	postgres = "postgres"
)

type PgConnectionManager struct {
	db *sqlx.DB
}

func NewPgConnectionManager(connectionString string) (*PgConnectionManager, error) {

	// Open a connection to the database
	db, err := sqlx.Open(postgres, connectionString)
	if err != nil {
		log.Printf("failed to open connection: %v", err)
		return nil, err
	}

	// Ping the database to verify the connection
	if err := db.Ping(); err != nil {
		log.Printf("failed to ping: %v", err)
		db.Close()
		return nil, err
	}

	return &PgConnectionManager{
		db: db,
	}, nil
}

func (cm *PgConnectionManager) Close() error {
	return cm.db.Close()
}

func (cm *PgConnectionManager) Query() *Query {
	return NewQuery(cm.db)
}
func (cm *PgConnectionManager) Transaction() *Transaction {
	return NewTransaction(cm.db)
}

func (cm *PgConnectionManager) CheckDatabaseHealth() error {
	if cm.db == nil {
		return fmt.Errorf("database connection is not initialized")
	}

	ctx := context.Background()
	if err := cm.db.PingContext(ctx); err != nil {
		return err
	}

	return nil
}
