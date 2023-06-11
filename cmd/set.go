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
		Short:                 "set appid or secret",
		Long:                  "set translate api appid or secret",
		Run: func(c *cobra.Command, args []string) {
			c.SetOut(os.Stderr)
			c.SetErr(os.Stderr)
			_ = c.Help()
		},
	}
	// sub command 'secret'
	setCmd.AddCommand(setSecretCmd)
	// sub command 'appid'
	setCmd.AddCommand(setAppIdCmd)
	return setCmd
}

// setSecretCmd secret subcommand
var setSecretCmd = &cobra.Command{
	Use:                   "secret (secret=SECRET_NAME)",
	DisableFlagsInUseLine: true,
	Short:                 "baidu translate api secret",
	Long:                  "baidu translate api secret",
	// at least one arg is needed
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO set secret
		fmt.Println("this is set secret", args)
	},
}

// setAppIdCmd appid subcommand
var setAppIdCmd = &cobra.Command{
	Use:                   "appid (appid=APPID)",
	DisableFlagsInUseLine: true,
	Short:                 "baidu translate api appid",
	Long:                  "baidu translate api appid",
	// at least one arg is needed
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// TODO set appid
		fmt.Println("this is appid", args)
	},
}
