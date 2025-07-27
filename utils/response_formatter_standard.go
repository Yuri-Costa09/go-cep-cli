package utils

import (
	"fmt"

	"github.com/Yuri-Costa09/go-cep-cli/shared"
)

func PrintEndereco(e shared.Endereco) {
	fmt.Println("🔎 Endereço encontrado:")
	fmt.Println("📍 Logradouro :", e.Logradouro)
	fmt.Println("🏘️ Bairro     :", e.Bairro)
	fmt.Println("🏙️ Cidade     :", e.Localidade)
	fmt.Println("🌎 Estado     :", e.UF)
	fmt.Println("📮 CEP        :", e.CEP)

	// Shows optional field if they are not empty
	if e.Complemento != "" {
		fmt.Println("🧩 Complemento:", e.Complemento)
	}
	if e.DDD != "" {
		fmt.Println("📞 DDD        :", e.DDD)
	}
}
