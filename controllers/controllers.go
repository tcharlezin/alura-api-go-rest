package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest-api/database"
	"go-rest-api/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page!")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func FindPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func CreatePersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	json.NewDecoder(r.Body).Decode(&personalidade)
	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletePersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func UpdatePersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade

	vars := mux.Vars(r)
	id := vars["id"]

	database.DB.First(&personalidade, id)
	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Save(&personalidade)
	json.NewEncoder(w).Encode(personalidade)
}
