package cmd

import (
	"fmt"

	"github.com/glup3/smeargle/pokemon"
	"github.com/spf13/cobra"
)

var histogramCmd = &cobra.Command{
	Use:   "histogram",
	Short: "prints color histogram for given Pokemon",
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

		for rgba, count := range p.ColorHistogram() {
			fmt.Println(rgba, ":", count)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(histogramCmd)
	histogramCmd.Flags().StringP("name", "n", "", "pokemon name as slug")
	histogramCmd.Flags().StringP("form", "f", "", "show alternate form")
	histogramCmd.Flags().BoolP("shiny", "s", false, "show shiny version")
	histogramCmd.MarkFlagRequired("name")
}
