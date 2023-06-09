// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: user.sql

package db

import (
	"context"
	"database/sql"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    full_name,
    phone_number,
    gender,
    type,
    avatar_url,
    hashed_password
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING username, email, full_name, phone_number, gender, type, status, old_status, avatar_url, hashed_password, password_changed_at, created_at, updated_at, deleted_at
`

type CreateUserParams struct {
	Username       string         `json:"username"`
	Email          string         `json:"email"`
	FullName       string         `json:"full_name"`
	PhoneNumber    sql.NullString `json:"phone_number"`
	Gender         NullGender     `json:"gender"`
	Type           UserType       `json:"type"`
	AvatarUrl      sql.NullString `json:"avatar_url"`
	HashedPassword string         `json:"hashed_password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.FullName,
		arg.PhoneNumber,
		arg.Gender,
		arg.Type,
		arg.AvatarUrl,
		arg.HashedPassword,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.PhoneNumber,
		&i.Gender,
		&i.Type,
		&i.Status,
		&i.OldStatus,
		&i.AvatarUrl,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT username, email, full_name, phone_number, gender, type, status, old_status, avatar_url, hashed_password, password_changed_at, created_at, updated_at, deleted_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.Email,
		&i.FullName,
		&i.PhoneNumber,
		&i.Gender,
		&i.Type,
		&i.Status,
		&i.OldStatus,
		&i.AvatarUrl,
		&i.HashedPassword,
		&i.PasswordChangedAt,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.DeletedAt,
	)
	return i, err
}
