package tests

import (
    "fmt"
		"io"
		"net/http"
    "net/http/httptest"
    "strings"
    "testing"
)

var (
    server   *httptest.Server
    reader   io.Reader //Ignore this for now
    usersUrl string
)

func init() {
    server = httptest.NewServer(Handlers())

    usersUrl = fmt.Sprintf("%s/invoices", server.URL)
}

func TestCreateInvoiceEndpoint(t *testing.T) {
	initDb()
	defer db.Close()

	invoiceJson := `{"document" : "12312311999"}`
	reader := strings.NewReader(invoiceJson)

	req, err := http.NewRequest("POST", usersUrl, reader)
	req.Header.Set("Token", "token#app1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusCreated {
		t.Errorf("Erro! Reposta não esperada: %d", res.StatusCode)
	}
}

func TestGetAllInvoicesEndpoint(t *testing.T) {
	initDb()
	defer db.Close()

	reader = strings.NewReader("")

	req, err := http.NewRequest("GET", usersUrl, reader)
	req.Header.Set("Token", "token#app1")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Erro! Reposta não esperada: %d", res.StatusCode)
	}
}
