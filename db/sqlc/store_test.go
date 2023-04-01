package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/hoangtk0100/go-healthy/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserTx(t *testing.T) {
	store := NewStore(testDB)

	user1 := createRandomUser(t)
	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	email := user1.Email
	arg := CreateUserParams{
		Username: email,
		Email:    email,
		FullName: util.RandomString(10),
		PhoneNumber: sql.NullString{
			String: util.RandomPhoneNumber(),
			Valid:  true,
		},
		HashedPassword: hashedPassword,
		Type:           UserTypeUSER,
	}

	result, err := store.CreateUserTx(context.Background(), CreateUserTxParams{
		CreateUserParams: arg,
		AfterCreate: func(user User) error {
			_, err := testQueries.CreateUser(context.Background(), arg)
			return err
		},
	})

	require.Error(t, err)
	require.Empty(t, result)
}
