package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/linuxfight/yandexCalcApi/internal/services/logger"
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

	app.Post("/api/v1/calculate", func(c fiber.Ctx) error {

		return c.JSON(fiber.Map{
			"Status": "ok",
		})
	})

	err := app.Listen(":8080")
	if err != nil {
		logger.Log.Panicf("listening error: %s", err.Error())
		return
	}
}
