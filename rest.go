package main

import (
	"encoding/json"
	// "fmt"
	// "log"
	"io/ioutil"
	"net/http"
)

func RestInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	invoices := GetAllInvoices()
	res, err := json.Marshal(invoices)
	checkErr(err)

	if invoices == nil {
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(jsonErr{Code: http.StatusNotFound, Text: "Not Found"})
		checkErr(err)
	}

	w.Write(res)
}

type test_struct struct {
	Test string
}

func RestCreateInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var err error
	var res []byte

	invoice := new(Invoice)
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	err = json.Unmarshal(body, &invoice)
	checkErr(err)

	inv := GetInvoiceByDoc(invoice.Document)

	if inv != (Invoice{}) {
		w.WriteHeader(http.StatusConflict)
		res, err = json.Marshal(jsonErr{Code: http.StatusConflict,
			Text: "Já existe um documento com esse número"})
		checkErr(err)
    w.Write(res)

    return
  } else {
    inv := CreateInvoice(invoice)
    res, err = json.Marshal(inv)
    checkErr(err)
    w.WriteHeader(http.StatusCreated)

    w.Write(res)
  }
}
