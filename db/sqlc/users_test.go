package db

import (
	"context"
	"database/sql"
	"time"

	"testing"
	"todo/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomAccount(t *testing.T) User {

	arg := CreateuserParams{
		FullName:   sql.NullString{String: util.RandomName(), Valid: true},
		Supervisor: sql.NullBool{Bool: true, Valid: true},
	}

	user, err := testQueries.Createuser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Supervisor, user.Supervisor)
	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)

	return user
}

func TestCreateAccount(t *testing.T) {
	CreateRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)
	account2, err := testQueries.Getuser(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.FullName, account2.FullName)
	require.Equal(t, account1.Supervisor, account2.Supervisor)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateRandomAccount(t)

	arg := UpdateuserParams{
		ID:         account1.ID,
		FullName:   sql.NullString{String: util.RandomName(), Valid: true},
		Supervisor: account1.Supervisor,
	}

	err := testQueries.Updateuser(context.Background(), arg)
	require.NoError(t, err)

}

func TestDeleteAccount(t *testing.T) {

	account1 := CreateRandomAccount(t)
	err := testQueries.Deleteuser(context.Background(), account1.ID)

	require.NoError(t, err)
	account2, err := testQueries.Getuser(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomAccount(t)

	}

	accounts, err := testQueries.Listusers(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, accounts)
}
