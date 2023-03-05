package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type AverageServer struct {
	app *fiber.App
}

func Run(as *AverageServer) {

	as.app = fiber.New(fiber.Config{
		AppName: "Analysis",
	})

	as.app.Use(cors.New())

	as.app.Post("/average", as.Average)

	as.app.Post("/weighted-average", as.WeightedAverage)

	as.app.Post("/median", as.Mediana)

	as.app.Listen(":3000")
}
