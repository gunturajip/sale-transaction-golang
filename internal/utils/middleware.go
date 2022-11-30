package utils

import "github.com/gofiber/fiber/v2"

func MiddlewareAuth(ctx *fiber.Ctx) error {
	token := ctx.Get("token")
	if token == "" {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// _, err := utils.VerifyToken(token)
	claims, err := DecodeToken(token)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	ctx.Locals("userid", claims["id"])
	ctx.Locals("useremail", claims["email"])

	// Go to next middleware:
	return ctx.Next()
}
