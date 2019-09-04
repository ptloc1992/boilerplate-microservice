package handler

import (
	"log"

	"github.com/iris-contrib/swagger"
	"github.com/iris-contrib/swagger/swaggerFiles"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	_ "github.com/ptloc1992/boilerplate-microservice/docs"
	"github.com/ptloc1992/boilerplate-microservice/repo"
	"github.com/urfave/cli"
)

func BuildEngine(appContext *cli.Context, logger *log.Logger, db *gorm.DB) *iris.Application {
	app := iris.Default()
	app.Logger().SetLevel(appContext.GlobalString("loglevel"))
	// config := &swagger.Config{
	// 	URL: "http://localhost:8686/swagger/doc.json", //The url pointing to API definition
	// }
	// use swagger middleware to
	//app.Get("/swagger/*any}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
	app.Get("/swagger/*any}", swagger.WrapHandler(swaggerFiles.Handler))
	// HealCheck
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
