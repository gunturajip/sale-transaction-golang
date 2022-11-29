package auth

import (
	"tugas_akhir/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	authcontroller "tugas_akhir/internal/pkg/auth/controller"
	authrepository "tugas_akhir/internal/pkg/auth/repository"
	authusecase "tugas_akhir/internal/pkg/auth/usecase"
)

func AuthRoute(r fiber.Router, containerConf *container.Container) {
	repo := authrepository.NewAuthRepository(containerConf.Mysqldb)
	usecase := authusecase.NewAuthRepository(repo)
	controller := authcontroller.NewAuthRepository(usecase)

	r.Post("/register", controller.Register)
}
