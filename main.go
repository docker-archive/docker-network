package main

import (
	"os"

	"github.com/docker/docker-network/Godeps/_workspace/src/github.com/codegangsta/cli"
)

var (
	networkCommand = cli.Command{
		Name:      "network",
		ShortName: "n",
		Usage:     "Manage networks",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "Shows list of created networks",
			},
			{
				Name:  "add",
				Usage: "Add network",
			},
			{
				Name:  "del",
				Usage: "Delete network",
			},
			{
				Name:  "show",
				Usage: "Shows info about network",
			},
		},
	}
	epCommand = cli.Command{
		Name:      "endpoint",
		ShortName: "ep",
		Usage:     "Manage endpoints in network",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "Shows list of endpoinds in network",
			},
			{
				Name:  "add",
				Usage: "Add endpoint to network",
			},
			{
				Name:  "del",
				Usage: "Delete endpoint from network",
			},
			{
				Name:  "show",
				Usage: "Shows info about endpoint",
			},
		},
	}
	nsCommand = cli.Command{
		Name:      "namespace",
		ShortName: "ns",
		Usage:     "Manage network namespaces",
		Subcommands: []cli.Command{
			{
				Name:  "list",
				Usage: "List of network namespaces which belongs to docknet",
			},
			{
				Name:  "add",
				Usage: "Add new network namespace",
			},
			{
				Name:  "del",
				Usage: "Delete network namespace",
			},
			{
				Name:  "join",
				Usage: "Join endpoint to specified namespace (this can be docknet namespace or path)",
			},
			{
				Name:  "exec",
				Usage: "Execute command in namespace",
			},
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "docknet"
	app.Usage = "Create and manage networks"
	app.Action = func(c *cli.Context) {
		println(app.Usage)
	}
	app.Commands = []cli.Command{
		networkCommand,
		epCommand,
		nsCommand,
	}
	app.Run(os.Args)
}
