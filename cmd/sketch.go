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
		slug, err := cmd.Flags().GetString("name")
		if err != nil {
			return err
		}

		shiny, err := cmd.Flags().GetBool("shiny")
		if err != nil {
			return err
		}

		form, err := cmd.Flags().GetString("form")
		if err != nil {
			return err
		}

		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		im, err := config.FindImage(slug, form, shiny)
		if err != nil {
			return err
		}

		p := pokemon.NewPokemon(slug, im)
		fmt.Println(p.String())

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sketchCmd)
	sketchCmd.Flags().StringP("name", "n", "", "pokemon name as slug")
	sketchCmd.Flags().StringP("form", "f", "", "show alternate form")
	sketchCmd.Flags().BoolP("shiny", "s", false, "show shiny version")
	sketchCmd.MarkFlagRequired("name")
}
