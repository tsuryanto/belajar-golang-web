package belajar_golang_web

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

/*
	http.ServeFile()
	Dengan menggunakan func ini, kita bisa menentukan file mana yang ingin kita tulis ke http response
*/

func ServeFile(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	// Opsi menampilkan file
	if name != "" {
		// Parameter ServeFile hanya berisi string file name
		// sehingga tidak bisa diimplementasi golang embed
		http.ServeFile(w, r, "./resources/ok.html")
	} else {
		http.ServeFile(w, r, "./resources/notfound.html")
	}
}

//go:embed resources/ok.html
var resourceOK string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name != "" {
		fmt.Fprint(w, resourceOK)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		// Handler: http.HandlerFunc(ServeFile),
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	server.ListenAndServe()
}
