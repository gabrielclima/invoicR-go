package controllers

import (
	"github.com/gorilla/mux"
)

// Handlers - all routes
func Handlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", ApiController).Methods("GET")

	router.HandleFunc("/invoices", GetAllInvoicesController).Methods("GET")
	router.HandleFunc("/invoices", CreateInvoiceController).Methods("POST")
	router.HandleFunc("/invoices/{document}", InvoiceByDocController).Methods("GET")
	router.HandleFunc("/invoices/{document}", DeleteInvoiceController).Methods("DELETE")

	return router
}
