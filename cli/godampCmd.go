package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var godampCmd = &cobra.Command{
	Use:   "godamp",
	Short: "Godamp is a minimal tcpdump written in go",
}

func Execute() {
	monitorCmd.Flags().StringVarP(&filter, "filter", "f", "", "Add filter for packets")
	godampCmd.AddCommand(listCmd)
	godampCmd.AddCommand(monitorCmd)

	if err := godampCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
