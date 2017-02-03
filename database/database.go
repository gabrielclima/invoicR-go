package database

import (
	"database/sql"
	"github.com/gabrielclima/go_rest_api/utils"
	_ "github.com/go-sql-driver/mysql"
)

var DBCon *sql.DB

// InitDb make the database connection
func init() {
	var err error
	DBCon, err = sql.Open("mysql", "root:1234@/invoices?parseTime=true")
	utils.CheckErr(err)
	err = DBCon.Ping()
	utils.CheckErr(err)
}
