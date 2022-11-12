package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/bertbr/poc-golang/models"
	"github.com/go-chi/chi/v5"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		log.Printf("Error converting id to int: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := models.Delete(id)
	if err != nil {
		log.Printf("Erro ao deletar registro: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram deletados %d registros", rows)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": fmt.Sprintf("Registro removido com sucesso!: %d", id),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
