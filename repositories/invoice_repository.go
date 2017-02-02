package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

func GetAllInvoices(params map[string][]string) (Invoices, error) {
	orderBy := strings.Join(params["orderBy"], "")
	year := strings.Join(params["year"], "")
	month := strings.Join(params["month"], "")
	limit := strings.Join(params["limit"], "")
	offset := strings.Join(params["offset"], "")

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
	if limit != "" {
		sql += " limit " + limit
		if offset != "" {
			sql += " offset " + offset
		}
	}

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

	return invoices, err
}

func GetInvoiceByDoc(document int) (Invoice, error) {
	var invoice Invoice
	stmt, err := db.Prepare("select * from invoices where document=?")
	checkErr(err)

	err = stmt.QueryRow(document).Scan(&invoice.Document, &invoice.Description, &invoice.Amount,
		&invoice.ReferenceMounth, &invoice.ReferenceYear,
		&invoice.CreatedAt, &invoice.IsActice, &invoice.DesactiveAt)
	if err == sql.ErrNoRows {
		return invoice, err
	}
	checkErr(err)
	defer stmt.Close()

	return invoice, err
}

func CreateInvoice(invoice *Invoice) (*Invoice, error) {
	var sql = "insert into invoices set document=?, description=?, amount=?,"
	sql += " reference_month=?, reference_year=?, created_at=NOW(), is_active=1"
	stmt, err := db.Prepare(sql)
	checkErr(err)

	_, err = stmt.Exec(invoice.Document, invoice.Description, invoice.Amount,
		invoice.ReferenceMounth, invoice.ReferenceYear)
	if err != nil {
		return invoice, err
	}

	defer stmt.Close()

	return invoice, err
}

func DeleteInvoice(document int) (string, error) {
	stmt, err := db.Prepare("update invoices set is_active=0, desactive_at=NOW() where document = ?")
	checkErr(err)

	_, err = stmt.Exec(document)
	checkErr(err)
	defer stmt.Close()

	return "deleted", err
}
