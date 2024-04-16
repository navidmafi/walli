package cmd

import (
	"navidmafi/walli/backends"
	"navidmafi/walli/config"
	"navidmafi/walli/logger"
	"navidmafi/walli/networking"
	"navidmafi/walli/providers"
	"strings"

	"github.com/urfave/cli/v2"
)

var WallpaperSetCmd = cli.Command{
	Name:      "set",
	Aliases:   []string{"s"},
	Usage:     "set a wallpaper based on the query",
	UsageText: "walli wallpaper set [query]",
	Action: func(ctx *cli.Context) error {

		query := strings.Join(ctx.Args().Slice(), " ")

		providerName := config.GetDefaultProvider()
		backendName := config.GetDefaultBackend()

		providerImpl := providers.Providers[providers.ProviderName(providerName)]
		backendImpl := backends.Backends[backends.BackendName(backendName)]

		providerSecret := config.GetSecret(providerName)

		if len(providerSecret) == 0 {
			logger.Logger.Fatalf("%[1]s provider needs authentication. use : walli auth %[1]s", providerName)
		}

		authErr := providerImpl.Authenticate(providers.Authentication{
			Strategy: providerImpl.GetSupportedAuthStrategy(),
			Secret:   providerSecret,
		})

		if authErr != nil {
			logger.Logger.Fatal("Something went wrong")
		}

		randomImgURL, err := providerImpl.GetRandom(query)

		if err != nil {
			logger.Logger.Fatal(err)
		}

		bufDownload := networking.FetchBinary(networking.RequestCtx{
			URL: randomImgURL,
		})

		backendImpl.Apply(bufDownload)

		return nil
	},
}
