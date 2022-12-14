package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/bigquery"
	"github.com/ahmadrosid/bq-cli/cmd"
	"github.com/ahmadrosid/bq-cli/service"
	"github.com/urfave/cli/v2"
)

func main() {
	projectId := os.Getenv("GOOGLE_PROJECT_ID")
	if projectId == "" {
		fmt.Printf("please provide env GOOGLE_PROJECT_ID\n")
		os.Exit(1)
	}

	bQClient, err := bigquery.NewClient(context.Background(), projectId)
	if err != nil {
		panic(err)
	}

	svc := service.NewBiqueryService(bQClient)
	ui := service.NewReadlineUI()
	bqCmd := cmd.NewBiqueryCommand(svc, ui)

	app := &cli.App{
		Name:  "bq-cli",
		Usage: "A cli app to execute bigquery from terminal.",
		Commands: []*cli.Command{
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Execute biquery",
				Action:  bqCmd.HandleQuery,
			},
			{
				Name:    "repl",
				Aliases: []string{"i"},
				Usage:   "Run interactive query",
				Action:  bqCmd.HandleInteractive,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
