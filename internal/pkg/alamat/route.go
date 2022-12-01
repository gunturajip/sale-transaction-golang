package alamat

import (
	"tugas_akhir/internal/infrastructure/container"
	"tugas_akhir/internal/utils"

	"github.com/gofiber/fiber/v2"

	alamatcontroller "tugas_akhir/internal/pkg/alamat/controller"

	alamatrepository "tugas_akhir/internal/pkg/alamat/repository"

	alamatusecase "tugas_akhir/internal/pkg/alamat/usecase"
)

func AlamatRoute(r fiber.Router, containerConf *container.Container) {
	repo := alamatrepository.NewAlamatRepository(containerConf.Mysqldb)
	usecase := alamatusecase.NewAlamatUseCase(repo)
	controller := alamatcontroller.NewAlamatController(usecase)

	r.Get("", utils.MiddlewareAuth, controller.GetAllAlamat)
	r.Get("/:id_alamat", utils.MiddlewareAuth, controller.GetAlamatByID)
	r.Post("", utils.MiddlewareAuth, controller.CreateAlamat)
	r.Put("/:id_alamat", utils.MiddlewareAuth, controller.UpdateAlamatByID)
	r.Delete("/:id_alamat", utils.MiddlewareAuth, controller.DeleteAlamatByID)
}
