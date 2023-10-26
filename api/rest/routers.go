package rest

import "github.com/gofiber/fiber/v2"

func routers(app *fiber.App) {
	app.Get("/health", func(c *fiber.Ctx) error { return c.SendString("OK") })

	v1 := app.Group("/v1")
	{
		message := v1.Group("/send")
		{
			message.Post("/:type", sendHandler)
		}
		v1.Get("/count", countHandler)
	}
}
