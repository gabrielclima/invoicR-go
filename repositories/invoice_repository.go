package repositories

import (
	"database/sql"
	db "github.com/gabrielclima/go_rest_api/database"
	"github.com/gabrielclima/go_rest_api/domain"
	"github.com/gabrielclima/go_rest_api/utils"
	"strings"
)

func GetAllInvoices(params map[string][]string) (domain.Invoices, error) {
	orderBy := strings.Join(params["orderBy"], "")
	year := strings.Join(params["year"], "")
	month := strings.Join(params["month"], "")
	limit := strings.Join(params["limit"], "")
	offset := strings.Join(params["offset"], "")

	var invoice domain.Invoice
	var invoices domain.Invoices

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

	rows, err := db.DBCon.Query(sql)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&invoice.Document, &invoice.Description, &invoice.Amount,
			&invoice.ReferenceMounth, &invoice.ReferenceYear,
			&invoice.CreatedAt, &invoice.IsActice, &invoice.DesactiveAt)
		invoices = append(invoices, invoice)
		utils.CheckErr(err)
	}

	defer rows.Close()

	return invoices, err
}

func GetInvoiceByDoc(document int) (domain.Invoice, error) {
	var invoice domain.Invoice
	stmt, err := db.DBCon.Prepare("select * from invoices where document=?")
	utils.CheckErr(err)

	err = stmt.QueryRow(document).Scan(&invoice.Id, &invoice.Document, &invoice.Description, &invoice.Amount,
		&invoice.ReferenceMounth, &invoice.ReferenceYear,
		&invoice.CreatedAt, &invoice.IsActice, &invoice.DesactiveAt)
	if err == sql.ErrNoRows {
		return invoice, err
	}
	utils.CheckErr(err)
	defer stmt.Close()

	return invoice, err
}

func CreateInvoice(invoice *domain.Invoice) (*domain.Invoice, error) {
	var sql = "insert into invoices set document=?, description=?, amount=?,"
	sql += " reference_month=?, reference_year=?, created_at=NOW(), is_active=1"
	stmt, err := db.DBCon.Prepare(sql)
	utils.CheckErr(err)

	_, err = stmt.Exec(invoice.Document, invoice.Description, invoice.Amount,
		invoice.ReferenceMounth, invoice.ReferenceYear)
	if err != nil {
		return invoice, err
	}

	defer stmt.Close()

	return invoice, err
}

func DeleteInvoice(document int) (string, error) {
	stmt, err := db.DBCon.Prepare("update invoices set is_active=0, desactive_at=NOW() where document = ?")
	utils.CheckErr(err)

	_, err = stmt.Exec(document)
	utils.CheckErr(err)
	defer stmt.Close()

	return "deleted", err
}
