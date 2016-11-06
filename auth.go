package main

import (
	// "encoding/json"
	// "github.com/gorilla/mux"
	// "io/ioutil"
	"net/http"
	// "strconv"
  // "fmt"
  "strings"
)

var validTokens = []string{"token#app1", "token#app2"}

func Authenticate(w http.ResponseWriter, r *http.Request) int {
  var status int
  header := r.Header
  token := strings.Join(header["Token"], "")
  if isValueInList(token, validTokens){
    status = http.StatusOK
  } else {
    status = http.StatusUnauthorized
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
