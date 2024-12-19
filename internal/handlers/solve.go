package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/linuxfight/yandexCalcApi/internal/logger"
	"github.com/linuxfight/yandexCalcApi/pkg/calc"
)

func SolveHandler(c fiber.Ctx) error {
	var body solveRequest
	if err := c.Bind().JSON(&body); err != nil || body.Expression == "" {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			&errorResponse{Error: "invalid json"},
		)
	}

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
	result, solveErr := calc.Solve(body.Expression)
	if solveErr != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(
			&errorResponse{Error: solveErr.Error()},
		)
	}

	return c.JSON(successResponse{Result: result})
}
