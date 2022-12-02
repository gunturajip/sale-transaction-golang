package categorycontroller

import (
	"tugas_akhir/internal/helper"
	categorydto "tugas_akhir/internal/pkg/category/dto"
	categoryusecase "tugas_akhir/internal/pkg/category/usecase"

	"github.com/gofiber/fiber/v2"
)

type CategoryController interface {
	GetAllCategories(ctx *fiber.Ctx) error
	GetCategoryByID(ctx *fiber.Ctx) error
	CreateCategory(ctx *fiber.Ctx) error
	UpdateCategoryByID(ctx *fiber.Ctx) error
	DeleteCategoryByID(ctx *fiber.Ctx) error
}

type CategoryControllerImpl struct {
	categoryusecase categoryusecase.CategoryUseCase
}

func NewCategoryController(categoryusecase categoryusecase.CategoryUseCase) CategoryController {
	return &CategoryControllerImpl{
		categoryusecase: categoryusecase,
	}
}

func (uc *CategoryControllerImpl) GetAllCategories(ctx *fiber.Ctx) error {
	c := ctx.Context()

	res, err := uc.categoryusecase.GetAllCategories(c)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *CategoryControllerImpl) GetCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryid := ctx.Params("id_category")
	if categoryid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	res, err := uc.categoryusecase.GetCategoryByID(c, categoryid, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *CategoryControllerImpl) CreateCategory(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	data := new(categorydto.CategoryReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.categoryusecase.CreateCategory(c, useridStr, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, "", res, fiber.StatusOK)
}

func (uc *CategoryControllerImpl) UpdateCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryid := ctx.Params("id_category")
	if categoryid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	data := new(categorydto.CategoryReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.categoryusecase.UpdateCategoryByID(c, categoryid, useridStr, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *CategoryControllerImpl) DeleteCategoryByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	categoryid := ctx.Params("id_category")
	if categoryid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	res, err := uc.categoryusecase.DeleteCategoryByID(c, categoryid, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}
