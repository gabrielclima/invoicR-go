package auth

import (
	"encoding/json"
	"github.com/gabrielclima/go_rest_api/utils"
	"net/http"
	"strings"
)

var validTokens = []string{"token#app1", "token#app2"}

// Authenticate requests based in validTokens list
func Authenticate(w http.ResponseWriter, r *http.Request) int {
	var status int
	var res []byte
	var err error
	header := r.Header
	token := strings.Join(header["Authorization"], "")

	if isValueInList(token, validTokens) {
		status = http.StatusOK
	} else {
		status = http.StatusUnauthorized
		w.WriteHeader(status)
		res, err = json.Marshal(utils.JsonErr{Code: http.StatusUnauthorized, Message: "Unauthorized"})
		utils.CheckErr(err)
		w.Write(res)
	}

	return status
}

func isValueInList(value string, list []string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}
