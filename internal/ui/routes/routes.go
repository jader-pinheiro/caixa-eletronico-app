package routes

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jadermoura/caixa-eletronico-app/internal/ui"
	"gorm.io/gorm"
)

func GetRoutes(conn *gorm.DB) {

	app := fiber.New()

	app.Post("/account/create", ui.Create(conn))
	//app.Post("/account/deposit", ui.Deposit(conn))
	//app.Post("/account/withdraw", ui.Withdraw(conn))
	app.Get("/account/balance/:accountID", ui.GetBalance(conn))
	// app.Get("/account/statement/:accountID", ui.Statement(conn))
	// app.Get("/account/transfer/:accountID", ui.Transfer(conn))

	app.Listen(":" + os.Getenv("PORT"))

	log.Fatal(app.Listen(":3000"))

}
