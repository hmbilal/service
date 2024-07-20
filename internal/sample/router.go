package sample

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmbilal/gofiber-start/pkg/fiberLib/api"
)

type router struct {
	handler *Handler
}

func NewRouter(handler *Handler) api.Router {
	return &router{handler: handler}
}

func (r *router) RegisterRoutes(app *fiber.App) {
	sampleGroup := app.Group("/sample")

	sampleGroup.Get("/get", r.handler.hello)
}
