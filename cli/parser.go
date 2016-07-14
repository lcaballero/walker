package cli

import (
	"github.com/jessevdk/go-flags"
	"github.com/lcaballero/walker/conf"
	"os"
)

func ParseArgs(args ...string) *conf.Config {
	opts := &conf.Config{}
	parser := flags.NewParser(opts, flags.Default)
	_, err := parser.ParseArgs(args)
	if err != nil {
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
	return opts
}
