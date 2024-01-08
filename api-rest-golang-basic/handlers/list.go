package handlers

import (
	"github.com/patrik-rangel/API-Rest-Go_CRUD/tree/main/api-rest-golang-basic/models"
	"log"
	"net/http"
	"encoding/json"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}