package server

import (
	"ProjectAnalysis/pkg/api"

	"github.com/gofiber/fiber/v2"
)

func (as *AverageServer) Average(ctx *fiber.Ctx) (err error) {

	body := &api.CreateAverage{}

	ctx.BodyParser(body)

	average := (body.Value1 + body.Value2) / 2

	return ctx.JSON(average)
}
