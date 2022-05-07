package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

// Writer di gunakan untuk menulis data kembalian
// Menggunakan Fprint untuk menulis di halaman browser
// Dengan membuka localhost:8080 di browser
func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request) {
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// http.Request berasal dari client
// Digunakan untuk mengetahui data request , misal HTTP method, URI, http header, dll

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
