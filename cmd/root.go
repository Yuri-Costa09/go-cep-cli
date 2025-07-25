package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "ceptools",
	Short: "Uma ferramenta CLI para busca de CEP",
	Long:  `Uma ferramenta CLI para busca de CEP, fornecendo respostas estruturadas, e respostas em JSON, construido por Yuri Costa.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
	Bem vindo ao CEPTOOLS!
	------------------------------------------------------------------------------------------------------

	Comandos:
		$ cepcli buscar {CEP}		BUSCAR CEP |  formatos permitidos: "12345-678" ou "12345678"

	Flags:
		--json				Retornar em JSON
			`)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing Zero '%s'\n", err)
		os.Exit(1)
	}
}
