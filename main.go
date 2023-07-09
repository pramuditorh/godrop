package main

import (
	"log"
	"os"

	"github.com/pramuditorh/godrop/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "godrop",
		Usage: "Provision DigitalOcean Droplet",
		Version: "0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Robby Hemawan Pramudito",
				Email: "pramuditorh@gmail.com",
			},
		},
		Commands: cmd.Commands(),
	}


	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}