package main

import (
	"testing"
)

var invoice = new(Invoice)

func TestCreateInvoice(t *testing.T) {
	initDb()
	defer db.Close()

	// invoice := new(Invoice)
	invoice.Document = 1111111132
	invoice.Description = "Testando API"
	invoice.Amount = 213123.22
	invoice.ReferenceYear = 2015
	invoice.ReferenceMounth = 2

	i, err := CreateInvoice(invoice)
	if err != nil {
		t.Fatal("Erro no teste de criação : ", err)
	}
	if i != invoice {
		t.Fatal("O invoice não foi criado")
	}
}

func TestGetInvoiceByDoc(t *testing.T) {
	initDb()
	defer db.Close()

	invoice, err := GetInvoiceByDoc(invoice.Document)
	if invoice == (Invoice{}) {
		t.Fatal("Erro ao tentar encontrar um invoice pelo id \n Erro : ", err)
	}
}

func TestGetAllInvoices(t *testing.T){
	initDb()
	defer db.Close()

	invoices, err := GetAllInvoices(nil)
	if invoices == nil {
		t.Fatal("Erro no teste de recuperar todos os invoices", err)
	}
}

func TestDeleteInvoice(t *testing.T){
	initDb()
	defer db.Close()

	deleted, err := DeleteInvoice(invoice.Document)
	if err != nil {
		t.Fatal("Erro no teste de deleção ", err)
	}
	if deleted != "deleted" {
		t.Fatal("Erro na deleção")
	}
	i, err := GetInvoiceByDoc(invoice.Document)
	if i.IsActice == 1 {
		t.Fatal("Erro na deleção", err)
	}
}
