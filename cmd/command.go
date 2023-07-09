package cmd

import (
	"github.com/pramuditorh/godrop/actions"
	"github.com/urfave/cli/v2"
)

func Commands() []*cli.Command {
	var commands = []*cli.Command{
		{
			Name: "provision",
			Aliases: []string{"prov"},
			Description: "Provision Digital Ocean Droplet",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "name",
					Usage: "Provide Droplet name",
					Required: true,
				},
				&cli.StringFlag{
					Name: "region",
					Value: "sgp1",
					Usage: "Provide Droplet region",
				},
				&cli.StringFlag{
					Name: "size",
					Usage: "Provide Droplet size",
					Required: true,
				},
				&cli.StringFlag{
					Name: "image",
					Usage: "Provide Droplet image",
					Required: true,
				},
				&cli.IntFlag{
					Name: "ssh-key-id",
					Usage: "Provide Droplet ssh-key ID",
				},
			},
			Action: func(cCtx *cli.Context) error {
				return actions.Provision(cCtx)
			},
		},
		{
			Name: "list",
			Aliases: []string{"ls"},
			Description: "List your Digital Ocean objects",
			Subcommands: []*cli.Command{
				{
					Name: "droplets",
					Aliases: []string{"drp"},
					Description: "List your running Droplets",
					Action: func(cCtx *cli.Context) error {
						return actions.ListDroplets(cCtx)
					},
				},
				{
					Name: "regions",
					Aliases: []string{"reg"},
					Description: "List Digital Ocean Regions",
					Action: func(cCtx *cli.Context) error {
						return actions.ListRegions(cCtx)
					},
				},
			},
		},

	}

	return commands
}