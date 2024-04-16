package cmd

import (
	"fmt"
	"navidmafi/walli/backends"
	"strings"

	"github.com/urfave/cli/v2"
)

var BackendCmd = cli.Command{
	Name:        "backend",
	Usage:       "configure backends and set the default",
	Aliases:     []string{"b"},
	UsageText:   fmt.Sprintf("walli backend [command] [backend] \n\nAvailable backends: %s", strings.Join(backends.GetAvailable(), ", ")),
	Subcommands: []*cli.Command{&BackendSetCmd},
}
