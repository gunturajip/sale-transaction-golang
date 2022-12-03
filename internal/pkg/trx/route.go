package trx

import (
	"tugas_akhir/internal/infrastructure/container"
	"tugas_akhir/internal/utils"

	"github.com/gofiber/fiber/v2"

	trxcontroller "tugas_akhir/internal/pkg/trx/controller"

	trxrepository "tugas_akhir/internal/pkg/trx/repository"

	alamtrepository "tugas_akhir/internal/pkg/alamat/repository"
	productrepository "tugas_akhir/internal/pkg/product/repository"

	trxusecase "tugas_akhir/internal/pkg/trx/usecase"
)

func TrxRoute(r fiber.Router, containerConf *container.Container) {
	repo := trxrepository.NewTrxRepository(containerConf.Mysqldb)
	productRepo := productrepository.NewProductRepository(containerConf.Mysqldb)
	alamattRepo := alamtrepository.NewAlamatRepository(containerConf.Mysqldb)
	usecase := trxusecase.NewTrxUseCase(repo, productRepo, alamattRepo, containerConf.Mysqldb)
	controller := trxcontroller.NewTrxController(usecase)

	r.Get("", utils.MiddlewareAuth, controller.GetAllTrxs)
	r.Get("/:id_trx", utils.MiddlewareAuth, controller.GetTrxByID)
	r.Post("", utils.MiddlewareAuth, controller.CreateTrx)
}
