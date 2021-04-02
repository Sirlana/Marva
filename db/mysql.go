package db

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

type MySQL struct {
	Username string
	Password string
	Database string
}

var (
	db  *sql.DB
	err error
)

func NewMySQL(username, password, database string) *MySQL {
	return &MySQL{
		Username: username,
		Password: password,
		Database: database,
	}
}

func (mysql MySQL) Connect() error {
	db, err = sql.Open("mysql", mysql.Username+":"+mysql.Password+"@/"+mysql.Database)
	return err
}

func (mysql MySQL) GetDB() *sql.DB {
	return db
}

func (mysql MySQL) Close() error {
	err = db.Close()
	return err
}
