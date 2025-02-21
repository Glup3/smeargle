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
		shinyOdds, err := cmd.Flags().GetFloat32("shiny-odds")
		if err != nil {
			return err
		}

		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		p, err := config.RandomPokemon(shinyOdds)
		if err != nil {
			return err
		}

		fmt.Println(p.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().Float32("shiny-odds", 1/128, "shiny probablity between 0.0 and 1.0")
}
