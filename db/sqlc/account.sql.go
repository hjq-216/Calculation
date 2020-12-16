// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
)

const createAccount = `-- name: CreateAccount :one
INSERT INTO accounts (
  owner,
  message,
  mistakes,
  password
) VALUES (
  $1, $2, $3, $4
) RETURNING id, owner, message, mistakes, password, created_at
`

type CreateAccountParams struct {
	Owner    string `json:"owner"`
	Message  string `json:"message"`
	Mistakes string `json:"mistakes"`
	Password string `json:"password"`
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, createAccount,
		arg.Owner,
		arg.Message,
		arg.Mistakes,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Message,
		&i.Mistakes,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE owner = $1
`

func (q *Queries) DeleteAccount(ctx context.Context, owner string) error {
	_, err := q.db.ExecContext(ctx, deleteAccount, owner)
	return err
}

const getAccount = `-- name: GetAccount :one
SELECT id, owner, message, mistakes, password, created_at FROM accounts
WHERE owner = $1 LIMIT 1
`

func (q *Queries) GetAccount(ctx context.Context, owner string) (Account, error) {
	row := q.db.QueryRowContext(ctx, getAccount, owner)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Message,
		&i.Mistakes,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id, owner, message, mistakes, password, created_at FROM accounts
ORDER BY owner
LIMIT $1
OFFSET $2
`

type ListAccountsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error) {
	rows, err := q.db.QueryContext(ctx, listAccounts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Account{}
	for rows.Next() {
		var i Account
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Message,
			&i.Mistakes,
			&i.Password,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAccount = `-- name: UpdateAccount :one
UPDATE accounts
SET message = $2,mistakes = $3,password = $4
WHERE owner = $1
RETURNING id, owner, message, mistakes, password, created_at
`

type UpdateAccountParams struct {
	Owner    string `json:"owner"`
	Message  string `json:"message"`
	Mistakes string `json:"mistakes"`
	Password string `json:"password"`
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error) {
	row := q.db.QueryRowContext(ctx, updateAccount,
		arg.Owner,
		arg.Message,
		arg.Mistakes,
		arg.Password,
	)
	var i Account
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Message,
		&i.Mistakes,
		&i.Password,
		&i.CreatedAt,
	)
	return i, err
}
