package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	PORT = os.Getenv("PORT")
	flag.Parse()
	exitCode := m.Run()

	os.Exit(exitCode)
}

func TestPort(t *testing.T) {
	exp := PortEnv()
	if PORT != exp {
		t.Errorf("handler returned unexpected Port: got %v want %v",
			PORT, exp)

	}
}

func TestRouting(t *testing.T) {
	srv := httptest.NewServer(Router())
	defer srv.Close()

	res, err := http.Get(fmt.Sprintf("%s/hello", srv.URL))
	if err != nil {
		t.Fatalf("could not send get request: %v ", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expect status ok ; got %v", res.StatusCode)
	}
}

func TestGreeting(t *testing.T) {
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
					status, tt.respCode)
			}

			if rr.Body.String() != tt.body {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), tt.body)
			}

		})
	}
}
