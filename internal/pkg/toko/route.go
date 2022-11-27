package toko

import "github.com/gofiber/fiber/v2"

func UserRoute(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
