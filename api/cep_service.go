package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Yuri-Costa09/go-cep-cli/utils"
)

type Endereco struct {
	CEP         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Unidade     string `json:"unidade"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	UF          string `json:"uf"`
	Estado      string `json:"estado"`
	Regiao      string `json:"regiao"`
	IBGE        string `json:"ibge"`
	GIA         string `json:"gia"`
	DDD         string `json:"ddd"`
	SIAFI       string `json:"siafi"`
}

func FetchEndereco(cep string) Endereco {
	cepFormatted := utils.FormatCep(cep)

	res, err := http.Get("https://viacep.com.br/ws/" + cepFormatted + "/json/")

	if err != nil {
		log.Fatal("Erro ao buscar CEP:", err)
	}
	defer res.Body.Close()

	var endereco Endereco

	err = json.NewDecoder(res.Body).Decode(&endereco)

	if err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	return endereco
}
