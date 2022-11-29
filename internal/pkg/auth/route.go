package auth

import "github.com/gofiber/fiber/v2"

func AuthRoute(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
