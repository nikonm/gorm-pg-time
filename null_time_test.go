package gormpgtime

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNullTime(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTime
	}
	tm, _ := time.Parse("15:04:05", "13:07:23")
	m := &Model{
		ID:   0,
		Name: "test",
		Time: (&NullTime{}).Set(&tm),
	}
	db := (&Storage{}).Init()
	db.Mock.ExpectQuery(`INSERT INTO .*`).
		WithArgs(sqlmock.AnyArg(), tm.Format("15:04:05")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(1))

	err := db.Db().Save(m).Error

	require.NoError(t, err)
}

func TestNullTime_Value(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTime
	}

	db := (&Storage{}).Init()
	db.Mock.ExpectQuery(`SELECT .*`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "time"}).
				AddRow(1, "test", nil))

	m := &Model{}
	err := db.Db().First(m).Error
	require.NoError(t, err)
	require.Equal(t, false, m.Time.Valid)
}
