package tests

import (
	"testing"
	"github.com/gabrielclima/go_rest_api/domain"
	db "github.com/gabrielclima/go_rest_api/database"
	"github.com/gabrielclima/go_rest_api/repositories"
)

var invoice = new(domain.Invoice)

func TestCreateInvoice(t *testing.T) {
	db.InitDb()
	defer db.DBCon.Close()

	// invoice := new(Invoice)
	invoice.Document = 21466713311348
	invoice.Description = "Testando API"
	invoice.Amount = 213123.22
	invoice.ReferenceYear = 2015
	invoice.ReferenceMounth = 2

	i, err := repositories.CreateInvoice(invoice)
	if err != nil {
		t.Error("Erro no teste de criação : ", err)
	}
	if i != invoice {
		t.Error("O invoice não foi criado")
	}
}

func TestGetInvoiceByDoc(t *testing.T) {
	db.InitDb()
	defer db.DBCon.Close()

	invoice, err := repositories.GetInvoiceByDoc(invoice.Document)
	if invoice == (domain.Invoice{}) {
		t.Error("Erro ao tentar encontrar um invoice pelo id \n Erro : ", err)
	}
}

func TestGetAllInvoices(t *testing.T){
	db.InitDb()
	defer db.DBCon.Close()

	invoices, err := repositories.GetAllInvoices(nil)
	if invoices == nil {
		t.Error("Erro no teste de recuperar todos os invoices", err)
	}
}

func TestDeleteInvoice(t *testing.T){
	db.InitDb()
	defer db.DBCon.Close()

	deleted, err := repositories.DeleteInvoice(invoice.Document)
	if err != nil {
		t.Error("Erro no teste de deleção ", err)
	}
	if deleted != "deleted" {
		t.Error("Erro na deleção")
	}
	i, err := repositories.GetInvoiceByDoc(invoice.Document)
	if i.IsActice == 1 {
		t.Error("Erro na deleção", err)
	}
}
