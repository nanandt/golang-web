package golang_web

import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// bikin server
	server := http.Server{
		Addr: "localhost:8080",
	}

	// menjalankan server
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
