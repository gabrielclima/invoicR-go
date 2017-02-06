package rest

import (
	"encoding/json"
	auth "github.com/gabrielclima/go_rest_api/auth"
	. "github.com/gabrielclima/go_rest_api/domain"
	"github.com/gabrielclima/go_rest_api/repositories"
	utils "github.com/gabrielclima/go_rest_api/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

// ApplicationJSON const for used in all Headers setting a Content-Type
const ApplicationJSON = "application/json; charset=UTF-8"

// InvoicesResource returns all actives invoices
func InvoicesResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)

	res := []byte(`[]`)
	var err error

	authenticate := auth.Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	params := r.URL.Query()
	invoices, err := repositories.GetAllInvoices(params)
	utils.CheckErr(err)
	if invoices != nil {
		res, err = json.Marshal(invoices)
		utils.CheckErr(err)
	}

	w.Write(res)
}

// InvoiceByDocResource return a the invoices parsed in request path
func InvoiceByDocResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)
	vars := mux.Vars(r)
	var res []byte
	var err error

	authenticate := auth.Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	var document int

	document, err = strconv.Atoi(vars["document"])
	utils.CheckErr(err)

	invoice, err := repositories.GetInvoiceByDoc(document)
	utils.CheckErr(err)

	if invoice != (Invoice{}) {
		res, err = json.Marshal(invoice)
		utils.CheckErr(err)
	} else {
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(utils.JsonErr{Code: http.StatusNotFound, Text: "Not Found"})
	}

	w.Write(res)
}

// CreateInvoiceResource create a invoice based on JSON body parsed in request
func CreateInvoiceResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)
	var err error
	var res []byte

	authenticate := auth.Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	invoice := new(Invoice)
	body, err := ioutil.ReadAll(r.Body)
	utils.CheckErr(err)

	err = json.Unmarshal(body, &invoice)
	utils.CheckErr(err)

	inv, err := repositories.GetInvoiceByDoc(invoice.Document)

	if inv != (Invoice{}) {
		w.WriteHeader(http.StatusConflict)
		res, err = json.Marshal(utils.JsonErr{Code: http.StatusConflict,
			Text: "Já existe um documento com esse número"})
		utils.CheckErr(err)
	} else {
		inv, err := repositories.CreateInvoice(invoice)
		utils.CheckErr(err)
		invoiceCreated, err := repositories.GetInvoiceByDoc(inv.Document)
		utils.CheckErr(err)
		res, err = json.Marshal(invoiceCreated)
		utils.CheckErr(err)
		w.WriteHeader(http.StatusCreated)
	}

	w.Write(res)
}

// DeleteInvoiceResource do a soft delete on a invoice parsed in request path
func DeleteInvoiceResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)

	vars := mux.Vars(r)
	var res []byte
	var err error

	authenticate := auth.Authenticate(w, r)
	if authenticate == http.StatusUnauthorized {
		return
	}

	var document int

	document, err = strconv.Atoi(vars["document"])
	utils.CheckErr(err)

	invoice, err := repositories.GetInvoiceByDoc(document)
	utils.CheckErr(err)

	if invoice == (Invoice{}) {
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(utils.JsonErr{Code: http.StatusNotFound, Text: "Not Found"})
	} else {

		deleted, err := repositories.DeleteInvoice(document)
		utils.CheckErr(err)
		if deleted  {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
	w.Write(res)
}
