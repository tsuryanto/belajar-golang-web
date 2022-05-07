package belajar_golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")

	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func SayHelloFullName(w http.ResponseWriter, r *http.Request) {
	first_name := r.URL.Query().Get("first_name")
	last_name := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Hello %s %s", first_name, last_name)
}

type TestFormat struct {
	Name     string
	Request  interface{}
	Expected string
}

type Name struct {
	FirstName string
	LastName  string
}

func TestQueryParameter(t *testing.T) {

	helloTest := []TestFormat{
		{
			Name:     "SayHello(Taufiq)",
			Request:  "Taufiq",
			Expected: "Hello Taufiq",
		},
		{
			Name:     "SayHello(Sri)",
			Request:  "Sri",
			Expected: "Hello Sri",
		},
	}

	for _, test := range helloTest {
		t.Run(test.Name, func(t *testing.T) {

			request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name="+test.Request.(string), nil)
			recorder := httptest.NewRecorder()

			SayHello(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)

			fmt.Println(string(body))
			assert.Equal(t, test.Expected, string(body))
		})
	}
}

func TestMultipleQueryParameter(t *testing.T) {

	helloTest := []TestFormat{
		{
			Name:     "SayHello(Taufiq)",
			Request:  Name{FirstName: "Taufiq", LastName: "Suryanto"},
			Expected: "Hello Taufiq Suryanto",
		},
		{
			Name:     "SayHello(Sri)",
			Request:  Name{FirstName: "Sri", LastName: "Nuryanti"},
			Expected: "Hello Sri Nuryanti",
		},
	}

	for _, test := range helloTest {
		t.Run(test.Name, func(t *testing.T) {

			nameObj := test.Request.(Name)

			request := httptest.NewRequest(
				http.MethodGet,
				"http://localhost:8080/hello?first_name="+nameObj.FirstName+"&last_name="+nameObj.LastName, nil)
			recorder := httptest.NewRecorder()

			SayHelloFullName(recorder, request)

			response := recorder.Result()
			body, _ := io.ReadAll(response.Body)

			fmt.Println(string(body))
			assert.Equal(t, test.Expected, string(body))
		})
	}
}

/* Mengambil multiple query parameter (name) (memiliki nama yang sama) */
func MultipleQueryParamValues(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	names := query["name"]
	fmt.Fprint(w, strings.Join(names, " "))
}

func TestMultipleQueryParamValues(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello?name=Taufiq&name=Suryanto&name=S.Kom", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParamValues(recorder, request)
	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

	assert.Equal(t, "Taufiq Suryanto S.Kom", string(body))
}
