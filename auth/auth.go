package auth

import (
	"encoding/json"
	"github.com/gabrielclima/go_rest_api/utils"
	"net/http"
	"strings"
	"log"
)

var validTokens = []string{"token#app1", "token#app2"}





// Middleware que autentica todas as requisições feitas na API baseando nos tokens
// presentes em validTokens
// Retorna 401 se não estiver com um token válido
func Authenticate(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	var status int
	var res []byte
	var err error
	header := r.Header
	token := strings.Join(header["Authorization"], "")

	if isValueInList(token, validTokens) {
		next(w, r)
	} else {
		status = http.StatusUnauthorized
		w.WriteHeader(status)
		res, err = json.Marshal(utils.JsonErr{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		if err != nil {
			log.Println(err)
		}
		w.Write(res)
	}
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
