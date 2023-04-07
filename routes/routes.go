package routes

import (
	"Go-API-Rest/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest(pPorta string) {
	gorilla := mux.NewRouter()
	gorilla.HandleFunc("/", controllers.Home)
	gorilla.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("GET")
	gorilla.HandleFunc("/api/personalidades/{id}", controllers.GetPersonalitie).Methods("GET")
	log.Fatalln(http.ListenAndServe(pPorta, gorilla))
}
