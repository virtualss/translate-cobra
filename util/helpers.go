package util

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"translate-cobra/pkg/types"
	utilerrors "translate-cobra/util/errors"
)

type ConfigReadChain struct {
	Path string
	Flag int
	Perm os.FileMode
	Next *ConfigReadChain
}

func (r *ConfigReadChain) read() (*os.File, error) {
	file, err := os.OpenFile(r.Path, r.Flag, r.Perm)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (r *ConfigReadChain) ProcessRead() *os.File {
	file, err := r.read()
	if err != nil {
		if r.Next != nil {
			return r.Next.ProcessRead()
		} else {
			return nil
		}
	} else {
		return file
	}
}

// CfgRead read from `CLI -> ./tl.json -> homedir`
func CfgRead(cfg string) (*types.TranslateConfig, error) {
	chain := &ConfigReadChain{
		// cli
		Path: cfg, Flag: os.O_RDONLY, Perm: 0666,
		// ./tl.json
		Next: &ConfigReadChain{Path: "./tl.json", Flag: os.O_RDWR, Perm: 0666}}

	// home dir
	dir, err := os.UserHomeDir()
	if err == nil {
		chain.Next.Next = &ConfigReadChain{Path: dir + string(os.PathSeparator) + "tl.json", Flag: os.O_RDWR, Perm: 0666}
	}
	c := chain.ProcessRead()
	defer func(c *os.File) {
		_ = c.Close()
	}(c)
	if c == nil {
		cobra.CheckErr("missing config file")
	}
	t := new(types.TranslateConfig)

	err = json.NewDecoder(c).Decode(t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func ValueCheck(config *types.TranslateConfig) error {
	// domain check
	supportDomain, domainSupport := types.Domains.Support(config.Domain)
	if !domainSupport {
		// language (FROM TO) check
		return utilerrors.Domain{Field: config.Domain}
	}

	// language FROM TO check
	if !types.Languages.Support(supportDomain, config.From, config.To) {
		return utilerrors.FromTo{From: config.From, To: config.To}
	}
	return nil
}

func UsageErrorf(cmd *cobra.Command, format string, args ...interface{}) error {
	msg := fmt.Sprintf(format, args...)
	return fmt.Errorf("%s\nSee '%s -h' for help and examples", msg, cmd.CommandPath())
}
