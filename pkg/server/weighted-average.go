package server

import (
	"ProjectAnalysis/pkg/api"

	"github.com/gofiber/fiber/v2"
)

func (as *AverageServer) WeightedAverage(ctx *fiber.Ctx) (err error) {

	body := &api.CreateWeightedAverage{}

	ctx.BodyParser(body)

	average := (body.Value1*body.Peso1 + body.Value2*body.Peso2) / 2

	return ctx.JSON(average)
}
