package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/hoangtk0100/go-healthy/util"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {
	hashedPassword, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	email := util.RandomEmail()
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

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.Type, user.Type)
	require.Equal(t, arg.AvatarUrl, user.AvatarUrl)
	require.True(t, user.PasswordChangedAt.IsZero())
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Email, user2.Email)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.PhoneNumber, user2.PhoneNumber)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.Type, user2.Type)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)

	require.WithinDuration(t, user1.PasswordChangedAt, user2.PasswordChangedAt, time.Second)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
