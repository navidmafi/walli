package cmd

import "github.com/urfave/cli/v2"

var WallpaperCmd = cli.Command{
	Name:        "wallpaper",
	Aliases:     []string{"w"},
	Usage:       "set and apply wallpapers",
	Subcommands: []*cli.Command{&WallpaperSetCmd},
}
