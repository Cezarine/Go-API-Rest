package controllers

import (
	"Go-API-Rest/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Home Page")
}

func AllPersonalities(res http.ResponseWriter, req *http.Request) {
	json.NewEncoder(res).Encode(models.Personalidades)
}

func GetPersonalitie(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	vId := vars["id"]

	for _, personalidade := range models.Personalidades {
		if strconv.Itoa(personalidade.Id) == vId {
			json.NewEncoder(res).Encode(personalidade)
		}
	}
}
