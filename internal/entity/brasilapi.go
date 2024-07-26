package entity

import "fmt"

type BrasilApiCep struct {
	Cep          string `json:"cep"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street       string `json:"street"`
	Service      string `json:"service"`
}

func (b BrasilApiCep) ToString() string {
	return fmt.Sprintf("Cep: %s, State: %s, City: %s, Neighborhood: %s, Street: %s, Service: %s",
		b.Cep, b.State, b.City, b.Neighborhood, b.Street, b.Service)
}
