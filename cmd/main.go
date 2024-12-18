package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/linuxfight/yandexCalcApi/internal/handlers"
	"github.com/linuxfight/yandexCalcApi/internal/logger"
	"time"
)

func main() {
	localTime := time.Now()
	timezone, _ := localTime.Zone()

	app := fiber.New()
	logger.New(
		true,
		timezone,
	)

	logger.Log.Debug("logger initialized")
	logger.Log.Debugf("current timezone: %s", timezone)

	app.Post("/api/v1/calculate", handlers.SolveHandler)

	err := app.Listen(":8080")
	if err != nil {
		logger.Log.Panicf("listening error: %s", err.Error())
		return
	}
}
