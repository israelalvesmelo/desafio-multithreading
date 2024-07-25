package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/israelalvesmelo/desafio-multithreading/internal/entity"
	"github.com/israelalvesmelo/desafio-multithreading/internal/infra/client"
)

type cepHandler struct {
	brasilApiClient client.CepInterface[entity.BrasilApiCep]
}

func NewCepHandler() *cepHandler {
	return &cepHandler{
		brasilApiClient: client.NewBrasilApiClient(),
	}
}

func (h *cepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if cep == "" {
		log.Print("Cep não pode ser vazio")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	resp, err := h.brasilApiClient.GetCep(ctx, cep)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)

}
