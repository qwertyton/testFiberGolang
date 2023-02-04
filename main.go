package main

import (
	"github.com/gofiber/fiber/v2"
	"goTestFiber/config"
	"goTestFiber/router"
)

func main() {
	config.Init()
	app := fiber.New()
	router.SetupRoutes(app)
	app.Listen(":3000")
}
