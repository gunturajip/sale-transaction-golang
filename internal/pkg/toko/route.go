package toko

import (
	"tugas_akhir/internal/infrastructure/container"
	"tugas_akhir/internal/utils"

	"github.com/gofiber/fiber/v2"

	tokocontroller "tugas_akhir/internal/pkg/toko/controller"
	tokorepository "tugas_akhir/internal/pkg/toko/repository"
	tokousecase "tugas_akhir/internal/pkg/toko/usecase"
)

func TokoRoute(r fiber.Router, containerConf *container.Container) {
	repo := tokorepository.NewTokoRepository(containerConf.Mysqldb)
	usecase := tokousecase.NewTokoUseCase(repo)
	controller := tokocontroller.NewTokoUseCase(usecase)

	r.Get("/my", utils.MiddlewareAuth, controller.MyToko)
	r.Get("/:id_toko", controller.FindByID)
	r.Get("/", controller.GetAll)
	r.Put("/:id_toko", utils.MiddlewareAuth, utils.HandleSingleFile, controller.UpdateByID)
}
