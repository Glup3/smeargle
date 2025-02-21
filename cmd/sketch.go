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

		rgbaOverrides, err := cmd.Flags().GetStringArray("override-rgba")
		if err != nil {
			return err
		}

		colorOverrides := make([]pokemon.RGBAOverride, len(rgbaOverrides))
		for _, override := range rgbaOverrides {
			colorOverride, err := pokemon.ParseRGBAOverride(override)
			if err != nil {
				return err
			}
			colorOverrides = append(colorOverrides, colorOverride)
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
		fmt.Println(p.String(colorOverrides))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(sketchCmd)
	sketchCmd.Flags().StringP("name", "n", "", "pokemon name as slug")
	sketchCmd.Flags().StringP("form", "f", "", "show alternate form")
	sketchCmd.Flags().BoolP("shiny", "s", false, "show shiny version")
	sketchCmd.Flags().StringArray("override-rgba", []string{}, "override a given rgba color. example \"120 90 23 255=99 18 44 255\"")
	sketchCmd.MarkFlagRequired("name")
}
