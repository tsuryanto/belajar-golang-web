package belajar_golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

/*
	Mux jika di framework2 bahasa lainnya disebut Router
*/

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	// Endpoint nya wajib unik , jika tidak maka akan tereplace

	/*
		URL Pattern dalam Serve Mux sederhana, hanya perlu menambahkan string sebagai endpoint
		tanpa harus memasukkan domain web kita1

		Untuk format URL Pattern '/images/' (terdapat slash di akhir)
		Jika seperti ini , ketika kita mengakses '/images/contoh' atau '/images/lihat'
		akan diarahkan ke handler '/images/'

		Namun , jika ada yg lebih panjang , misal '/images/thumbnail/' ini akan di prioritaskan
		karena URL Pattern nya lebih panjang
	*/

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hi Taufiq")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Images")
	})

	mux.HandleFunc("/images/thumbnail/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Thumbnail")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
