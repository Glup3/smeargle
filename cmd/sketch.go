package cmd

import (
	"fmt"

	"github.com/glup3/smeargle/pokemon"
	"github.com/spf13/cobra"
)

var sketchCmd = &cobra.Command{
	Use:   "sketch",
	Short: "paints the Pokemon",
	RunE: func(cmd *cobra.Command, args []string) error {
		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		name := "pikachu"
		form := ""

		im, err := config.GetImage(name, form)
		if err != nil {
			return err
		}

		p := pokemon.NewPokemon(name, im)
		fmt.Println(p.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sketchCmd)
}
