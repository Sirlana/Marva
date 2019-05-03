package db

import (
	"database/sql"
)

type Database interface {
	Connect() error
	GetDB() *sql.DB
	Close() error
	Fetch(sql string) (map[string]interface{}, error)
	Insert(sql string) error
	Update(sql string) error
	Delete(sql string) error
}
