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

		p, err := config.RandomPokemon(shinyOdds)
		if err != nil {
			return err
		}

		fmt.Println(p.String(colorOverrides))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
	randomCmd.Flags().Float32("shiny-odds", 1/128, "shiny probablity between 0.0 and 1.0")
	randomCmd.Flags().StringArray("override-rgba", []string{}, "override a given rgba color. example \"120 90 23 255=99 18 44 255\"")
}
