package helper

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
)

// Buat versi struct
type HttpTest struct {
	Param  string
	Method string
}

func (test *HttpTest) Request() *http.Request {
	return httptest.NewRequest(test.Method, test.Param, nil)
}

func (test *HttpTest) Recorder() *httptest.ResponseRecorder {
	return httptest.NewRecorder()
}

func (test *HttpTest) ResultBody(recorder *httptest.ResponseRecorder) []byte {
	body, _ := io.ReadAll(recorder.Result().Body)
	return body
}

func RunHttpTest(handler func(w http.ResponseWriter, r *http.Request), param string, method string) {
	ht := HttpTest{
		Param:  param,
		Method: method,
	}
	recorder := ht.Recorder()
	handler(recorder, ht.Request())

	body := ht.ResultBody(recorder)
	fmt.Println(string(body))
}
