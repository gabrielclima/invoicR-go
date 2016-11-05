package main

import (
	"net/http"
	"encoding/json"
)

func RestInvoices(w http.ResponseWriter, r *http.Request) {
	invoices := GetAllInvoices()
	res, err := json.Marshal(invoices)
	checkErr(err)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	if invoices == nil{
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(jsonErr{Code: http.StatusNotFound, Text: "Not Found"})
		checkErr(err)
	}

	w.Write(res)
}

func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}
