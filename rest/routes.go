package rest

import (
	"github.com/gorilla/mux"
)

// Handlers - all routes
func Handlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", ApiResource).Methods("GET")

	router.HandleFunc("/invoices", InvoicesResource).Methods("GET")
	router.HandleFunc("/invoices", CreateInvoiceResource).Methods("POST")
	router.HandleFunc("/invoices/{document}", InvoiceByDocResource).Methods("GET")
	router.HandleFunc("/invoices/{document}", DeleteInvoiceResource).Methods("DELETE")

	return router
}
