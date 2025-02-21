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
		slugs, err := pokemon.LoadSlugs()
		if err != nil {
			return err
		}

		fmt.Println(strings.Join(slugs, "\n"))
		return nil
	},
}

func init() {
	rootCmd.AddCommand(namesCmd)
}
