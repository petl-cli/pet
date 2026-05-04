package commands

import "github.com/spf13/cobra"

var petsCmd = &cobra.Command{
	Use:   "pets",
	Short: "",
}

func init() {
	rootCmd.AddCommand(petsCmd)
}
