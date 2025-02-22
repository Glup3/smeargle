package cmd

import (
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

		orderByString, err := cmd.Flags().GetString("order-by")
		if err != nil {
			return err
		}

		orderBy, err := pokemon.ParseOrderByString(orderByString)
		if err != nil {
			return err
		}

		sortDirectionString, err := cmd.Flags().GetString("sort-direction")
		if err != nil {
			return err
		}

		sortDirection, err := pokemon.ParseSortDirectionString(sortDirectionString)
		if err != nil {
			return err
		}

		gens, err := pokemon.ParseGenerationString(generationsString)
		if err != nil {
			return err
		}

		config, err := pokemon.NewPokemonConfig()
		if err != nil {
			return err
		}

		slugs, err := config.GetSlugs(gens, orderBy, sortDirection)
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
	namesCmd.Flags().String("order-by", "alphabet", "order names by (alphabet|idx), default is alphabet")
	namesCmd.Flags().String("sort-direction", "asc", "sort direction (asc|desc), default is asc")
}
