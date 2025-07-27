package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// THIS IS THE ROOT COMMAND
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

// THIS IS THE FUNCTION THAT EXECUTES THE ROOT COMMAND
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An error while executing ceptools '%s'\n", err)
		os.Exit(1)
	}
}
