package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier
}

type SqlStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SqlStore{
		Queries: New(db),
		db:      db,
	}
}

func (store *SqlStore) ExecTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {

		return err
	}

	q := New(tx)

	errTx := fn(q)

	if errTx != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("rb Error: %v, tx: %v", rbErr, errTx)
		}
		return errTx
	}

	return tx.Commit()
}

// func (store *S)
