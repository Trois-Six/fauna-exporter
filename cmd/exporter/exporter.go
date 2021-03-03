// Package exporter contains the exporter cmd.
package exporter

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Trois-Six/fauna-exporter/pkg/exporter"
	"github.com/Trois-Six/fauna-exporter/pkg/handlers"
	"github.com/Trois-Six/fauna-exporter/pkg/logger"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/urfave/cli/v2"
)

// Run executes user service.
func Run(cliContext *cli.Context) error {
	logger.Setup()

	if err := setupEnvVars(cliContext); err != nil {
		return err
	}

	prometheus.MustRegister(version.NewCollector("fauna_exporter"))
	prometheus.MustRegister(exporter.NewExporter(
		cliContext.Int("fauna-days"),
		cliContext.String("fauna-email"),
		cliContext.String("fauna-password"),
	))

	mux := http.NewServeMux()
	mux.Handle(cliContext.String("metrics-path"), promhttp.InstrumentMetricHandler(
		prometheus.DefaultRegisterer,
		promhttp.HandlerFor(
			prometheus.DefaultGatherer,
			promhttp.HandlerOpts{},
		),
	))

	handler := handlers.New(cliContext.String("metrics-path"))
	mux.HandleFunc("/", handler.Index)
	mux.HandleFunc("/-/healthy", handler.OK)
	mux.HandleFunc("/-/ready", handler.OK)

	return http.ListenAndServe(cliContext.String("host"), mux)
}

func setupEnvVars(context *cli.Context) error {
	vars := map[string]string{
		"FAUNA_EMAIL":    "fauna-email",
		"FAUNA_PASSWORD": "fauna-password",
		"FAUNA_DAYS":     "fauna-days",
		"METRICS_PATH":   "metrics-path",
	}

	for name, flag := range vars {
		if err := os.Setenv(name, context.String(flag)); err != nil {
			return fmt.Errorf("failed to set environment variable: %w", err)
		}
	}

	return nil
}
