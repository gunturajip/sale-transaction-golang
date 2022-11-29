package provincecity

import (
	"tugas_akhir/internal/infrastructure/container"

	provincecitycontroller "tugas_akhir/internal/pkg/provincecity/controller"
	provincecityrepository "tugas_akhir/internal/pkg/provincecity/repository"
	provincecityusecase "tugas_akhir/internal/pkg/provincecity/usecase"

	"github.com/gofiber/fiber/v2"
)

func ProvinceCityRoute(r fiber.Router, containerConf *container.Container) {
	repo := provincecityrepository.NewProviceCityRepository(containerConf.Mysqldb)
	usecase := provincecityusecase.NewProviceCityRepository(repo)
	controller := provincecitycontroller.NewProviceCityController(usecase)

	r.Get("/listprovincies", controller.GetListProvince)
	r.Get("/listcities/:prov_id", controller.GetListCity)
	r.Get("/detailprovince/:prov_id", controller.GetDetailProvince)
	r.Get("/detailcity/:city_id", controller.GetDetailCity)
}
