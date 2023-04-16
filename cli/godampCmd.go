package cli

import (
	"fmt"
	"github.com/krak7602/godamp/internals"
	"github.com/spf13/cobra"
	"os"
)

var godampCmd = &cobra.Command{
	Use:   "tcpdumpgo <device>",
	Short: "Godamp is a minimal TCPdump written in go",
	Args:  cobra.ExactArgs(1),
	Run:   internals.Godamp,
}

func Execute() {
	if err := godampCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
