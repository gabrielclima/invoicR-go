package rest

import (
	"fmt"
	"net/http"
)

const TextHTML = "text/html; charset=UTF-8"

func ApiResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", TextHTML)
	text := "<h2>Golang REST API</h2>"
	text += "<span>Essa é uma API RESTFull escrita em Golang.</span>"
	text += "<span>A documentação contendo todas as informações" +
		"e as rotas disponíveis se encontra em " +
		"<a href=\"https://github.com/gabrielclima/go_rest_api\">" +
		"github.com/gabrielclima/go_rest_api</a>.</span>"
	fmt.Fprint(w, text)
}
