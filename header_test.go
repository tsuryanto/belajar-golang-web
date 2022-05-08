package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

var localhost string = "http://localhost:8080/"

func RequestHeader(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	fmt.Fprint(w, contentType)
}

func TestRequestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func ResponseHeader(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-powered-by", "tsuryanto")
	fmt.Fprint(w, w.Header())
}

func TestResponseHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	request.Header.Add("content-type", "application/json")

	recorder := httptest.NewRecorder()

	ResponseHeader(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
