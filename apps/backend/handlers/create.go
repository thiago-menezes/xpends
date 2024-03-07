package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"xpends/webapi/models"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func Create(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var spend models.Spend
	err := json.NewDecoder(r.Body).Decode(&spend)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(spend)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		w.WriteHeader(http.StatusOK)
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Gasto inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)

}
