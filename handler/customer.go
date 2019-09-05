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

// @Description get customer info by ID
// @Tags customers
// @Accept  json
// @Produce  json
// @Param   id     path    string     true        "Some ID"
// @Success 200 {string} string	"ok"
// @Failure 404 {object} handler.StatusError
// @Router /customers/get/{id} [get]
func (handler customerHandlerImpl) get(c iris.Context) {
	id := c.Params().GetUintDefault("id", 0)
	customer, err := handler.customerRepo.Find(id)
	if err != nil {
		simpleReturnHandler(c, customer, NewNotFound(err))
		return
	}
	simpleReturnHandler(c, customer, nil)
}
