package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store interface {
	Querier // allow access to all the methods which use *Queries
	CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error)
}

type SQLStore struct {
	// All the individual query functions provided by Queries will be available to Store
	*Queries

	// To manage DB transaction
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("txErr: %v - rbErr: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}
