package db

import (
	"context"
	"database/sql"

	"testing"
	"todo/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomCommit(t *testing.T) Commit {

	arg := CreatecommitParams{
		Title:        sql.NullString{String: util.RandomName(), Valid: true},
		Comment:      sql.NullString{String: util.RandomName(), Valid: true},
		Category:     sql.NullString{String: util.RandomCategory(), Valid: true},
		SupervisorID: sql.NullInt32{Int32: int32(12), Valid: true},
		TaskID:       sql.NullInt32{Int32: util.RandomInt(1, 15), Valid: true},
	}

	commit, err := testQueries.Createcommit(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, commit)
	require.Equal(t, arg.Title, commit.Title)

	require.Equal(t, arg.TaskID, commit.TaskID)
	require.Equal(t, arg.Comment, commit.Comment)
	require.Equal(t, arg.Category, commit.Category)

	return commit
}

func TestCreateCommit(t *testing.T) {
	CreateRandomCommit(t)
}

func TestGetCommit(t *testing.T) {
	commit1 := CreateRandomCommit(t)
	commit2, err := testQueries.Getcommit(context.Background(), commit1.CommitID)
	require.NoError(t, err)
	require.NotEmpty(t, commit2)

	require.Equal(t, commit1.CommitID, commit2.CommitID)
	require.Equal(t, commit1.Title, commit2.Title)
	require.Equal(t, commit1.Comment, commit2.Comment)
	require.Equal(t, commit1.Category, commit2.Category)

}
func TestUpdateCommit(t *testing.T) {
	commit1 := CreateRandomCommit(t)

	arg := UpdatecommitParams{
		Title:    sql.NullString{String: util.RandomName(), Valid: true},
		Comment:  sql.NullString{String: util.RandomName(), Valid: true},
		CommitID: commit1.CommitID,
	}

	err := testQueries.Updatecommit(context.Background(), arg)
	require.NoError(t, err)

}

func TestDeleteCommit(t *testing.T) {

	commit1 := CreateRandomCommit(t)
	err := testQueries.Deletecommit(context.Background(), commit1.CommitID)

	require.NoError(t, err)
	commit2, err := testQueries.Getcommit(context.Background(), commit1.CommitID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, commit2)
}

func TestListCommits(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomCommit(t)

	}

	commits, err := testQueries.Listcommits(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, commits)
}
