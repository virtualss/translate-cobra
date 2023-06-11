package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"translate-cobra/pkg"
	"translate-cobra/pkg/types"
	"translate-cobra/util"
)

var (
	cfg          string
	q            string
	translateCfg = &types.TranslateConfig{
		From:   "en",
		To:     "zh",
		Domain: "it",
	}
)

var rootCmd = &cobra.Command{
	Use:   "tl",
	Short: "tl is a CLI for translate",
	Long:  `A translate tool provide translate capability power by baidu translate api`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO root cmd
		trimStr := strings.TrimSpace(q)
		pkg.DoRequest(trimStr, translateCfg)
		//fmt.Println("run", args)
	},
}

func init() {
	cobra.OnInitialize(initCfg)
	rootCmd.Flags().StringVarP(&cfg, "configFile", "c", "", "Default config file location (HOME_DIR/tl.json)")
	rootCmd.Flags().StringVarP(&translateCfg.From, "from", "f", "en", "Origin language(en or zh)")
	rootCmd.Flags().StringVarP(&translateCfg.To, "to", "t", "zh", "Target language(en or zh)")
	rootCmd.Flags().StringVarP(&translateCfg.Domain, "domain", "d", "it", "Translate field (see also 'tl field' for more details)")
	rootCmd.Flags().StringVarP(&q, "query", "q", "", "The words that should be translated")
	_ = rootCmd.MarkFlagRequired("query")
	rootCmd.AddCommand(NewVersionCmd())
	rootCmd.AddCommand(NewField())
	rootCmd.AddCommand(NewSetCmd())
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

// initCfg is called after init function
func initCfg() {
	// read from CLI -> ./tl.json -> HOME_DIR/tl.json
	t, err := util.CfgRead(cfg)
	cobra.CheckErr(err)

	err = util.ValueCheck(t)

	if err != nil {
		cobra.CheckErr(err)
	}
	translateCfg = t
}
