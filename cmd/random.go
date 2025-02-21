package cmd

import (
	"fmt"

	"github.com/glup3/smeargle/pokemon"
	"github.com/spf13/cobra"
)

var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "paints a random pokemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		p, err := config.RandomPokemon()
		if err != nil {
			return err
		}

		fmt.Println(p.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}
