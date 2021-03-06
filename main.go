package main

//go:generate ./version.sh

import (
	"os"

	"github.com/degica/barcelona-cli/cmd"
	"github.com/degica/barcelona-cli/config"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "bcn"
	app.Version = Version
	app.Usage = "Barcelona Command Line Client"
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug, d",
			Usage:       "Enable debug mode",
			Destination: &config.Debug,
		},
	}
	app.Commands = []cli.Command{
		cmd.LoginCommand,
		cmd.DeployCommand,
		cmd.CreateCommand,
		cmd.DistrictCommand,
		cmd.EndpointCommand,
		cmd.APICommand,
		cmd.EnvCommand,
		cmd.RunCommand,
		cmd.SSHCommand,
		cmd.ReleaseCommand,
		cmd.NotificationCommand,
		cmd.AppCommand,
	}

	app.Run(os.Args)
}
