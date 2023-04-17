package cli

import (
	"github.com/krak7602/godamp/internals"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor",
	Short: "Monitor a interface activity",
	Args:  cobra.ExactArgs(1),
	Run:   internals.Monitor,
}
