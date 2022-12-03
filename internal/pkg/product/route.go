package product

import (
	"tugas_akhir/internal/infrastructure/container"
	"tugas_akhir/internal/utils"

	"github.com/gofiber/fiber/v2"

	productcontroller "tugas_akhir/internal/pkg/product/controller"

	productrepository "tugas_akhir/internal/pkg/product/repository"

	tokorepository "tugas_akhir/internal/pkg/toko/repository"

	productusecase "tugas_akhir/internal/pkg/product/usecase"
)

func ProductRoute(r fiber.Router, containerConf *container.Container) {
	repo := productrepository.NewProductRepository(containerConf.Mysqldb)
	tokoRepo := tokorepository.NewTokoRepository(containerConf.Mysqldb)
	usecase := productusecase.NewProductUseCase(repo, tokoRepo)
	controller := productcontroller.NewProductController(usecase)

	r.Get("", controller.GetAllProducts)
	r.Get("/:id_product", controller.GetProductByID)
	r.Post("", utils.MiddlewareAuth, utils.HandleMultiplePartFile, controller.CreateProduct)
	r.Put("/:id_product", utils.MiddlewareAuth, utils.HandleMultiplePartFile, controller.UpdateProductByID)
	r.Delete("/:id_product", utils.MiddlewareAuth, controller.DeleteProductByID)
}
