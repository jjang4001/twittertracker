package main

import (
	"log"
	"net/http"
	"os"
	"twittertracker/datastore"
	"twittertracker/handlers"
	"twittertracker/middleware"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	WEBSERVERPORT = ":3000"
)

func main() {
	dbConnectionString := "127.0.0.1:6379"
	db, err := datastore.NewDatastore(datastore.REDIS, dbConnectionString)

	if err != nil {
		log.Print(err)
	}

	defer db.Close()
	env := datastore.Env{DB: db}

	r := mux.NewRouter()

	// handlers

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/socket", handlers.SocketHandler)
	r.Handle("/example/{exampleId}", handlers.ExampleHandler(&env)).Methods("GET", "POST")

	// middleware

	http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	http.ListenAndServe(WEBSERVERPORT, nil)
}
