package handler

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"github.com/ptloc1992/boilerplate-microservice/repo"
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
	// Customer handler
	customerHandlerImpl := customerHandlerImpl{
		customerRepo: repo.CustomerRepoImpl{
			DB: db,
		},
		log: logger,
	}
	customerHandlerImpl.inject(app)
	// // HTTP ReqResIn service
	// reqResInHanler := reqResInHandlerImpl{
	// 	reqResService: service.NewReqResIn("https://reqres.in"),
	// 	log:           logger,
	// }
	return app
}

func simpleReturnHandler(c iris.Context, result interface{}, err Error) {
	if err != nil {
		c.StatusCode(err.Status())
		c.JSON(iris.Map{
			"error": err.Error(),
		})
		return
	}
	c.JSON(result)
}
