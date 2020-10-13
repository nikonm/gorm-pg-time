package gormpgtime

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	Mock sqlmock.Sqlmock
	db   *gorm.DB
}

func (s *Storage) Init() *Storage {
	db, mock, _ := sqlmock.New()
	s.Mock = mock
	s.db, _ = gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{SkipDefaultTransaction: true})
	return s
}

func (s *Storage) Db() *gorm.DB {
	return s.db
}
