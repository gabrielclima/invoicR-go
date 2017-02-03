package auth

import (
	"encoding/json"
	"net/http"
  "strings"
)

var validTokens = []string{"token#app1", "token#app2"}

func Authenticate(w http.ResponseWriter, r *http.Request) int {
	var status int
	var res []byte
	var err error
  header := r.Header
	token := strings.Join(header["Authorization"], "")

  if isValueInList(token, validTokens){
    status = http.StatusOK
  } else {
    status = http.StatusUnauthorized
		w.WriteHeader(status)
		res, err = json.Marshal(jsonErr{Code: http.StatusUnauthorized, Text: "Unauthorized"})
		checkErr(err)
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
