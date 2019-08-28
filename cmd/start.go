package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/ptloc1992/boilerplate-microservice/server"
	"github.com/urfave/cli"
	"go.uber.org/fx"
)

// handleHTTPServer handles http server
func handleHTTPServer(lc fx.Lifecycle, appContext *cli.Context, logger *log.Logger, db *gorm.DB) {
	//irisApp := handler.BuildEngine(appContext, logger, db)
	s := server.Server{
		//IrisApp: irisApp,
		Address: appContext.String("address"),
		Port:    appContext.String("port"),
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			fmt.Println("!Start")
			go s.Start(appContext)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			fmt.Println("!Stop")
			s.Stop()
			return nil
		},
	})
}

func startAction(appContext *cli.Context) {
	app := fx.New(
		fx.Provide(
			func() *cli.Context {
				return appContext
			},
			newLogger,
			newDB,
		),
		fx.StopTimeout(time.Second*30),
		fx.Invoke(handleHTTPServer),
	)
	app.Run()
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
