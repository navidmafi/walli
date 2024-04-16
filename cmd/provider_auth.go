package cmd

import (
	"bufio"
	"fmt"
	"navidmafi/walli/config"
	"navidmafi/walli/logger"
	"navidmafi/walli/providers"
	"navidmafi/walli/validation"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

var ProviderAuthCmd = cli.Command{
	Name:      "auth",
	Usage:     "Authenticate with a provider",
	Aliases:   []string{"a"},
	UsageText: fmt.Sprintf("walli provider auth [%s]", strings.Join(providers.GetAvailable(), "|")),
	Action: func(ctx *cli.Context) error {
		targetProvider := ctx.Args().First()

		err := validation.Validate.Var(targetProvider, "validProvider")
		if err != nil {
			logger.Logger.Fatalf("Usage: %s", ctx.Command.UsageText)
		}

		var targetProviderImpl = providers.Providers[providers.ProviderName(targetProvider)]

		logger.Logger.Infof("%s\n%s", targetProviderImpl.GetAuthURL(), targetProviderImpl.GetAuthHelp())

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		token := scanner.Text()

		authErr := targetProviderImpl.Authenticate(providers.Authentication{
			Strategy: targetProviderImpl.GetSupportedAuthStrategy(),
			Secret:   token,
		})

		if authErr != nil {
			logger.Logger.Fatalf("Internal Error. please open an issue. %s", authErr)
		}

		testResult, err := targetProviderImpl.GetRandom("black cat")

		if err != nil {
			logger.Logger.Fatalf("Auth failed. please try again. %s", err)
		}

		logger.Logger.Debug(testResult)
		logger.Logger.Infof("Authenticated successfully with %s.", targetProvider)
		config.SetSecret(targetProvider, token)
		return nil
	},
}
