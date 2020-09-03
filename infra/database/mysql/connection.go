package mysql

import "database/sql"

var db *sql.DB

func Connect() *sql.DB {
	var err error
	db, err = sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}
	return db
}

func CloseConn() {
	if db != nil {
		db.Close()
	}
}
