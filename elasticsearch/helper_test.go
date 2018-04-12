package elasticsearch

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

var wd, _ = os.Getwd()
var testdata = wd + "/testdata"

func CannedResponder(s string) *Connection {
	return mockConnection(CannedResponse{id: s})
}

type CannedResponse struct {
	id string
}

func (cr CannedResponse) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	respondWith(w, cr.id)
}

func mockConnection(handler http.Handler) *Connection {
	server := httptest.NewServer(handler)

	conn, err := Connect(server.URL)
	if err != nil {
		panic(err)
	}
	return conn
}

func respondWith(w io.Writer, response string) {
	f, err := os.Open(testdata + "/response/" + response)
	defer f.Close()
	if err != nil {
		panic(err)
	}

	io.Copy(w, f)
}
