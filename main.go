package main

import (
	"log"
	"os"

	"github.com/alcb1310/go-cli/utils"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:        "go-cli",
		Description: "go-cli is a tool that will allow you to interact with your git branches more easily",
		Copyright:   "(c) 2023 - 2024 Andrés Court <andres@andrescourt.com>\ngithub: https://github.com/alcb1310/go-cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "name",
				Aliases: []string{"n"},
				Usage:   "name of the new branch",
			},
		},

		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	if c.String("name") != "" {
		utils.Create(c)
		return nil
	}

	utils.List(c)
	return nil
}
