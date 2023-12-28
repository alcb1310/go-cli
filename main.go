package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "go-cli",
		Description: "go-cli is a tool that will allow you to interact with your git branches more easily",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
