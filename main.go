package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/ptloc1992/boilerplate-microservice/cmd"
	"github.com/urfave/cli"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := cli.NewApp()

	app.Name = "Micro-service Boilerplate"
	app.Usage = "Go Micro-service Boilerplate"
	app.Version = "v1.0.0"

	// GLOBAL OPTIONS
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "database, db",
			Value:  "root:123@/notes?charset=utf8&parseTime=True&loc=Local",
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
		cmd.Test,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Panic(err)
		log.Fatal(err)
	}
}
