package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Yuri-Costa09/go-cep-cli/shared"
	"github.com/Yuri-Costa09/go-cep-cli/utils"
)

func FetchEndereco(cep string) shared.Endereco {
	cepFormatted := utils.FormatCep(cep)

	res, err := http.Get("https://viacep.com.br/ws/" + cepFormatted + "/json/")

	if err != nil {
		log.Fatal("Erro ao buscar CEP:", err)
	}
	defer res.Body.Close()

	var endereco shared.Endereco

	err = json.NewDecoder(res.Body).Decode(&endereco)

	if err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	return endereco
}
