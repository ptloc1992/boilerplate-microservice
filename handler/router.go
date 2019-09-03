package handler

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/urfave/cli"
)

func BuildEngine(appContext *cli.Context, logger *log.Logger, db *gorm.DB) *iris.Application {
	app := iris.Default()
	app.Logger().SetLevel(appContext.GlobalString("loglevel"))
	// your handler
	// 1. healCheck
	healCheckHandlerImpl := healthCheckHandlerImpl{
		log: logger,
	}
	healCheckHandlerImpl.inject(app)
	// // HTTP ReqResIn service
	// reqResInHanler := reqResInHandlerImpl{
	// 	reqResService: service.NewReqResIn("https://reqres.in"),
	// 	log:           logger,
	// }
	return app
}
