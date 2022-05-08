package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request) {
	first_name := r.PostFormValue("first_name")
	last_name := r.PostFormValue("last_name")

	fmt.Fprintf(w, "Halo %s %s", first_name, last_name)
}

/*
	Bentuk Form Post sama dengan Query Parameter, namun bedanya
	form_post ada di Body, sementara Query Parameter di URL

	untuk Test menggunakan Form Post wajib memasukkan "Content-Type" nya "application/x-www-form-urlencoded"
	untuk memberitahu bahwa ini form POST
*/

func TestPostForm(t *testing.T) {
	requestBody := strings.NewReader("first_name=Taufiq&last_name=Suryanto")
	request := httptest.NewRequest(http.MethodPost, localhost, requestBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
