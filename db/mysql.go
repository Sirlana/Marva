package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	username string
	password string
	database string
}

var (
	db  *sql.DB
	err error
)

func NewMySQL(username, password, database string) *MySQL {
	return &MySQL{
		username: username,
		password: password,
		database: database,
	}
}

func (mysql MySQL) Connect() error {
	db, err = sql.Open("mysql", mysql.username+":"+mysql.password+"@/"+mysql.database)
	return err
}

func (mysql MySQL) GetDB() *sql.DB {
	return db
}

func (mysql MySQL) Close() error {
	err = db.Close()
	return err
}

func (mysql MySQL) Fetch(sql string) (map[string]interface{}, error) {
	rows, err := mysql.GetDB().Query(sql)
	if err != nil {
		return nil, err
	}
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	m := make(map[string]interface{})
	for rows.Next() {
		columns := make([]interface{}, len(cols))
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
		}
	}
	return m, nil
}

func (mysql MySQL) Insert(sql string) error {
	stmt, err := mysql.GetDB().Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
func (mysql MySQL) Update(sql string) error {
	stmt, err := mysql.GetDB().Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func (mysql MySQL) Delete(sql string) error {
	stmt, err := mysql.GetDB().Prepare(sql)
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}
