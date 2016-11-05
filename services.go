package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllInvoices() Invoices {
	var invoice Invoice
	var invoices Invoices

	var sql = "select * from invoices"
	sql += " where document = ?"
  fmt.Println(sql)
	rows, err := db.Query(sql, "12312310")
	checkErr(err)

	for rows.Next() {
		err = rows.Scan(&invoice.Document, &invoice.Description, &invoice.Amount,
			&invoice.ReferenceMounth, &invoice.ReferenceYear,
			&invoice.CreatedAt, &invoice.IsActice, &invoice.DesactiveAt)
		invoices = append(invoices, invoice)
		checkErr(err)
	}

	defer rows.Close()

	return invoices
}

// func GetInvoiceByDocument(int document) Invoice {}
