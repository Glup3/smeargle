package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

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

		generations, err := parseGenerationString(generationsString)
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

func parseGenerationString(input string) ([]int, error) {
	if input == "" {
		return []int{}, nil
	}

	var result []int
	seen := pokemon.NewSet[int]()
	parts := strings.Split(input, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			// Handle range
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("invalid range: %s", part)
			}

			start, err := strconv.Atoi(rangeParts[0])
			if err != nil {
				return nil, fmt.Errorf("invalid start number: %s", rangeParts[0])
			}

			end, err := strconv.Atoi(rangeParts[1])
			if err != nil {
				return nil, fmt.Errorf("invalid end number: %s", rangeParts[1])
			}

			if start > end {
				return nil, fmt.Errorf("start of range is greater than end: %s", part)
			}

			for i := start; i <= end; i++ {
				if !seen.Has(i) {
					result = append(result, i)
					seen.Add(i)
				}
			}
		} else {
			// Handle individual number
			num, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid number: %s", part)
			}
			if !seen.Has(num) {
				result = append(result, num)
				seen.Add(num)
			}
		}
	}

	return result, nil
}
