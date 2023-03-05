package server

import (
	"ProjectAnalysis/pkg/api"
	"sort"

	"github.com/gofiber/fiber/v2"
)

func (as *AverageServer) Median(ctx *fiber.Ctx) (err error) {
	body := &api.CreateMedian{}

	ctx.BodyParser(body)

	value1 := body.Value1

	sort.Slice(value1, func(i, j int) bool {
		return value1[i] < value1[j]
	})

	lenValue := len(value1)

	if lenValue%2 == 0 {

		middleware := lenValue / 2
		middlewareBack := middleware - 1
		mediana1 := value1[middleware]
		mediana2 := value1[middlewareBack]

		result := (mediana1 + mediana2) / 2

		return ctx.JSON(result)
	}

	middleware := lenValue / 2

	result := value1[middleware]

	return ctx.JSON(result)

}
