package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/hyperscale/hyperdash/cmd/hyperdash/app/commands"
	"github.com/hyperscale/hyperdash/pkg/hyperdash/version"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	cli.VersionPrinter = func(c *cli.Context) {
		_, _ = fmt.Fprintf(c.App.Writer, "%v version: %v (build: %v)\n", c.App.Name, c.App.Version, version.Get().BuildDate)
		_, _ = fmt.Fprintf(c.App.Writer, "Go runtime version: %v\n", version.Get().GoVersion)
		_, _ = fmt.Fprintf(c.App.Writer, "Platform: %v\n", version.Get().Platform)
	}

	// Get the command line args.
	binName := filepath.Base(os.Args[0])

	app := &cli.App{
		Name:    binName,
		Version: version.Get().Version,
		//Description: "Generic automation framework for acceptance testing.",
		Usage:     "Generate metric dashboard",
		Copyright: "Axel Etcheverry",
		Flags:     []cli.Flag{},
		Commands: []*cli.Command{
			{
				Name:   "run",
				Usage:  "run dashboard",
				Action: commands.RunAction,
			},
			{
				Name: "provider",
				Subcommands: []*cli.Command{
					{
						Name:   "ls",
						Usage:  "list all available provider",
						Action: commands.ProviderListAction,
					},
				},
			},
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
