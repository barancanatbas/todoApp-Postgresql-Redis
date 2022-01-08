package main

import (
	"todo/internal/config"
	"todo/internal/router"
	"todo/repository"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	config.Init()
	repository.Set()
	router.Init(e)

	e.Logger.Fatal(e.Start(":8080"))
}
