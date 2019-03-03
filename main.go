package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

var (
	// VERSION keeps the version of application
	VERSION string
	// PORT contain the expose port
	PORT string
)

func main() {

	//init port
	PortEnv()

	// init http server
	log.Printf("Server version:%v Start On Port:%v", VERSION, PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, Router()))
}

// PortEnv return port from os env
func PortEnv() string {
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}
	return PORT
}

// Router return gorilla mux router
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/hello", greeting).Methods("GET")
	return r
}

// greeting get name and say hello to it if not name provided counsider you as a stranger! .
func greeting(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	// TODO: shuld i use regex for name whats the name condition?
	w.WriteHeader(http.StatusOK)
	if name == "" {
		fmt.Fprint(w, "Hello stranger!.")
	} else {
		fmt.Fprint(w, "Hello ", name, "!.")
	}

}
