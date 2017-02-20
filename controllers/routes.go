package controllers

import (
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/gabrielclima/go_rest_api/auth"
)

// Rotas da API
func Handlers() *negroni.Negroni {
	protectedRoutes := mux.NewRouter().StrictSlash(true)

	protectedRoutes.HandleFunc("/invoices", GetAllInvoicesController).Methods("GET")
	protectedRoutes.HandleFunc("/invoices", CreateInvoiceController).Methods("POST")
	protectedRoutes.HandleFunc("/invoices/{document}", InvoiceByDocController).Methods("GET")
	protectedRoutes.HandleFunc("/invoices/{document}", DeleteInvoiceController).Methods("DELETE")

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.HandlerFunc(auth.Authenticate))
	n.UseHandler(protectedRoutes)

	return n
}
