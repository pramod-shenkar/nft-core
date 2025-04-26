package rest

import (
	"nftledger/common/server"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RegisterRoute(a *server.App, handlers *Handlers) {

	a.Use(logger.New())
	v1 := a.Group("v1")
	{
		request := v1.Group("request")
		{
			request.Get("/:id", handlers.Request.Get)
			request.Get("/", handlers.Request.GetAll)
			request.Post("/", handlers.Request.Save)
			request.Patch("/:id", handlers.Request.Update)
			request.Delete("/", handlers.Request.Delete)
		}
	}

}
