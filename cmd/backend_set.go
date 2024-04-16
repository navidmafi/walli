package cmd

import (
	"fmt"
	"navidmafi/walli/backends"
	"navidmafi/walli/config"
	"navidmafi/walli/logger"
	"navidmafi/walli/validation"
	"strings"

	"github.com/urfave/cli/v2"
)

var BackendSetCmd = cli.Command{
	Name:      "set",
	Usage:     "Set the default backend",
	Aliases:   []string{"s"},
	UsageText: fmt.Sprintf("walli backend set [%s]", strings.Join(backends.GetAvailable(), "|")),
	Action: func(ctx *cli.Context) error {
		targetBackend := ctx.Args().First()

		err := validation.Validate.Var(targetBackend, "validBackend")
		if err != nil {
			logger.Logger.Fatalf("Usage: %s", ctx.Command.UsageText)
		}

		config.SetDefaultBackend(targetBackend)

		fmt.Printf("%s is now the default backend\n", targetBackend)
		return nil

	},
}
