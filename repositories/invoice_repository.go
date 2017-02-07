package repositories

import (
	"database/sql"
	"fmt"
	db "github.com/gabrielclima/go_rest_api/database"
	"github.com/gabrielclima/go_rest_api/models"
	"github.com/gabrielclima/go_rest_api/utils"
	"strings"
)

var DBCon *sql.DB

func init() {
	DBCon = db.DBCon
}

func GetAllInvoices(params map[string][]string) (models.Invoices, error) {
	orderBy := strings.Join(params["orderBy"], "")
	year := strings.Join(params["year"], "")
	month := strings.Join(params["month"], "")
	limit := strings.Join(params["limit"], "")
	offset := strings.Join(params["offset"], "")

	var invoice models.Invoice
	var invoices models.Invoices

	var sql = "select i.id, i.document, i.description, i.amount, " +
		"i.reference_month, i.reference_year, i.created_at, i.is_active  " +
		"from invoices i "

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

	if year == "" && month == "" {
		sql += " where is_active = 1 "
	} else {
		sql += " and is_active = 1"
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

	fmt.Printf(sql)

	rows, err := DBCon.Query(sql)
	utils.CheckErr(err)

	for rows.Next() {
		err = rows.Scan(&invoice.Id, &invoice.Document, &invoice.Description,
			&invoice.Amount, &invoice.ReferenceMounth, &invoice.ReferenceYear,
			&invoice.CreatedAt, &invoice.IsActice)
		invoices = append(invoices, invoice)
		utils.CheckErr(err)
	}

	defer rows.Close()

	return invoices, err
}

func GetInvoiceByDoc(document int) (models.Invoice, error) {
	var invoice models.Invoice
	stmt, err := DBCon.Prepare("select i.id, i.document, i.description, i.amount, " +
		"i.reference_month, i.reference_year, i.created_at, i.is_active  " +
		"from invoices i " +
		"where document=? and is_active ")
	utils.CheckErr(err)

	err = stmt.QueryRow(document).Scan(&invoice.Id, &invoice.Document,
		&invoice.Description, &invoice.Amount,
		&invoice.ReferenceMounth, &invoice.ReferenceYear,
		&invoice.CreatedAt, &invoice.IsActice)

	if err == sql.ErrNoRows {
		return invoice, err
	}

	utils.CheckErr(err)
	defer stmt.Close()

	return invoice, err
}

func CreateInvoice(invoice *models.Invoice) (*models.Invoice, error) {
	var sql = "insert into invoices set document=?, description=?, amount=?, " +
		"reference_month=?, reference_year=?, created_at=NOW(), is_active=1, desactive_at='0000-00-00 00:00:00'"

	stmt, err := DBCon.Prepare(sql)
	utils.CheckErr(err)

	_, err = stmt.Exec(invoice.Document, invoice.Description, invoice.Amount,
		invoice.ReferenceMounth, invoice.ReferenceYear)
	if err != nil {
		return invoice, err
	}

	defer stmt.Close()

	return invoice, err
}

func DeleteInvoice(document int) (bool, error) {
	stmt, err := DBCon.Prepare("update invoices set is_active=0, desactive_at=NOW() where document = ?")
	utils.CheckErr(err)

	_, err = stmt.Exec(document)
	utils.CheckErr(err)
	defer stmt.Close()

	return true, err
}
