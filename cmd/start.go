package cmd

import (
	"fmt"

	"github.com/urfave/cli"
)

func startAction(appContext *cli.Context) {
	fmt.Println("Xin chao ban loc")
}

var Start = cli.Command{
	Name:  "start",
	Usage: "Khoi-dong-chuong-trinh",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:   "port, p",
			Value:  "8080",
			Usage:  "Listen on port",
			EnvVar: "PORT",
		},
	},
	Action: func(appContext *cli.Context) error {
		startAction(appContext)
		return nil
	},
}
