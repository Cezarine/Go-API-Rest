package routes

import (
	"Go-API-Rest/controllers"
	"Go-API-Rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest(pPorta string) {
	gorilla := mux.NewRouter()
	gorilla.Use(middleware.ContentTypeMiddleware)
	gorilla.HandleFunc("/", controllers.Home)
	gorilla.HandleFunc("/api/personalidades", controllers.AllPersonalities).Methods("GET")
	gorilla.HandleFunc("/api/personalidade/{id}", controllers.GetPersonality).Methods("GET")
	gorilla.HandleFunc("/api/personalidade", controllers.InsertNewPersonality).Methods("POST")
	gorilla.HandleFunc("/api/personalidade/{id}", controllers.DeletePersonality).Methods("DELETE") //METHOD delete, porém execute PUT
	gorilla.HandleFunc("/api/personalidade/{id}", controllers.UpdatePersonality).Methods("PUT")
	log.Fatalln(http.ListenAndServe(pPorta, handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}))(gorilla)))
}
