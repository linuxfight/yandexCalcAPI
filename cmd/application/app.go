package application

import (
	"github.com/gofiber/fiber/v3"
	"github.com/linuxfight/yandexCalcApi/internal/logger"
	"github.com/linuxfight/yandexCalcApi/internal/server"
	"time"
)

type Application struct {
	Http *fiber.App
}

func (a Application) Start() {
	err := a.Http.Listen(":8080")
	if err != nil {
		logger.Log.Panicf("listening error: %s", err.Error())
		return
	}
}

func New() *Application {
	localTime := time.Now()
	timezone, _ := localTime.Zone()

	app := fiber.New()
	logger.New(
		true,
		timezone,
	)

	logger.Log.Debug("logger initialized")
	logger.Log.Debugf("current timezone: %s", timezone)

	app.Use(server.Recovery())
	app.Post("/api/v1/calculate", server.SolveHandler)

	return &Application{Http: app}
}
