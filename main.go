package main

import (
	"fmt"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

var (
	// VERSION keeps the version of application
	VERSION = 0.1
)

func main() {
	// TODO: for keep it simple i dont use router !
	http.HandleFunc("/hello", greeting)

	// get port from
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}

	// init http server
	log.Println("Starting Server On Port:", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, nil))
}

// greeting get name and say hello to it if not name provided counsider you as a stranger! .
func greeting(w http.ResponseWriter, r *http.Request) {
	// TODO: use router to avoid check method in each function.
	if r.Method == "GET" {
		name := r.FormValue("name")
		// TODO: shuld i use regex for name whats the name condition?
		w.WriteHeader(http.StatusOK)
		if name == "" {
			fmt.Fprint(w, "Hello stranger!.")
		} else {
			fmt.Fprint(w, "Hello ", name, "!.")
		}

	} else {
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "Method Not Implemented")
	}
}
