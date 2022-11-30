package authcontroller

import (
	"tugas_akhir/internal/helper"
	authdto "tugas_akhir/internal/pkg/auth/dto"
	authusecase "tugas_akhir/internal/pkg/auth/usecase"

	"github.com/gofiber/fiber/v2"
)

type AuthUseCase interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type AuthUseCaseImpl struct {
	authusecase authusecase.AuthUseCase
}

func NewAuthUseCase(authusecase authusecase.AuthUseCase) AuthUseCase {
	return &AuthUseCaseImpl{
		authusecase: authusecase,
	}
}

func (ar *AuthUseCaseImpl) Login(ctx *fiber.Ctx) error {
	c := ctx.Context()

	user := new(authdto.LoginRequest)
	if err := ctx.BodyParser(user); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := ar.authusecase.LoginUC(c, *user)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, "", res, fiber.StatusOK)
}

func (ar *AuthUseCaseImpl) Register(ctx *fiber.Ctx) error {
	c := ctx.Context()

	user := new(authdto.RegisterRequest)
	if err := ctx.BodyParser(user); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := ar.authusecase.RegisterUC(c, *user)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, "", res, fiber.StatusCreated)
}
