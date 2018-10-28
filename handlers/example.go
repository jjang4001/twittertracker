package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"twittertracker/datastore"
	"twittertracker/models"

	"github.com/gorilla/mux"
)

func DisplayExample(w http.ResponseWriter, r *http.Request, e *datastore.Env, exampleId string) {

	ex, err := e.DB.GetExample(exampleId)
	if err != nil {
		log.Print(err)
	}
	example := models.NewExample(ex.ExampleId, "my example value")

	json.NewEncoder(w).Encode(example)
}

func AddExample(w http.ResponseWriter, r *http.Request, e *datastore.Env, exampleId string) {

	ex := models.NewExample(exampleId, "my example value")

	err := e.DB.CreateExample(ex)
	if err != nil {
		log.Print(err)
	}

	example, err := e.DB.GetExample(exampleId)
	if err != nil {
		log.Print(err)
	} else {
		fmt.Printf("Fetch example result: %+v\n", example)
	}
}

func DisplayDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Println("default example endpoint here")
}

func ExampleHandler(e *datastore.Env) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		v := mux.Vars(r)

		switch r.Method {
		case "GET":
			DisplayExample(w, r, e, v["exampleId"])
		case "POST":
			AddExample(w, r, e, v["exampleId"])
		default:
			DisplayDefault(w, r)
		}
	})
}
