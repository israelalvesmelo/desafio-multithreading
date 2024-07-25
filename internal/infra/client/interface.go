package client

import "context"

type CepInterface[T interface{}] interface {
	GetCep(ctx context.Context, cep string) (T, error)
}
