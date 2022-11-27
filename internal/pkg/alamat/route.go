package alamat

import "github.com/gofiber/fiber/v2"

func AlamatRoute(r fiber.Router) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
