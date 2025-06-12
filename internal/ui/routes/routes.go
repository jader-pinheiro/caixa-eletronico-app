package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/jadermoura/caixa-eletronico-app/internal/ui"
	"gorm.io/gorm"
)

func GetRoutes(conn *gorm.DB) {

	app := fiber.New()

	app.Post("/account/create", ui.Create(conn))

	log.Fatal(app.Listen(":3000"))

}
