package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/linuxfight/yandexCalcApi/internal/logger"
)

func Recovery() fiber.Handler {
	return func(c fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				logger.Log.Errorf("panic occurred: %v", err)

				fiberErr := c.Status(fiber.StatusInternalServerError).JSON(
					&errorResponse{Error: "internal server error"},
				)
				if fiberErr != nil {
					logger.Log.Errorf("fiber error: %v", fiberErr)
				}
			}
		}()

		return c.Next()
	}
}
