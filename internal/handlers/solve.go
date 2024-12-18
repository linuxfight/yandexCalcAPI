package handlers

import (
	"github.com/gofiber/fiber/v3"
	"github.com/linuxfight/yandexCalcApi/pkg/calc"
)

type SolveRequest struct {
	Expression string `json:"expression"`
}

func SolveHandler(c fiber.Ctx) error {
	var body SolveRequest
	if err := c.Bind().JSON(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			ErrorResponse{Error: err.Error()},
		)
	}

	result, solveErr := calc.Solve(body.Expression)
	if solveErr != nil {

	}
}
