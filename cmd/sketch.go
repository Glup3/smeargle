package cmd

import (
	"fmt"

	"github.com/glup3/smeargle/pokemon"
	"github.com/spf13/cobra"
)

var (
	shiny bool
	form  string
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

		im, err := config.GetImage(name, form, shiny)
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
	sketchCmd.Flags().BoolVarP(&shiny, "shiny", "s", false, "show shiny version")
	sketchCmd.Flags().StringVarP(&form, "form", "f", "", "show alternate form")
}
