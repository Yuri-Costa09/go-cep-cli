package cmd

import (
	"github.com/Yuri-Costa09/go-cep-cli/api"
	"github.com/Yuri-Costa09/go-cep-cli/utils"
	"github.com/spf13/cobra"
)

// WE CREATE A COMMAND BY USING &cobra.Command{}
var FindInfoCmd = &cobra.Command{
	Use:     "findinfo",                            // COMMAND NAME
	Aliases: []string{"findinfo, findinformation"}, // EXTRA NAMES THAT CAN BE USED TO CALL THE COMMAND
	Short:   "Finds information about a CEP",       // SHORT DESCRIPTION OF THE COMMAND
	Args:    cobra.ExactArgs(1),                    // HOW ARGUMENTS THAT THE COMMAND EXPECTS
	Run: func(cmd *cobra.Command, args []string) {
		cep := args[0]                     // GET THE CEP FROM THE ARGUMENTS
		endereco := api.FetchEndereco(cep) // FETCH THE ENDERECO FROM THE API
		utils.PrintEndereco(endereco)      // PRINT THE ENDERECO
	},
}

// THIS IS A FUNCTION THAT ADDS THE COMMAND TO THE ROOT COMMAND
// ? WE MUST ADD ALL COMMANDS THAT WE CREATE TO THE ROOT COMMAND, SO THEY CAN BE ACCESSIBLE
func init() {
	rootCmd.AddCommand(FindInfoCmd)
}
