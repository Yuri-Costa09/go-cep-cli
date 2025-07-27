package api

import (
	"encoding/json" // LIBRARY FOR JSON ENCODING AND DECODING
	"log"
	"net/http" // LIBRARY FOR HTTP REQUESTS

	"github.com/Yuri-Costa09/go-cep-cli/shared"
	"github.com/Yuri-Costa09/go-cep-cli/utils"
)

func FetchEndereco(cep string) shared.Endereco {
	cepFormatted := utils.FormatCep(cep) // FORMAT CEP TO BE COMPATIBLE WITH THE API

	res, err := http.Get("https://viacep.com.br/ws/" + cepFormatted + "/json/") // MAKE HTTP REQUEST TO THE API
	if err != nil {
		log.Fatal("Erro ao buscar CEP:", err)
	}
	defer res.Body.Close()

	var endereco shared.Endereco

	// decodes the response body into the endereco struct
	// by taking the memory address of the endereco, and changing
	// the original value
	err = json.NewDecoder(res.Body).Decode(&endereco)

	if err != nil {
		log.Fatal("Erro ao decodificar JSON:", err)
	}

	return endereco
}
