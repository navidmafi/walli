package cmd

import (
	"fmt"
	"navidmafi/walli/config"
	"navidmafi/walli/logger"
	"navidmafi/walli/providers"
	"navidmafi/walli/validation"
	"strings"

	"github.com/urfave/cli/v2"
)

var ProviderSetCmd = cli.Command{
	Name:      "set",
	Aliases:   []string{"s"},
	Usage:     "Set the default provider",
	UsageText: fmt.Sprintf("walli provider set [%s]", strings.Join(providers.GetAvailable(), "|")),
	Action: func(ctx *cli.Context) error {
		targetProvider := ctx.Args().First()

		err := validation.Validate.Var(targetProvider, "validProvider")
		if err != nil {
			logger.Logger.Fatalf("Usage: %s", ctx.Command.UsageText)
		}

		config.SetDefaultProvider(targetProvider)

		fmt.Printf("%s is now the default provider\n", targetProvider)
		return nil

	},
}
