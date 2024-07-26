package entity

import "fmt"

type ViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func (v ViaCep) ToString() string {
	return fmt.Sprintf("Cep: %s, Logradouro: %s, Complemento: %s, Unidade: %s, Bairro: %s, Localidade: %s, Uf: %s, Ibge: %s, Gia: %s, Ddd: %s, Siafi: %s",
		v.Cep, v.Logradouro, v.Complemento, v.Unidade, v.Bairro, v.Localidade, v.Uf, v.Ibge, v.Gia, v.Ddd, v.Siafi)
}
