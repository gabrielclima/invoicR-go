package controllers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

// Handlers map all routes
func Handlers() *negroni.Negroni {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", ApiController).Methods("GET")

	r.HandleFunc("/invoices", GetAllInvoicesController).Methods("GET")
	r.HandleFunc("/invoices", CreateInvoiceController).Methods("POST")
	r.HandleFunc("/invoices/{document}", InvoiceByDocController).Methods("GET")
	r.HandleFunc("/invoices/{document}", DeleteInvoiceController).Methods("DELETE")

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	return n
}
