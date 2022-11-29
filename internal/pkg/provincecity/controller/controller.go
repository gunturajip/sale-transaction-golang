package provincecitycontroller

import (
	"tugas_akhir/internal/helper"
	provincecityusecase "tugas_akhir/internal/pkg/provincecity/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProviceCityController interface {
	GetListProvince(ctx *fiber.Ctx) error
	GetListCity(ctx *fiber.Ctx) error
	GetDetailProvince(ctx *fiber.Ctx) error
	GetDetailCity(ctx *fiber.Ctx) error
}

type ProviceCityControllerImpl struct {
	provincecityusecase provincecityusecase.ProviceCityUseCase
}

func NewProviceCityController(provincecityusecase provincecityusecase.ProviceCityUseCase) ProviceCityController {
	return &ProviceCityControllerImpl{
		provincecityusecase: provincecityusecase,
	}
}

func (pc *ProviceCityControllerImpl) GetListProvince(ctx *fiber.Ctx) error {
	res, err := pc.provincecityusecase.GetListProvince()
	if err.Err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (pc *ProviceCityControllerImpl) GetListCity(ctx *fiber.Ctx) error {
	provID := ctx.Params("prov_id")

	if provID == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "ERROR NO CITY PARAMS", nil, fiber.StatusBadRequest)
	}

	res, err := pc.provincecityusecase.GetListCity(provID)
	if err.Err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (pc *ProviceCityControllerImpl) GetDetailProvince(ctx *fiber.Ctx) error {
	provID := ctx.Params("prov_id")
	if provID == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "ERROR NO PROVINCE PARAMS", nil, fiber.StatusBadRequest)
	}

	res, err := pc.provincecityusecase.GetDetailProvince(provID)
	if err.Err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (pc *ProviceCityControllerImpl) GetDetailCity(ctx *fiber.Ctx) error {
	cityID := ctx.Params("city_id")

	if cityID == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "ERROR NO CITY PARAMS", nil, fiber.StatusBadRequest)
	}

	res, err := pc.provincecityusecase.GetDetailCity(cityID)
	if err.Err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}
