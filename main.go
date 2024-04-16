package main

import (
	"navidmafi/walli/cmd"
	"navidmafi/walli/config"
	"navidmafi/walli/logger"
	"navidmafi/walli/validation"
	"os"

	"github.com/urfave/cli/v2"
)

var version string

func main() {

	logger.Init()
	config.Init()
	validation.Init()

	// swwwBackend := backends.NewSwwwBackend()
	// buf, readErr := os.ReadFile("/mnt/Data/Wallpapers and Photos/pexels-ron-lach-9999332.jpg")
	// if readErr != nil {
	// 	logger.Logger.Fatal(readErr)
	// }
	// backendErr := backends.Backends[backends.Gnome].Apply(*bytes.NewBuffer(buf))
	// swwwBackend.Apply(*bytes.NewBuffer(buf))
	// afasfderr := swwwBackend.ApplyFile("/mnt/Data/Wallpapers and Photos/pexels-ron-lach-9999332.jpg")
	// logger.Logger.Fatal(backendErr)
	// fmt.Println(readErr)
	// os.Exit(0)

	app := cli.App{

		Name:        "walli",
		Version:     version,
		Usage:       "PotD engine for linux desktops",
		Description: "",
		Commands:    []*cli.Command{&cmd.ProviderCmd, &cmd.BackendCmd, &cmd.WallpaperCmd},
	}

	app.Run(os.Args)
	logger.Logger.Info("WALLI IS BETA SOFTWARE")

}
