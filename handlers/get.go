package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/bertbr/poc-golang/models"
	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Printf("Error converting id to int: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	todo, err := models.Get(id)
	if err != nil {
		log.Printf("Erro ao buscar registro: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
