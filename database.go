package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDb() {
	var err error
	db, err = sql.Open("mysql", "root:1234@/desafio_stone?parseTime=true")
  err = db.Ping()
	checkErr(err)
}
