package api

import "github.com/gofiber/fiber/v2"

type Router interface {
	RegisterRoutes(app *fiber.App)
}

func RegisterRouters(app *fiber.App, routers ...Router) {
	for _, r := range routers {
		r.RegisterRoutes(app)
	}
}
