package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest) // Response Code : 400
		fmt.Fprint(w, "Name is empty")
	} else {
		// Default response code OK is 200
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestResponseCodeInvalid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost, nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, localhost+"?name=Taufiq", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
}
