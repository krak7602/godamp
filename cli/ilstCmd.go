package cli

import (
	"github.com/krak7602/godamp/internals"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available interfaces",
	Args:  cobra.NoArgs,
	Run:   internals.List,
}
