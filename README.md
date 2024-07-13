# Go SQL Library

This SQL library provides a simple and efficient way to interact with PostgreSQL databases in Go applications. It wraps the `sqlx` package, offering an easy-to-use interface for database operations including queries, transactions, and connection management.

## Features

- Connection management for PostgreSQL
- Support for regular queries and transactions
- Context-aware database operations
- Easy-to-use methods for common database operations (Select, Get, Exec, Query, QueryRow, Prepare)
- Database health check functionality

## Installation

To use this SQL library in your Go project, run:

```bash
go get github.com/ahghazey/sql
```

## Usage
### Initializing the Connection
To create a new PostgreSQL connection manager:

```go
import "github.com/ahghazey/sql"

connectionString := "postgres://username:password@localhost/dbname?sslmode=disable"
cm, err := sql.NewPgConnectionManager(connectionString)
if err != nil {
    // Handle error
}
defer cm.Close()
```
### Performing Queries
To perform a query:
```go
query := cm.Query()

var users []User
err := query.Select(context.Background(), &users, "SELECT * FROM users WHERE active = $1", true)
if err != nil {
    // Handle error
}
```
### Using Transactions
To perform operations within a transaction:
```go
tx := cm.Transaction()
err := tx.Begin(context.Background())
if err != nil {
// Handle error
}

// Perform operations
_, err = tx.Exec(context.Background(), "INSERT INTO users (name, email) VALUES ($1, $2)", "folan el folany", "folan@example.com")
if err != nil {
tx.Rollback(context.Background())
// Handle error
}

err = tx.Commit(context.Background())
if err != nil {
// Handle error
}
```

### Checking Database Health
To check the health of the database connection:
```go
err := cm.CheckDatabaseHealth()
if err != nil {
    // Handle error
}
```

## API Reference

### Connection Manager
* NewPgConnectionManager(connectionString string) (*PgConnectionManager, error)
* Close() error
* Query() *Query
* Transaction() *Transaction
* CheckDatabaseHealth() error

### Query
* Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
* Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
* Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
* Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
* QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row
* Prepare(ctx context.Context, query string) (*sqlx.Stmt, error)
* Close() error

### Transaction 
* Begin(ctx context.Context) error
* Commit(ctx context.Context) error
* Rollback(ctx context.Context) error
* Query(ctx context.Context, query string, args ...interface{}) (*sqlx.Rows, error)
* QueryRow(ctx context.Context, query string, args ...interface{}) *sqlx.Row
* Exec(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
* Prepare(ctx context.Context, query string) (*sqlx.Stmt, error)
* Select(ctx context.Context, dest interface{}, query string, args ...interface{}) error
* Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
* Close() error