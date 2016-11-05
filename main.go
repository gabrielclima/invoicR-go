package main

import (
	// "fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	initDb()
	defer db.Close()
	router := mux.NewRouter().StrictSlash(true)

	// declarando rotas seus respectivos métodos http e as funções a serem executadas
	router.HandleFunc("/invoices", RestInvoices).Methods("GET")
	// router.HandleFunc("/invoices", RestCreateInvoice).Methods("POST")
	// router.HandleFunc("/invoices", RestInvoiceByDoc).Methods("GET")
	// router.HandleFunc("/invoices/:document", RestDeleteInvoices).Methods("DELETE")

	// colocando a API online
	log.Fatal(http.ListenAndServe(":8080", router))
}
