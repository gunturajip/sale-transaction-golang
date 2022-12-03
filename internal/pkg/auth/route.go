package auth

import (
	"tugas_akhir/internal/infrastructure/container"

	"github.com/gofiber/fiber/v2"

	authcontroller "tugas_akhir/internal/pkg/auth/controller"
	authrepository "tugas_akhir/internal/pkg/auth/repository"
	provincecityrepository "tugas_akhir/internal/pkg/provincecity/repository"
	tokorepository "tugas_akhir/internal/pkg/toko/repository"

	authusecase "tugas_akhir/internal/pkg/auth/usecase"
)

func AuthRoute(r fiber.Router, containerConf *container.Container) {
	repoToko := tokorepository.NewTokoRepository(containerConf.Mysqldb)
	repoProv := provincecityrepository.NewProviceCityRepository(containerConf.Mysqldb)
	repo := authrepository.NewAuthRepository(containerConf.Mysqldb)
	usecase := authusecase.NewAuthUseCase(repo, repoToko, repoProv, containerConf.Mysqldb)
	controller := authcontroller.NewAuthUseCase(usecase)

	r.Post("/register", controller.Register)
	r.Post("/login", controller.Login)
}
