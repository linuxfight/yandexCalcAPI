package handlers

import "github.com/gofiber/fiber/v3"

func SolveHandler(c fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
	})
}
