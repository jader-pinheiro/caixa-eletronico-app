package ui

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jadermoura/caixa-eletronico-app/internal/infra/db/data"
	"github.com/jadermoura/caixa-eletronico-app/internal/infra/db/entity"
	"gorm.io/gorm"
)

type Success struct {
	Status string `json:"status"`
	Msg    string `json:"msg,omitempty"`
}

func Create(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		// Use `db` aqui

		var ac entity.Account
		if err := ctx.BodyParser(&ac); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"msg":    err.Error,
			})
		}

		id, err := data.Create(db, ctx.Context(), ac)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"msg":    err.Error,
			})
		}

		return ctx.JSON(fiber.Map{
			"message":  "Conta criada com sucesso!",
			"api_code": fiber.StatusCreated,
			"id":       id,
		})
	}
}

func GetBalance(db *gorm.DB) func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {

		number := ctx.Params("accountID")

		if number == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": fiber.StatusBadRequest,
				"msg":    "Número da conta não informado",
			})
		}
		// intVal, _ := strconv.Atoi(number)

		account, err := data.GetBalance(db, ctx.Context(), number)

		if err != nil {
			fmt.Printf("Erro ao obter saldo da conta: %v\n", err)
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"status": fiber.StatusNotFound,
				"msg":    err.Error(),
			})
		}

		return ctx.JSON(fiber.Map{
			"account": account,
			"status":  fiber.StatusOK,
			"msg":     "Saldo obtido com sucesso!",
		})
	}
}
