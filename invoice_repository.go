package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"strings"
	// "log"
)

func GetAllInvoices(params map[string][]string) Invoices {
	fmt.Println(params)

	orderBy := strings.Join(params["orderBy"], "")
	year := strings.Join(params["year"], "")
	month := strings.Join(params["month"], "")

	var invoice Invoice
	var invoices Invoices

	var sql = "select * from invoices"

	if year != "" {
		sql += "  where reference_year = " + year

		if month != "" {
			sql += "  and reference_month = " + month
		}

	} else {

		if month != "" {
			sql += " where reference_month = " + month
		}

	}

	if orderBy != "" {
		sql += " order by " + orderBy
	}

	fmt.Println(sql)

	rows, err := db.Query(sql)
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

func GetInvoiceByDoc(document int) Invoice {
	var invoice Invoice
	stmt, err := db.Prepare("select * from invoices where document=?")
	checkErr(err)
	// row := db.QueryRow("select * from invoices where document=?;", document)
	err = stmt.QueryRow(document).Scan(&invoice.Document, &invoice.Description, &invoice.Amount,
		&invoice.ReferenceMounth, &invoice.ReferenceYear,
		&invoice.CreatedAt, &invoice.IsActice, &invoice.DesactiveAt)
	if err == sql.ErrNoRows {
		return invoice
	}
	checkErr(err)
	defer stmt.Close()

	return invoice
}

func CreateInvoice(invoice *Invoice) *Invoice {
	var sql = "insert into invoices set document=?, description=?, amount=?,"
	sql += " reference_month=?, reference_year=?, created_at=NOW(), is_active=1"
	stmt, err := db.Prepare(sql)
	checkErr(err)

	_, err = stmt.Exec(invoice.Document, invoice.Description, invoice.Amount,
		invoice.ReferenceMounth, invoice.ReferenceYear)
	checkErr(err)

	defer stmt.Close()

	return invoice
}

func DeleteInvoice(document int) string {
	stmt, err := db.Prepare("update invoices set is_active=0, desactive_at=NOW() where document = ?")
	checkErr(err)

	_, err = stmt.Exec(document)
	checkErr(err)
	defer stmt.Close()

	return "deleted"
}

// func GetInvoiceByDocument(int document) Invoice {}