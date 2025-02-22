package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/glup3/smeargle/pokemon"
	"github.com/spf13/cobra"
)

var namesCmd = &cobra.Command{
	Use:   "names",
	Short: "print all pokemon names separated by line breaks",

	RunE: func(cmd *cobra.Command, args []string) error {
		generationsString, err := cmd.Flags().GetString("generations")
		if err != nil {
			return err
		}

		gens, err := pokemon.ParseGenerationString(generationsString)
		if err != nil {
			return err
		}
		for _, gen := range gens {
			if gen <= 0 || gen > 8 {
				return errors.New("generation has to be between 1 and 8")
			}
		}

		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		slugs, err := config.GetSlugs(gens)
		if err != nil {
			return err
		}

		fmt.Println(strings.Join(slugs, "\n"))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(namesCmd)
	namesCmd.Flags().StringP("generations", "g", "", "provide a list of generations separated by comma (1,3,5,6), OR a range (1-4), OR both (1-3,4,6-8)")
}
