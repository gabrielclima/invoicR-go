package database

import (
	"database/sql"
	utils "github.com/gabrielclima/go_rest_api/utils"
)

var db *sql.DB

// InitDb make the database connection
func InitDb() {
	var err error
	db, err = sql.Open("mysql", "root:1234@/invoices?parseTime=true")
	utils.CheckErr(err)
	err = db.Ping()
	utils.CheckErr(err)
	defer db.Close()
}
