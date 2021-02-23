package main

import (
	"os"

	"github.com/Trois-Six/fauna-exporter/cmd/exporter"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "fauna-exporter CLI",
		Usage: "Run fauna-exporter",
		Commands: []*cli.Command{
			exporterCommand(),
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("Error while  executing command")
	}
}

func exporterCommand() *cli.Command {
	return &cli.Command{
		Name:        "exporter",
		Usage:       "Export Fauna metrics",
		Description: "Export Fauna metrics from the Fauna dashboard.",
		Action:      exporter.Run,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "host",
				Usage:   "Host to listen on.",
				EnvVars: []string{"HOST"},
			},
			&cli.StringFlag{
				Name:    "fauna-email",
				Usage:   "Email to connect to the Fauna dashboard.",
				EnvVars: []string{"FAUNA_EMAIL"},
			},
			&cli.StringFlag{
				Name:    "fauna-password",
				Usage:   "Password to connect to the Fauna dashboard.",
				EnvVars: []string{"FAUNA_PASSWORD"},
			},
			&cli.IntFlag{
				Name:    "fauna-days",
				Usage:   "Number of days of the metrics.",
				EnvVars: []string{"FAUNA_DAYS"},
				Value:   7,
			},
			&cli.StringFlag{
				Name:    "metrics-path",
				Usage:   "Path of the metrics.",
				EnvVars: []string{"METRICS_PATH"},
				Value:   "/metrics",
			},
		},
	}
}
