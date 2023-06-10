package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Print the version number of 'tl'",
		Long:    "this is tl's version",
		Example: `tl version`,
		Run: func(cmd *cobra.Command, args []string) {
			_, _ = fmt.Fprintln(os.Stdout, "tl version: [v0.1]")
		},
	}
}
