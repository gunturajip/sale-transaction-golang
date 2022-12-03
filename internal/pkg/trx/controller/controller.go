package trxcontroller

import (
	"log"
	"tugas_akhir/internal/helper"
	trxdto "tugas_akhir/internal/pkg/trx/dto"
	trxusecase "tugas_akhir/internal/pkg/trx/usecase"

	"github.com/gofiber/fiber/v2"
)

type TrxController interface {
	GetAllTrxs(ctx *fiber.Ctx) error
	GetTrxByID(ctx *fiber.Ctx) error
	CreateTrx(ctx *fiber.Ctx) error
}

type TrxControllerImpl struct {
	trxusecase trxusecase.TrxUseCase
}

func NewTrxController(trxusecase trxusecase.TrxUseCase) TrxController {
	return &TrxControllerImpl{
		trxusecase: trxusecase,
	}
}

func (uc *TrxControllerImpl) GetAllTrxs(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	filter := new(trxdto.TrxFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.trxusecase.GetAllTrxs(c, useridStr, *filter)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *TrxControllerImpl) GetTrxByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	trxid := ctx.Params("id_trx")
	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	if trxid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	res, err := uc.trxusecase.GetTrxByID(c, trxid, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *TrxControllerImpl) CreateTrx(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	data := new(trxdto.TrxReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.trxusecase.CreateTrx(c, useridStr, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, "", res, fiber.StatusOK)
}
