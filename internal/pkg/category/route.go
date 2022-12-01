package category

import (
	"tugas_akhir/internal/infrastructure/container"
	"tugas_akhir/internal/utils"

	"github.com/gofiber/fiber/v2"

	categorycontroller "tugas_akhir/internal/pkg/category/controller"

	categoryrepository "tugas_akhir/internal/pkg/category/repository"
	userrepository "tugas_akhir/internal/pkg/user/repository"

	categoryusecase "tugas_akhir/internal/pkg/category/usecase"
)

func CategoryRoute(r fiber.Router, containerConf *container.Container) {
	repo := categoryrepository.NewCategoryRepository(containerConf.Mysqldb)
	userRepo := userrepository.NewUserRepository(containerConf.Mysqldb)
	usecase := categoryusecase.NewCategoryUseCase(repo, userRepo)
	controller := categorycontroller.NewCategoryController(usecase)

	r.Get("", controller.GetAllCategories)
	r.Get("/:id_category", utils.MiddlewareAuth, controller.GetCategoryByID)
	r.Post("", utils.MiddlewareAuth, controller.CreateCategory)
	r.Put("/:id_category", utils.MiddlewareAuth, controller.UpdateCategoryByID)
	r.Delete("/:id_category", utils.MiddlewareAuth, controller.DeleteCategoryByID)
}
