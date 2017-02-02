package main

import (
	"github.com/gorilla/mux"
)

func Handlers() *mux.Router {
  router := mux.NewRouter().StrictSlash(true)

  router.HandleFunc("/invoices", RestInvoices).Methods("GET")
	router.HandleFunc("/invoices", RestCreateInvoice).Methods("POST")
	router.HandleFunc("/invoices/{document}", RestInvoiceByDoc).Methods("GET")
	router.HandleFunc("/invoices/{document}", RestDeleteInvoice).Methods("DELETE")

  return router
}
