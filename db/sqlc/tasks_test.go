package db

import (
	"context"
	"database/sql"
	"time"

	"testing"
	"todo/util"

	"github.com/stretchr/testify/require"
)

func CreateRandomTask(t *testing.T) Task {

	arg := CreatetaskParams{
		Name:       sql.NullString{String: util.RandomName(), Valid: true},
		Supervisor: sql.NullInt32{Int32: util.RandomInt(1, 10)},
		Category:   sql.NullString{String: util.RandomCategory(), Valid: true},
	}

	task, err := testQueries.Createtask(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, task)
	require.Equal(t, arg.Name, task.Name)

	require.Equal(t, arg.Category, task.Category)

	return task
}

func TestCreateTask(t *testing.T) {
	CreateRandomTask(t)
}

func TestGetTask(t *testing.T) {
	task1 := CreateRandomTask(t)
	task2, err := testQueries.Gettask(context.Background(), task1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, task2)

	require.Equal(t, task1.ID, task2.ID)
	require.Equal(t, task1.Name, task2.Name)
	require.Equal(t, task1.Supervisor, task2.Supervisor)
	require.WithinDuration(t, task1.CreatedAt, task2.CreatedAt, time.Second)

}
func TestUpdateTask(t *testing.T) {
	task1 := CreateRandomTask(t)

	arg := UpdatetaskParams{
		Name:       sql.NullString{String: util.RandomName(), Valid: true},
		Supervisor: sql.NullInt32{Int32: util.RandomInt(1, 10)},

		ID: task1.ID,
	}

	err := testQueries.Updatetask(context.Background(), arg)
	require.NoError(t, err)

}
func TestDeleteTask(t *testing.T) {

	task1 := CreateRandomTask(t)
	err := testQueries.Deletetask(context.Background(), task1.ID)

	require.NoError(t, err)
	task2, err := testQueries.Gettask(context.Background(), task1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, task2)
}

func TestListTask(t *testing.T) {
	for i := 0; i < 10; i++ {
		CreateRandomTask(t)

	}

	tasks, err := testQueries.Listtasks(context.Background())
	require.NoError(t, err)
	require.NotEmpty(t, tasks)
}
