package tests

import (
	"testing"
)

var invoice = new(Invoice)

func TestCreateInvoice(t *testing.T) {
	initDb()
	defer db.Close()

	// invoice := new(Invoice)
	invoice.Document = 21466713311348
	invoice.Description = "Testando API"
	invoice.Amount = 213123.22
	invoice.ReferenceYear = 2015
	invoice.ReferenceMounth = 2

	i, err := CreateInvoice(invoice)
	if err != nil {
		t.Error("Erro no teste de criação : ", err)
	}
	if i != invoice {
		t.Error("O invoice não foi criado")
	}
}

func TestGetInvoiceByDoc(t *testing.T) {
	initDb()
	defer db.Close()

	invoice, err := GetInvoiceByDoc(invoice.Document)
	if invoice == (Invoice{}) {
		t.Error("Erro ao tentar encontrar um invoice pelo id \n Erro : ", err)
	}
}

func TestGetAllInvoices(t *testing.T){
	initDb()
	defer db.Close()

	invoices, err := GetAllInvoices(nil)
	if invoices == nil {
		t.Error("Erro no teste de recuperar todos os invoices", err)
	}
}

func TestDeleteInvoice(t *testing.T){
	initDb()
	defer db.Close()

	deleted, err := DeleteInvoice(invoice.Document)
	if err != nil {
		t.Error("Erro no teste de deleção ", err)
	}
	if deleted != "deleted" {
		t.Error("Erro na deleção")
	}
	i, err := GetInvoiceByDoc(invoice.Document)
	if i.IsActice == 1 {
		t.Error("Erro na deleção", err)
	}
}
