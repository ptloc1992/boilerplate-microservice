package handler

import (
	"log"

	"github.com/kataras/iris"
	"github.com/ptloc1992/boilerplate-microservice/repo"
)

type customerHandlerImpl struct {
	customerRepo repo.CustomerRepo
	log          *log.Logger
}

func (handler customerHandlerImpl) inject(app *iris.Application) {
	group := app.Party("/customers")
	group.Get("/{id:uint}", handler.get)
}

func (handler customerHandlerImpl) get(c iris.Context) {
	id := c.Params().GetUintDefault("id", 0)
	customer, err := handler.customerRepo.Find(id)
	if err != nil {
		simpleReturnHandler(c, customer, NewNotFound(err))
		return
	}
	simpleReturnHandler(c, customer, nil)
}
