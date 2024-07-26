package client

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/israelalvesmelo/desafio-multithreading/internal/entity"
)

type viaCepClient struct {
}

func NewViaCepClient() CepInterface[entity.ViaCep] {
	return &viaCepClient{}
}

func (b *viaCepClient) GetCep(ctx context.Context, cep string) (entity.ViaCep, error) {
	var c entity.ViaCep
	url := fmt.Sprintf("http://viacep.com.br/ws/%s/json", cep)
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
