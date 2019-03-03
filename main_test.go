package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRouting(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		path     string
		param    string
		body     string
		respCode int
	}{
		{name: "get", method: "GET", path: "/hello", param: "", respCode: 200, body: "Hello stranger!."},
		{name: "getWithParam", method: "GET", path: "/hello", param: "?name=nikaein", respCode: 200, body: "Hello nikaein!."},
		{name: "post", method: "POST", path: "/hello", param: "", respCode: 501, body: "Method Not Implemented"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			url := tt.path + tt.param

			req, err := http.NewRequest(tt.method, url, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(greeting)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.respCode {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			if rr.Body.String() != tt.body {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.body)
			}

		})
	}
}
