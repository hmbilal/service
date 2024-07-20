package health

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
	healthGroup := app.Group("/health")

	healthGroup.Get("/ping", r.handler.Ping)
	healthGroup.Get("/details", r.handler.Details)
}
