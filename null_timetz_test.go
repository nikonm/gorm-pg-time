package gormpgtime

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestNullTimeTZ(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTimeTZ
	}
	tm, _ := time.ParseInLocation("15:04:05", "13:07:23", time.FixedZone("", 0))
	m := &Model{
		ID:   0,
		Name: "test",
		Time: (&NullTimeTZ{}).Set(&tm),
	}
	db := (&Storage{}).Init()
	db.Mock.ExpectQuery(`INSERT INTO .*`).
		WithArgs(sqlmock.AnyArg(), tm.Format("15:04:05")+"+00:00").
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(1))

	err := db.Db().Save(m).Error

	require.NoError(t, err)
}

func TestNullTimeTZPositive(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTimeTZ
	}
	tm, _ := time.ParseInLocation("15:04:05", "13:07:23", time.FixedZone("", 3600))
	m := &Model{
		ID:   0,
		Name: "test",
		Time: (&NullTimeTZ{}).Set(&tm),
	}
	db := (&Storage{}).Init()
	db.Mock.ExpectQuery(`INSERT INTO .*`).
		WithArgs(sqlmock.AnyArg(), tm.Format("15:04:05")+"+01:00").
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(1))

	err := db.Db().Save(m).Error

	require.NoError(t, err)
}

func TestNullTimeTZNegative(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTimeTZ
	}
	tm, _ := time.ParseInLocation("15:04:05", "13:07:23", time.FixedZone("", -3600))
	m := &Model{
		ID:   0,
		Name: "test",
		Time: (&NullTimeTZ{}).Set(&tm),
	}
	db := (&Storage{}).Init()
	db.Mock.ExpectQuery(`INSERT INTO .*`).
		WithArgs(sqlmock.AnyArg(), tm.Format("15:04:05")+"-01:00").
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).
				AddRow(1))

	err := db.Db().Save(m).Error

	require.NoError(t, err)
}

func TestNullTimeTZ_Value(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTimeTZ
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

func TestNullTimeTZ_Scan(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTimeTZ
	}

	tms := "15:42:31"
	tz := "+03:00"
	db := (&Storage{}).Init()
	db.Mock.ExpectQuery(`SELECT .*`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "time"}).
				AddRow(1, "test", tms+tz))

	m := &Model{}
	err := db.Db().First(m).Error
	require.NoError(t, err)
	require.Equal(t, true, m.Time.Valid)
	require.Equal(t, tms, m.Time.Time.Format("15:04:05"))
	_, offset := m.Time.Time.Zone()
	require.Equal(t, 3600*3, offset)
}

func TestNullTimeTZNegative_Scan(t *testing.T) {
	type Model struct {
		ID   uint
		Name string
		Time NullTimeTZ
	}

	tms := "15:42:31"
	tz := "-03:00"
	db := (&Storage{}).Init()
	db.Mock.ExpectQuery(`SELECT .*`).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "name", "time"}).
				AddRow(1, "test", tms+tz))

	m := &Model{}
	err := db.Db().First(m).Error
	require.NoError(t, err)
	require.Equal(t, true, m.Time.Valid)
	require.Equal(t, tms, m.Time.Time.Format("15:04:05"))
	_, offset := m.Time.Time.Zone()
	require.Equal(t, -3600*3, offset)
}
