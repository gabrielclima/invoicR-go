package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// DBCon storage my database connection
var DBCon *sql.DB

// Init make the database connection
func init() {
	var err error
	DBCon, err = sql.Open("mysql", "root:1234@/invoices?parseTime=true")
	if err != nil {
		log.Panic("Erro na conexão com o banco de dados", err)
	}
	err = DBCon.Ping()
	if err != nil {
		log.Panic("Erro na conexão com o banco de dados", err)
	}
}
