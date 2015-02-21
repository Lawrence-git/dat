package runner

import "database/sql"

// Connection is a connection to the database with an EventReceiver
type Connection struct {
	DB *sql.DB
	*Queryable
}

// NewConnection instantiates a Connection for a given database/sql connection
func NewConnection(db *sql.DB) *Connection {
	return &Connection{db, &Queryable{db}}
}
