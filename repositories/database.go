package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb() {
	var err error
	db, err = sql.Open("mysql", "root:1234@/invoices?parseTime=true")
	checkErr(err)
  err = db.Ping()
	checkErr(err)
}
