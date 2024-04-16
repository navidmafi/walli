package cmd

import (
	"fmt"
	"navidmafi/walli/providers"
	"strings"

	"github.com/urfave/cli/v2"
)

var ProviderCmd = cli.Command{
	Name:        "provider",
	Usage:       "configure providers and set the default",
	Aliases:     []string{"p"},
	UsageText:   fmt.Sprintf("walli provider [command] [provider] \n\nAvailable providers: %s", strings.Join(providers.GetAvailable(), ", ")),
	Subcommands: []*cli.Command{&ProviderSetCmd, &ProviderAuthCmd},
}
