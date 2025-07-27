package utils

import "strings"

func FormatCep(cep string) string {
	return strings.ReplaceAll(cep, "-", "")
}
