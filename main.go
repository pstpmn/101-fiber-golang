package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"data": "Hello World",
			"code": "0",
		})
	})

	app.Listen(":3000")
}
