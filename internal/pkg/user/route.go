package user

import (
	"tugas_akhir/internal/infrastructure/container"
	"tugas_akhir/internal/utils"

	"github.com/gofiber/fiber/v2"

	usercontroller "tugas_akhir/internal/pkg/user/controller"

	userrepository "tugas_akhir/internal/pkg/user/repository"

	userusecase "tugas_akhir/internal/pkg/user/usecase"
)

func UserRoute(r fiber.Router, containerConf *container.Container) {
	repo := userrepository.NewUserRepository(containerConf.Mysqldb)
	usecase := userusecase.NewUserUseCase(repo)
	controller := usercontroller.NewUserController(usecase)

	r.Get("", utils.MiddlewareAuth, controller.GetMyProfile)
	r.Put("", utils.MiddlewareAuth, controller.UpdateMyProfile)
}
