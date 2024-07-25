package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/israelalvesmelo/desafio-multithreading/internal/entity"
)

type brasilApiClient struct {
}

func NewBrasilApiClient() CepInterface[entity.BrasilApiCep] {
	return &brasilApiClient{}
}

func (b *brasilApiClient) GetCep(ctx context.Context, cep string) (entity.BrasilApiCep, error) {
	var c entity.BrasilApiCep
	url := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return c, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return c, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c, err
	}

	err = json.Unmarshal(body, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}
