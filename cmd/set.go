package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

func NewSetCmd() *cobra.Command {
	setCmd := &cobra.Command{
		Use:                   "set SUBCOMMAND",
		DisableFlagsInUseLine: true,
		Short:                 "Set appid or secret",
		Long:                  "Set translate api appid or secret",
		Run: func(c *cobra.Command, args []string) {
			c.SetOut(os.Stderr)
			c.SetErr(os.Stderr)
			_ = c.Help()
		},
	}

	setCmd.AddCommand(setSecretCmd)
	return setCmd
}

// setSecretCmd secret subcommand
var setSecretCmd = &cobra.Command{
	Use:                   "secret (secret=SECRET_NAME)",
	DisableFlagsInUseLine: true,
	Short:                 "baidu translate api secret",
	Long:                  "baidu translate api secret",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("this is set secret", args)
	},
}

// setAppIdCmd appid subcommand
var setAppIdCmd = &cobra.Command{
	Use:                "appid (appid=APPID)",
	DisableSuggestions: true,
	Short:              "baidu translate api appid",
	Long:               "baidu translate api appid",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
