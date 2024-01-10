package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("ping")
	})

	app.Get("/mini-app-versions", func(c *fiber.Ctx) error {
		return c.JSON(map[string]map[string]interface{}{
			"MiniAppOne": {
				"name":    "MiniAppOne",
				"code":    "one",
				"version": "1.0.5",
			},
			"MiniAppTwo": {
				"name":    "MiniAppTwo",
				"code":    "two",
				"version": "1",
			},
		})
	})

	app.Static("/bundle", "./bundle")

	app.Listen(":8080")
}
