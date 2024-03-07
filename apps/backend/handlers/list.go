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
	json.NewEncoder(w).Encode(spends)
}
