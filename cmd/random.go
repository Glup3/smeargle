package cmd

import (
	"errors"
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

		ignoreForms, err := cmd.Flags().GetBool("no-forms")
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

		generationsString, err := cmd.Flags().GetString("generations")
		if err != nil {
			return err
		}

		generations, err := pokemon.ParseGenerationString(generationsString)
		if err != nil {
			return err
		}
		for _, gen := range generations {
			if gen <= 0 || gen > 8 {
				return errors.New("generation has to be between 1 and 8")
			}
		}

		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		p, err := config.RandomPokemon(pokemon.RandomPokemonOptions{
			ShinyOdds:   shinyOdds,
			IgnoreForms: ignoreForms,
			Generations: generations,
		})
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
	randomCmd.Flags().StringP("generations", "g", "", "provide a list of generations separated by comma (1,3,5,6), OR a range (1-4), OR both (1-3,4,6-8)")
	randomCmd.Flags().Bool("no-forms", false, "ignore alternate forms when true and always paint base form")
}
