package runner

import (
	"database/sql"

	"github.com/mgutz/dat"
)

// Tx is a transaction for the given Session
type Tx struct {
	*sql.Tx
	*Queryable
}

// Begin creates a transaction for the given session
func (sess *Session) Begin() (*Tx, error) {
	tx, err := sess.DB.Begin()
	if err != nil {
		return nil, dat.Events.EventErr("begin.error", err)
	}
	dat.Events.Event("begin")

	return &Tx{tx, &Queryable{tx}}, nil
}

// Commit finishes the transaction
func (tx *Tx) Commit() error {
	err := tx.Tx.Commit()
	if err != nil {
		return dat.Events.EventErr("commit.error", err)
	}
	dat.Events.Event("commit")
	return nil
}

// Rollback cancels the transaction
func (tx *Tx) Rollback() error {
	err := tx.Tx.Rollback()
	if err != nil {
		return dat.Events.EventErr("rollback", err)
	}
	dat.Events.Event("rollback")
	return nil
}

// RollbackUnlessCommitted rollsback the transaction unless it has already been committed or rolled back.
// Useful to defer tx.RollbackUnlessCommitted() -- so you don't have to handle N failure cases
// Keep in mind the only way to detect an error on the rollback is via the event log.
func (tx *Tx) RollbackUnlessCommitted() {
	err := tx.Tx.Rollback()
	if err == sql.ErrTxDone {
		// ok
	} else if err != nil {
		dat.Events.EventErr("rollback_unless_committed", err)
	} else {
		dat.Events.Event("rollback")
	}
}
