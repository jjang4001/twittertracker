package main

import (
	"log"
	"net/http"
	"os"
	"twittertracker/api/handlers"
	"twittertracker/api/middleware"
	"twittertracker/common"
	"twittertracker/datastore"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConnectionString := os.Getenv(common.LocalRedis)
	port := os.Getenv(common.Port)

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
	r.Handle("/example/{exampleId}", handlers.ExampleHandler(&env)).Methods(common.Get, common.Post)

	// middleware

	http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	http.ListenAndServe(port, nil)
}
