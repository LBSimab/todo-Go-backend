package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {

	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v , rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TaskTxParams struct {
	Name sql.NullString `json:"name"`

	Supervisor sql.NullInt32 `json:"supervisor"`

	Category sql.NullString `json:"category"`
}

type TaskTxResults struct {
	Task Task `json:"task"`
}

func (store *Store) TaskTx(ctx context.Context, arg TaskTxParams) (TaskTxResults, error) {

	var result TaskTxResults

	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		result.Task, err = q.Createtask(ctx, CreatetaskParams{
			Name:       arg.Name,
			Supervisor: arg.Supervisor,
			Category:   arg.Category,
		})
		if err != nil {
			return err
		}
		return nil
	})
	return result, err

}
