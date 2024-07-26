package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/israelalvesmelo/desafio-multithreading/internal/entity"
	"github.com/israelalvesmelo/desafio-multithreading/internal/infra/client"
)

type cepHandler struct {
	brasilApiClient client.CepInterface[entity.BrasilApiCep]
	viacepClient    client.CepInterface[entity.ViaCep]
}

func NewCepHandler(brasilApiClient client.CepInterface[entity.BrasilApiCep],
	viacepClient client.CepInterface[entity.ViaCep]) *cepHandler {
	return &cepHandler{
		brasilApiClient: brasilApiClient,
		viacepClient:    viacepClient,
	}
}

func (h *cepHandler) GetCep(w http.ResponseWriter, r *http.Request) {
	cep := chi.URLParam(r, "cep")
	if !isCepValid(cep) {
		fmt.Printf("Invalid CEP %s", cep)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	c1 := make(chan entity.BrasilApiCep)
	c2 := make(chan entity.ViaCep)

	go func() {
		resp, _ := h.brasilApiClient.GetCep(ctx, cep)
		c1 <- resp
	}()
	go func() {
		resp, _ := h.viacepClient.GetCep(ctx, cep)
		c2 <- resp
	}()

	select {
	case resp := <-c1:
		fmt.Printf("Received from BrasilAPI [%s] \n", resp.ToString())
		w.WriteHeader(http.StatusOK)

	case resp := <-c2:
		fmt.Printf("Received from ViaCep [%s] \n", resp.ToString())
		w.WriteHeader(http.StatusOK)

	case <-ctx.Done():
		fmt.Println("Timeout exceeded")
		w.WriteHeader(http.StatusNotFound)
	}
}

func isCepValid(value string) bool {
	_, err := strconv.Atoi(value)
	return err == nil
}
