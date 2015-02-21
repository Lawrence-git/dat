package runner

import "database/sql"

// Session represents a business unit of execution for some connection
type Session struct {
	DB *sql.DB
	*Queryable
}

// NewSession instantiates a Session for the Connection
func (cxn *Connection) NewSession() *Session {
	return &Session{cxn.DB, &Queryable{cxn.DB}}
}
