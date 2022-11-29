package http

import (
	"tugas_akhir/internal/pkg/auth"
	"tugas_akhir/internal/pkg/provincecity"

	"tugas_akhir/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Route")
	})

	provcityAPI := api.Group("/provcity")
	provincecity.ProvinceCityRoute(provcityAPI, containerConf)

	authAPI := api.Group("/auth")
	auth.AuthRoute(authAPI)

	// userAPI := api.Group("/user")
	// user.UserRoute(userAPI)
	// alamatAPI := userAPI.Group("/alamat")

	// tokoAPI := api.Group("/toko")

	// productAPI := api.Group("/product")

	// categoryAPI := api.Group("/category")

	// authAPI := api.Group("/auth")

	// trxAPI := api.Group("/trx")

}
