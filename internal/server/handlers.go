package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/linuxfight/yandexCalcApi/pkg/calc"
	"net/http"
)

func SolveHandler(c fiber.Ctx) error {
	if c.Get("Content-Type") != "application/json" {
		return c.Status(http.StatusUnprocessableEntity).JSON(
			&errorResponse{Error: notJson})
	}

	var body solveRequest
	if err := c.Bind().JSON(&body); err != nil || body.Expression == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			&errorResponse{Error: invalidJson},
		)
	}

	result, solveErr := calc.Solve(body.Expression)
	if solveErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			&errorResponse{Error: solveErr.Error()},
		)
	}

	return c.JSON(successResponse{Result: result})
}
