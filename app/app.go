package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// App is Application structure
type App struct {
	Version string
	Port    string
	Router  *mux.Router
	DB      *gorm.DB
}

// InitializeDB init db
func (a *App) InitializeDB() {

	// dbURI from os env
	var dbURI string
	if dbURI = os.Getenv("dbURI"); dbURI == "" {
		log.Fatal("could not find dbURI")
	}

	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatalf("Could not connect database err:%v", err)
	}

	a.DB = DBMigrate(db)
}

// Run application
func (a *App) Run() {

	// port from os env
	if a.Port = os.Getenv("PORT"); a.Port == "" {
		a.Port = "8080"
	}
	// init router
	a.RouterInit()

	// init http server
	log.Printf("Server version:%v Start On Port:%v", a.Version, a.Port)
	log.Fatal(http.ListenAndServe(":"+a.Port, a.Router))
}

// RouterInit initalize all route
func (a *App) RouterInit() {
	a.Router = mux.NewRouter()
	a.Router.HandleFunc("/hello", logRequest(a.greeting)).Methods("GET")
}

func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("RequestURI: %s RemoteAddr: %v", r.RequestURI, r.RemoteAddr)
		next.ServeHTTP(w, r)
	}
}

// greeting get name and say hello to it if not name provided counsider you as a stranger! .
func (a *App) greeting(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	// TODO: shuld i use regex for name whats the name condition?
	w.WriteHeader(http.StatusOK)
	if name == "" {
		fmt.Fprint(w, "Hello stranger!.")
	} else {
		// inset name to db
		InsertName(a.DB, &Name{Name: name})
		fmt.Fprint(w, "Hello ", name, "!.")
	}

}
