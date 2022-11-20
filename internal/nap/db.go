package nap

import (
	"context"
	"database/sql"
	"sync/atomic"
)

//go:generate mockgen -destination=db_mock.go -package=nap -source db.go

// Neighbor Access Protocol (NAP)

type SqlDb interface {
	Close() error
	Ping() error
	PingContext(ctx context.Context) error
	Begin() (*sql.Tx, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Slave() *sql.DB
	Master() *sql.DB
}

// DB is a logical database with multiple underlying physical databases
// forming a single master multiple slaves topology.
// Reads and writes are automatically directed to the correct physical db.
type sqlDb struct {
	pdbs  []*sql.DB // Physical databases
	count uint64    // Monotonically incrementing counter on each query
}

// Close closes all physical databases concurrently, releasing any open resources.
func (db *sqlDb) Close() error {
	return scatter(len(db.pdbs), func(i int) error {
		return db.pdbs[i].Close()
	})
}

// Ping verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (db *sqlDb) Ping() error {
	return scatter(len(db.pdbs), func(i int) error {
		return db.pdbs[i].Ping()
	})
}

// PingContext verifies a connection to the database is still alive,
// establishing a connection if necessary.
func (db *sqlDb) PingContext(ctx context.Context) error {
	return scatter(len(db.pdbs), func(i int) error {
		return db.pdbs[i].PingContext(ctx)
	})
}

// Begin starts a transaction. The default isolation level is dependent on
// the driver.
func (db *sqlDb) Begin() (*sql.Tx, error) {
	return db.Master().Begin()
}

// BeginTx starts a transaction.
//
// The provided context is used until the transaction is committed or rolled back.
// If the context is canceled, the sql package will roll back
// the transaction. Tx.Commit will return an error if the context provided to
// BeginTx is canceled.
//
// The provided TxOptions is optional and may be nil if defaults should be used.
// If a non-default isolation level is used that the driver doesn't support,
// an error will be returned.
func (db *sqlDb) BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return db.Master().BeginTx(ctx, opts)
}

// Exec executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// Exec uses the master as the physical db.
func (db *sqlDb) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Master().Exec(query, args...)
}

// ExecContext executes a query without returning any rows.
// The args are for any placeholder parameters in the query.
// ExecContext uses the master as the physical db.
func (db *sqlDb) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.Master().ExecContext(ctx, query, args...)
}

// Query executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
// Query uses a slave as the physical db.
func (db *sqlDb) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Slave().Query(query, args...)
}

// QueryContext executes a query that returns rows, typically a SELECT.
// The args are for any placeholder parameters in the query.
// QueryContext uses a slave as the physical db.
func (db *sqlDb) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	return db.Slave().QueryContext(ctx, query, args...)
}

// QueryRow executes a query that is expected to return at most one row.
// QueryRow always returns a non-nil value. Errors are deferred until
// Row's Scan method is called.
// If the query selects no rows, the *Row's Scan will return ErrNoRows.
// Otherwise, the *Row's Scan scans the first selected row and discards
// the route.
// QueryRow uses a slave as the physical db.
func (db *sqlDb) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.Slave().QueryRow(query, args...)
}

// Slave returns one of the physical databases which is a slave
func (db *sqlDb) Slave() *sql.DB {
	return db.pdbs[db.slave(len(db.pdbs))]
}

// Master returns the master physical database
func (db *sqlDb) Master() *sql.DB {
	return db.pdbs[0]
}

func (db *sqlDb) slave(n int) int {
	if n <= 1 {
		return 0
	}
	return int(1 + (atomic.AddUint64(&db.count, 1) % uint64(n-1)))
}

// Open concurrently opens each underlying physical db.
// dataSourceNames must be a list of connection strings with the first
// one being used as the master and the route as slaves.
func Open(driverName string, dataSourceNames []string) (SqlDb, error) {
	db := &sqlDb{pdbs: make([]*sql.DB, len(dataSourceNames))}

	err := scatter(len(db.pdbs), func(i int) (err error) {
		db.pdbs[i], err = sql.Open(driverName, dataSourceNames[i])
		return err
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
