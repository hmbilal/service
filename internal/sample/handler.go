package sample

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmbilal/gofiber-start/pkg/container"
)

type Handler struct {
	container *container.Container
}

func NewHandler(container *container.Container) *Handler {
	return &Handler{
		container: container,
	}
}

func (h *Handler) hello(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello, World!")
}
