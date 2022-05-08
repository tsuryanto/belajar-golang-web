package belajar_golang_web

import (
	"embed"
	_ "embed"
	"io/fs"
	"net/http"
	"testing"
)

/*
	Dengan FileServer, kita tidak perlu me-load file degan cara manual.
	Berguna untuk menggunakan static file server (html, js , css, images, dll)
*/

func TestFileServer(t *testing.T) {
	dir := http.Dir("./resources")
	fileServer := http.FileServer(dir)

	mux := http.NewServeMux()

	// strip prefix untuk menghapus prefix (disini '/static/') yang kita tambahkan,
	// agar tidak terjadi 404 Not Found
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources
var resources embed.FS

func TestFileServerGolangEmbed(t *testing.T) {
	// digunakan agar prefix resources nya tidak terinclude di url
	directory, _ := fs.Sub(resources, "resources")

	// file server nya ambil dari embed
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
