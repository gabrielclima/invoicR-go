package controllers

import (
	"database/sql"
	"encoding/json"
	"github.com/gabrielclima/go_rest_api/models"
	"github.com/gabrielclima/go_rest_api/repositories"
	"github.com/gabrielclima/go_rest_api/utils"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"log"
)

// Constante usada para setar todos os Content-Type
const ApplicationJSON = "application/json; charset=UTF-8"

// GET /invoices
// Controller que retorna todos os Invoices ativos na base de dados
// Retorna o status code 200 com [] caso não tenha nenhum e 200 com uma lista de
// Invoices, possua no banco de dados
func GetAllInvoicesController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)

	res := []byte(`[]`)
	var err error

	params := r.URL.Query()
	invoices, err := repositories.GetAllInvoices(params)
	if err != nil {
		log.Println(err)
	}
	if invoices != nil {
		res, err = json.Marshal(invoices)
		if err != nil {
			log.Println(err)
		}
	}

	w.Write(res)
}

// GET /invoice/{document}
// Controller responsável por retornar um invoice baseado no
// parâmetro "document" passado no PATH da requisição
// Retorna 200 se encontrar algum Invoice com o número de documento
// correspondente e 404 se não encontrar
func InvoiceByDocController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)
	vars := mux.Vars(r)
	var res []byte
	var err error

	var document int

	document, err = strconv.Atoi(vars["document"])
	if err != nil {
		log.Println(err)
	}

	invoice, err := repositories.GetInvoiceByDoc(document)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			res, err = json.Marshal(utils.JsonErr{Code: http.StatusNotFound,
				Message: "Não foi encontrada nenhuma nota fiscal com esse número de documento."})
		} else {
			log.Println(err)
		}
	}

	if invoice != (models.Invoice{}) {
		res, err = json.Marshal(invoice)
		log.Println(err)
	}

	w.Write(res)
}

// POST /invoices
// Controller responsável pela criação de um Invoice passado pelo body da requisição
// Ex:
// {
//   "document":"1231241231",
//   "description":"Uma nota fiscal qualquer",
//   "amount": "123.00",
//   "reference_mounth":"12",
//   "reference_year":"2014"
// }
// Retorna 201 caso seja criado com sucesso e 409 caso o número de documento já exista na base
func CreateInvoiceController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)
	var err error
	var res []byte

	invoice := new(models.Invoice)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(body, &invoice)
	if err != nil {
		log.Println(err)
	}

	inv, err := repositories.GetInvoiceByDoc(invoice.Document)

	if inv != (models.Invoice{}) {
		w.WriteHeader(http.StatusConflict)
		res, err = json.Marshal(utils.JsonErr{Code: http.StatusConflict,
			Message: "Já existe uma nota fiscal com esse número"})
		if err != nil {
			log.Println(err)
		}
	} else {
		inv, err := repositories.CreateInvoice(invoice)
		if err != nil {
			log.Println(err)
		}
		invoiceCreated, err := repositories.GetInvoiceByDoc(inv.Document)
		if err != nil {
			log.Println(err)
		}
		res, err = json.Marshal(invoiceCreated)
		if err != nil {
			log.Println(err)
		}
		w.WriteHeader(http.StatusCreated)
	}

	w.Write(res)
}

// DELETE /invoices/{document}
// Controller responsável por fazer um soft delete, ou seja, apenas setar o status
// para inativo baseado no documento que foi passado no path da requisição
// Retorna 200 caso seja deletado com sucesso e 404 caso não tenha achado
func DeleteInvoiceController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", ApplicationJSON)

	vars := mux.Vars(r)
	var res []byte
	var err error

	var document int

	document, err = strconv.Atoi(vars["document"])
	if err != nil {
		log.Println(err)
	}

	invoice, err := repositories.GetInvoiceByDoc(document)

	if invoice == (models.Invoice{}) {
		w.WriteHeader(http.StatusNotFound)
		res, err = json.Marshal(utils.JsonErr{Code: http.StatusNotFound,
			Message: "Não foi encontrada nenhuma nota fiscal ativa com esse número de documento."})
	} else {
		deleted, err := repositories.DeleteInvoice(document)
		if err != nil {
			log.Println(err)
		}
		if deleted {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
	w.Write(res)
}
