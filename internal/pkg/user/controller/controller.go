package usercontroller

import (
	"tugas_akhir/internal/helper"
	userdto "tugas_akhir/internal/pkg/user/dto"
	userusecase "tugas_akhir/internal/pkg/user/usecase"

	"github.com/gofiber/fiber/v2"
)

type UserUseCase interface {
	GetMyProfile(ctx *fiber.Ctx) error
	UpdateMyProfile(ctx *fiber.Ctx) error
}

type UserUseCaseImpl struct {
	userusecase userusecase.UserUseCase
}

func NewUserUseCase(userusecase userusecase.UserUseCase) UserUseCase {
	return &UserUseCaseImpl{
		userusecase: userusecase,
	}
}

func (uc *UserUseCaseImpl) GetMyProfile(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)

	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	res, err := uc.userusecase.MyProfile(c, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *UserUseCaseImpl) UpdateMyProfile(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)

	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	data := new(userdto.UserUpdateReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.userusecase.UpdateMyProfile(c, useridStr, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}
