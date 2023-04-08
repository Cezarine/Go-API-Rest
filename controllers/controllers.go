package controllers

import (
	"Go-API-Rest/database"
	"Go-API-Rest/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const personalitysActives = "ativo = true"
const personalitysInactives = "ativo = false"

func Home(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Home Page")
}

func AllPersonalities(res http.ResponseWriter, req *http.Request) {
	var personalidades []models.Personalidade
	database.DB.Where(personalitysActives).Find(&personalidades)
	json.NewEncoder(res).Encode(personalidades)
}

func GetPersonality(res http.ResponseWriter, req *http.Request) {
	vVarRequest := mux.Vars(req)
	vId := vVarRequest["id"]
	var personalidade models.Personalidade
	database.DB.Where(personalitysActives).First(&personalidade, vId)
	json.NewEncoder(res).Encode(personalidade)
}

func InsertNewPersonality(res http.ResponseWriter, req *http.Request) {
	var personality models.Personalidade
	json.NewDecoder(req.Body).Decode(&personality)
	database.DB.Create(&personality)
	json.NewEncoder(res).Encode(personality)
}

func DeletePersonality(res http.ResponseWriter, req *http.Request) {
	vVarRequest := mux.Vars(req)
	vId := vVarRequest["id"]
	var personality models.Personalidade
	database.DB.First(&personality, vId)
	database.DB.Model(&personality).Update("ativo", false)
	json.NewEncoder(res).Encode(personality)
}

func UpdatePersonality(res http.ResponseWriter, req *http.Request) {
	vVarRequest := mux.Vars(req)
	vId := vVarRequest["id"]
	var personality models.Personalidade
	database.DB.First(&personality, vId)
	json.NewDecoder(req.Body).Decode(&personality)
	database.DB.Save(&personality)
	json.NewEncoder(res).Encode(personality)
}
