package main

import (
	"log"
	"os"

	"./cmd"
	"github.com/joho/godotenv"
	"github.com/urfave/cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := cli.NewApp()

	app.Name = "Ten-Project-cua-Ban"
	app.Usage = "Mo-ta-ve-Project-cua-ban-o-day-nhe"
	app.Version = "v1"

	// GLOBAL OPTIONS
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "database, db",
			Value:  "default:secret@/notes?charset=utf8&parseTime=True&loc=Local",
			Usage:  "Database connection",
			EnvVar: "DATABASE",
		},
		cli.StringFlag{
			Name:   "loglevel",
			Value:  "debug",
			Usage:  "logging level",
			EnvVar: "LOG_LEVEL",
		},
	}

	app.Commands = []cli.Command{
		cmd.Start,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
