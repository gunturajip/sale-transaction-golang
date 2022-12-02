package tokocontroller

import (
	"fmt"
	"log"
	"tugas_akhir/internal/helper"
	tokodto "tugas_akhir/internal/pkg/toko/dto"
	tokousecase "tugas_akhir/internal/pkg/toko/usecase"

	"github.com/gofiber/fiber/v2"
)

type TokoUseCase interface {
	MyToko(ctx *fiber.Ctx) error
	FindByID(ctx *fiber.Ctx) error
	GetAll(ctx *fiber.Ctx) error
	UpdateByID(ctx *fiber.Ctx) error
}

type TokoUseCaseImpl struct {
	tokousecase tokousecase.TokoUseCase
}

func NewTokoUseCase(tokousecase tokousecase.TokoUseCase) TokoUseCase {
	return &TokoUseCaseImpl{
		tokousecase: tokousecase,
	}
}

func (tu *TokoUseCaseImpl) MyToko(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)

	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	res, err := tu.tokousecase.MyToko(c, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (tu *TokoUseCaseImpl) FindByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	tokoid := ctx.Params("id_toko")

	if tokoid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusUnprocessableEntity)
	}

	res, err := tu.tokousecase.FindByID(c, tokoid)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (tu *TokoUseCaseImpl) GetAll(ctx *fiber.Ctx) error {
	c := ctx.Context()

	filter := new(tokodto.TokoFilterRequest)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := tu.tokousecase.GetAll(c, *filter)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (tu *TokoUseCaseImpl) UpdateByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	tokoid := ctx.Params("id_toko")
	userid := ctx.Locals("userid")

	if tokoid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusUnprocessableEntity)
	}

	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	data := new(tokodto.TokoUpdateReq)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	// mengambil context filename
	var filenameStr string
	filename := ctx.Locals("filename")
	if filename != nil {
		filenameStr = fmt.Sprintf("%v", filename)
		data.UrlFoto = filenameStr
	}

	res, err := tu.tokousecase.UpdateByID(c, useridStr, tokoid, *data)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDUPDATEDATA, "", res, fiber.StatusOK)
}
