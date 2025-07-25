package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func hello() {
	fmt.Printf("Hello!\n\n")
}

var AddCmd = &cobra.Command{
	Use:     "hello",
	Aliases: []string{"sayhello, falaroi"},
	Short:   "Says Hello",
	Long:    "Says hello -long",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		hello()
	},
}

func init() {
	rootCmd.AddCommand(AddCmd)
}
