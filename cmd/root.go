/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/iwanhae/kaustar/pkg/config"
	"github.com/iwanhae/kaustar/pkg/server/proxy"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var cfgFile *string

func Execute() {
	cfgFile = rootCmd.Flags().StringP("config", "c", "./default.yaml", "config file")
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var rootCmd = &cobra.Command{
	Use:          "kaustar",
	Short:        "Reverse proxy of Kubernetes API server for Audit, Authorization, and Authentication",
	RunE:         run,
	SilenceUsage: true,
}

func run(cmd *cobra.Command, args []string) error {
	cfg, err := config.Load(*cfgFile)
	if err != nil {
		return err
	}

	e := echo.New()
	e.HidePort = true
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	{ // Porxy to Kubernetes API server
		config, err := proxy.LoadConfig(cfg.Kubernetes)
		if err != nil {
			return errors.Wrap(err, "failed to load kubernetes config")
		}
		e.Any("/proxy/*", proxy.KubeAPIServer(config),
			middleware.AddTrailingSlashWithConfig(middleware.TrailingSlashConfig{Skipper: middleware.DefaultSkipper, RedirectCode: 301}),
			middleware.Rewrite(map[string]string{"/proxy/*": "/$1"}),
		)
	}

	e.Logger.Info("Starting kaustar")
	e.Logger.Fatal(e.Start(cfg.Server.Address))
	return nil
}
