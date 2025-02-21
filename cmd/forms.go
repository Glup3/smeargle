package cmd

import (
	"fmt"
	"strings"

	"github.com/glup3/smeargle/pokemon"
	"github.com/spf13/cobra"
)

var formsCmd = &cobra.Command{
	Use:   "forms",
	Short: "print all forms for a Pokemon separated by line breaks",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		fmt.Println(strings.Join(config.GetForms(args[0]), "\n"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(formsCmd)
}
