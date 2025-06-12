package ui

import (
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
