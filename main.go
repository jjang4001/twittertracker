package main

import (
	"net/http"
	"os"
	"twittertracker/handlers"
	"twittertracker/middleware"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

const (
	WEBSERVERPORT = ":8080"
)

func main() {
	r := mux.NewRouter()

	// handlers

	r.HandleFunc("/", handlers.HomeHandler)
	r.HandleFunc("/socket", handlers.SocketHandler)

	// middleware

	http.Handle("/", middleware.ContextExampleHandler(middleware.PanicRecoveryHandler(ghandlers.LoggingHandler(os.Stdout, r))))

	http.ListenAndServe(WEBSERVERPORT, nil)
}
