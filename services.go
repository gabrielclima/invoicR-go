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

func CreateInvoice(invoice *Invoice) *Invoice {
	var sql = "insert into invoices set document=?, description=?, amount=?,"
	sql += " reference_month=?, reference_year=?, created_at=NOW(), is_active=1"
	stmt, err := db.Prepare(sql)
	checkErr(err)

	_, err = stmt.Exec(invoice.Document, invoice.Description, invoice.Amount,
	invoice.ReferenceMounth, invoice.ReferenceYear)

	defer stmt.Close()

	return invoice
}

// func GetInvoiceByDocument(int document) Invoice {}
