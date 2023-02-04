package router

import (
	"github.com/gofiber/fiber/v2"
	"goTestFiber/database"
	"goTestFiber/internal/usecase"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/sum", GetSum)
}
func GetSum(c *fiber.Ctx) error {
	conDB := database.ConnectDB()
	response := usecase.SumFromDB(conDB)
	return c.JSON(response)
}
