package http

import (
	"fmt"
	"tugas_akhir/internal/helper"
	"tugas_akhir/internal/pkg/alamat"
	"tugas_akhir/internal/pkg/auth"
	"tugas_akhir/internal/pkg/category"
	"tugas_akhir/internal/pkg/product"
	"tugas_akhir/internal/pkg/provincecity"
	"tugas_akhir/internal/pkg/toko"
	"tugas_akhir/internal/pkg/trx"
	"tugas_akhir/internal/pkg/user"

	"tugas_akhir/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"
)

func HTTPRouteInit(r *fiber.App, containerConf *container.Container) {
	api := r.Group("/api/v1") // /api

	r.Static("/public", helper.ProjectRootPath+"/public/img")
	fmt.Println("ProjectRootPath", helper.ProjectRootPath)

	api.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! Route")
	})

	provcityAPI := api.Group("/provcity")
	provincecity.ProvinceCityRoute(provcityAPI, containerConf)

	authAPI := api.Group("/auth")
	auth.AuthRoute(authAPI, containerConf)

	userAPI := api.Group("/user")
	user.UserRoute(userAPI, containerConf)

	alamatAPI := userAPI.Group("/alamat")
	alamat.AlamatRoute(alamatAPI, containerConf)

	tokoAPI := api.Group("/toko")
	toko.TokoRoute(tokoAPI, containerConf)

	productAPI := api.Group("/product")
	product.ProductRoute(productAPI, containerConf)

	categoryAPI := api.Group("/category")
	category.CategoryRoute(categoryAPI, containerConf)

	trxAPI := api.Group("/trx")
	trx.TrxRoute(trxAPI, containerConf)

}
