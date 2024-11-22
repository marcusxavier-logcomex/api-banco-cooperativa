// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const decreaseUserBalance = `-- name: DecreaseUserBalance :exec
UPDATE users
SET balance = balance - $1
WHERE id = $2
`

type DecreaseUserBalanceParams struct {
	Balance int32
	ID      int32
}

func (q *Queries) DecreaseUserBalance(ctx context.Context, arg DecreaseUserBalanceParams) error {
	_, err := q.db.Exec(ctx, decreaseUserBalance, arg.Balance, arg.ID)
	return err
}

const getLastTenTransactions = `-- name: GetLastTenTransactions :many
SELECT id, user_id, value, type, description, created_at FROM transactions
WHERE user_id = $1
ORDER BY id DESC
LIMIT 10
`

func (q *Queries) GetLastTenTransactions(ctx context.Context, userID pgtype.Int4) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, getLastTenTransactions, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Value,
			&i.Type,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getUser = `-- name: GetUser :one
SELECT id, name, credit_limit, balance FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreditLimit,
		&i.Balance,
	)
	return i, err
}

const getUserForUpdate = `-- name: GetUserForUpdate :one
SELECT id, name, credit_limit, balance FROM users
WHERE id = $1
FOR UPDATE
LIMIT 1
`

func (q *Queries) GetUserForUpdate(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRow(ctx, getUserForUpdate, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreditLimit,
		&i.Balance,
	)
	return i, err
}

const increaseUserBalance = `-- name: IncreaseUserBalance :exec
UPDATE users
SET balance = balance + $1
WHERE id = $2
`

type IncreaseUserBalanceParams struct {
	Balance int32
	ID      int32
}

func (q *Queries) IncreaseUserBalance(ctx context.Context, arg IncreaseUserBalanceParams) error {
	_, err := q.db.Exec(ctx, increaseUserBalance, arg.Balance, arg.ID)
	return err
}

const registerTransaction = `-- name: InsertBalanceTransaction :exec
INSERT INTO transactions (
    user_id,
    value,
    type,
    description
) VALUES (
    $1,
    $2,
    $3,
    $4
)
`

type RegisterTransactionParams struct {
	UserID      pgtype.Int4
	Value       int32
	Type        string
	Description pgtype.Text
}

func (q *Queries) RegisterTransaction(ctx context.Context, arg RegisterTransactionParams) error {
	_, err := q.db.Exec(ctx, registerTransaction,
		arg.UserID,
		arg.Value,
		arg.Type,
		arg.Description,
	)
	return err
}

const updateUserBalance = `-- name: UpdateUserBalance :exec
UPDATE users
SET balance = $1
WHERE id = $2
`

type UpdateUserBalanceParams struct {
	Balance int32
	ID      int32
}

func (q *Queries) UpdateUserBalance(ctx context.Context, arg UpdateUserBalanceParams) error {
	_, err := q.db.Exec(ctx, updateUserBalance, arg.Balance, arg.ID)
	return err
}
