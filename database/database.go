package database

import (
	"database/sql"
	"github.com/gabrielclima/go_rest_api/utils"
)

var db *sql.DB

func initDb() {
	var err error
	db, err = sql.Open("mysql", "root:1234@/invoices?parseTime=true")
	checkErr(err)
	err = db.Ping()
	checkErr(err)
}
