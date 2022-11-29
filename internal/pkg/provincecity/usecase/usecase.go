package provincecityusecase

import (
	"errors"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	provincecityrepository "tugas_akhir/internal/pkg/provincecity/repository"

	"github.com/gofiber/fiber/v2"
)

type ProviceCityUseCase interface {
	GetListProvince() (provList []dao.Province, err helper.ErrorStruct)
	GetListCity(provinceID string) (cityList []dao.City, err helper.ErrorStruct)
	GetDetailProvince(provinceID string) (provdetail dao.Province, err helper.ErrorStruct)
	GetDetailCity(cityID string) (citydetail dao.City, err helper.ErrorStruct)
}

type ProviceCityUseCaseImpl struct {
	provincecityrepository provincecityrepository.ProviceCityRepository
}

func NewProviceCityRepository(provincecityrepository provincecityrepository.ProviceCityRepository) ProviceCityUseCase {
	return &ProviceCityUseCaseImpl{
		provincecityrepository: provincecityrepository,
	}
}

func (pcu *ProviceCityUseCaseImpl) GetListProvince() (provList []dao.Province, err helper.ErrorStruct) {
	res, errRepo := pcu.provincecityrepository.GetListProvince()
	if errRepo != nil {
		return provList, helper.ErrorStruct{
			Err:  errRepo,
			Code: fiber.StatusInternalServerError,
		}
	}
	return res, err
}

func (pcu *ProviceCityUseCaseImpl) GetListCity(provinceID string) (cityList []dao.City, err helper.ErrorStruct) {
	res, errRepo := pcu.provincecityrepository.GetListCity(provinceID)
	if errRepo != nil {
		return cityList, helper.ErrorStruct{
			Err:  errors.New("data not found"),
			Code: fiber.StatusNotFound,
		}
	}
	return res, err
}
func (pcu *ProviceCityUseCaseImpl) GetDetailProvince(provinceID string) (provdetail dao.Province, err helper.ErrorStruct) {
	res, errRepo := pcu.provincecityrepository.GetDetailProvince(provinceID)
	if errRepo != nil {
		return provdetail, helper.ErrorStruct{
			Err:  errors.New("data not found"),
			Code: fiber.StatusNotFound,
		}
	}
	return res, err
}
func (pcu *ProviceCityUseCaseImpl) GetDetailCity(cityID string) (citydetail dao.City, err helper.ErrorStruct) {
	res, errRepo := pcu.provincecityrepository.GetDetailCity(cityID)
	if errRepo != nil {
		return citydetail, helper.ErrorStruct{
			Err:  errors.New("data not found"),
			Code: fiber.StatusNotFound,
		}
	}
	return res, err
}
