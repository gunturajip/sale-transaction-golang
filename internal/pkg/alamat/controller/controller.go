package alamatcontroller

import (
	"log"
	"tugas_akhir/internal/helper"
	alamatdto "tugas_akhir/internal/pkg/alamat/dto"
	alamatusecase "tugas_akhir/internal/pkg/alamat/usecase"

	"github.com/gofiber/fiber/v2"
)

type AlamatController interface {
	GetAllAlamat(ctx *fiber.Ctx) error
	GetAlamatByID(ctx *fiber.Ctx) error
	CreateAlamat(ctx *fiber.Ctx) error
	UpdateAlamatByID(ctx *fiber.Ctx) error
	DeleteAlamatByID(ctx *fiber.Ctx) error
}

type AlamatControllerImpl struct {
	alamatusecase alamatusecase.AlamatUseCase
}

func NewAlamatController(alamatusecase alamatusecase.AlamatUseCase) AlamatController {
	return &AlamatControllerImpl{
		alamatusecase: alamatusecase,
	}
}

func (uc *AlamatControllerImpl) GetAllAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	filter := new(alamatdto.AlamatFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.alamatusecase.GetAllAlamat(c, filter.JudulAlamat, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *AlamatControllerImpl) GetAlamatByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	alamatid := ctx.Params("id_alamat")
	if alamatid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	res, err := uc.alamatusecase.GetAlamatByID(c, alamatid)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *AlamatControllerImpl) CreateAlamat(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	data := new(alamatdto.AlamatReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.alamatusecase.CreateAlamat(c, useridStr, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, "", res, fiber.StatusOK)
}

func (uc *AlamatControllerImpl) UpdateAlamatByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	alamatid := ctx.Params("id_alamat")
	if alamatid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	data := new(alamatdto.AlamatReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.alamatusecase.UpdateAlamatByID(c, alamatid, useridStr, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *AlamatControllerImpl) DeleteAlamatByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	alamatid := ctx.Params("id_alamat")
	if alamatid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	res, err := uc.alamatusecase.DeleteAlamatByID(c, alamatid, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}
