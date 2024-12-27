package main

import "github.com/linuxfight/yandexCalcApi/cmd/application"

func main() {
	app := application.New()

	app.Start()
}
