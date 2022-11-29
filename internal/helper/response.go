package helper

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
}

func BuildResponse(ctx *fiber.Ctx, status bool, message string, err, data interface{}, code int) error {
	return ctx.Status(code).JSON(&Response{
		Status:  status,
		Message: message,
		Errors:  err,
		Data:    data,
	})
}
