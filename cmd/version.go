package cmd

import (
	"fmt"

	"github.com/andersonlira/golangspell-txtdb/config"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "golangspell-txtdb-version",
	Short: "golangspell-txtdb version number",
	Long:  `Shows the golangspell-txtdb current installed version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("golangspell-txtdb v%s -- HEAD\n", config.Version)
	},
}
