/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/iwanhae/kaustar/pkg/config"
	"github.com/iwanhae/kaustar/pkg/server"
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

	e, err := server.NewServer(cfg)
	if err != nil {
		return err
	}

	go func() {
		fmt.Println("Starting kaustar controller")
		runController()
	}()
	fmt.Println("Starting kaustar")
	log.Fatal(e.Start(cfg.Server.Address))
	return nil
}
