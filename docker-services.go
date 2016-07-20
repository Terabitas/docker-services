package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/nildev/docker-services/config"
	"github.com/nildev/docker-services/version"
	"github.com/nildev/lib/log"
)

func main() {
	app := cli.NewApp()
	app.Name = "docker-services"
	app.Usage = "Tool which allows you to define and run dockerized applications"
	app.Version = version.Version
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "verbosity",
			Value: 0,
			Usage: "logging level",
		},
		cli.StringFlag{
			Name:  "dir",
			Value: "",
			Usage: "root directory where to store configuration files",
		},
		cli.StringFlag{
			Name:  "config",
			Value: "",
			Usage: "path to config file to use",
		},
	}
	app.Before = func(c *cli.Context) error {
		err := config.LoadConfig(c.String("config"))
		if err != nil {
			log.Fatal(err)
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "setup",
			Aliases: []string{"i"},
			Usage:   "setup services",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) {
				// create directories
				// download `docker-compose.yml` files
			},
		},
		{
			Name:    "start",
			Aliases: []string{"s"},
			Usage:   "start services",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) {

			},
		},
		{
			Name:    "stop",
			Aliases: []string{"t"},
			Usage:   "stop services",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) {

			},
		},
		{
			Name:    "update",
			Aliases: []string{"u"},
			Usage:   "update services",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) {

			},
		},
		{
			Name:    "remove",
			Aliases: []string{"rm"},
			Usage:   "remove services",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) {

			},
		},
		{
			Name:    "describe",
			Aliases: []string{"d"},
			Usage:   "describe service, print out available tags, datasets",
			Flags:   []cli.Flag{},
			Action: func(c *cli.Context) {

			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Debugf("Could not run docker-services cmd, %s", err)
	}
}
