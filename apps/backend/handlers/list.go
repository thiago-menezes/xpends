package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"xpends/webapi/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	spends, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros %v", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(spends)
}
