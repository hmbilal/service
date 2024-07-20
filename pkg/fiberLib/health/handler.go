package health

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hmbilal/gofiber-start/pkg/checker"
	"net/http"
)

type Handler struct {
	checkerPool checker.Pool
}

func NewHandler(pool checker.Pool) *Handler {
	return &Handler{checkerPool: pool}
}

func (h *Handler) Ping(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(map[string]interface{}{"ping": h.checkerPool.Status()})
}

func (h *Handler) Details(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(h.checkerPool.Details())
}
