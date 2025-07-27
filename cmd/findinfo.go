package cmd

import (
	"fmt"

	"github.com/Yuri-Costa09/go-cep-cli/api"
	"github.com/spf13/cobra"
)

var FindInfoCmd = &cobra.Command{
	Use:     "findinfo",
	Aliases: []string{"findinfo, findinformation"},
	Short:   "Finds information about a CEP",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		cep := args[0]
		endereco := api.FetchEndereco(cep)
		fmt.Println(endereco)
	},
}

func init() {
	rootCmd.AddCommand(FindInfoCmd)
}
