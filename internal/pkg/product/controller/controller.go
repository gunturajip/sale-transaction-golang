package productcontroller

import (
	"fmt"
	"log"
	"tugas_akhir/internal/helper"
	productdto "tugas_akhir/internal/pkg/product/dto"
	productusecase "tugas_akhir/internal/pkg/product/usecase"

	"github.com/gofiber/fiber/v2"
)

type ProductController interface {
	GetAllProducts(ctx *fiber.Ctx) error
	GetProductByID(ctx *fiber.Ctx) error
	CreateProduct(ctx *fiber.Ctx) error
	UpdateProductByID(ctx *fiber.Ctx) error
	DeleteProductByID(ctx *fiber.Ctx) error
}

type ProductControllerImpl struct {
	productusecase productusecase.ProductUseCase
}

func NewProductController(productusecase productusecase.ProductUseCase) ProductController {
	return &ProductControllerImpl{
		productusecase: productusecase,
	}
}

func (uc *ProductControllerImpl) GetAllProducts(ctx *fiber.Ctx) error {
	c := ctx.Context()
	filter := new(productdto.ProductFilter)
	if err := ctx.QueryParser(filter); err != nil {
		log.Println(err)
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.productusecase.GetAllProducts(c, *filter)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *ProductControllerImpl) GetProductByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	productid := ctx.Params("id_product")

	if productid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	res, err := uc.productusecase.GetProductByID(c, productid)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *ProductControllerImpl) CreateProduct(ctx *fiber.Ctx) error {
	c := ctx.Context()
	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	// VALIDATION REQUIRED IMAGE
	// mengambil context filename
	var filenamesData []string
	filenames := ctx.Locals("filenames")
	if filenames == nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, "photos product is required", nil, fiber.StatusBadRequest)
	} else {
		filenamesData = filenames.([]string)
	}

	if len(filenamesData) < 1 {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, "photos product is required", nil, fiber.StatusBadRequest)
	}

	data := new(productdto.ProductReqCreate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.productusecase.CreateProduct(c, useridStr, *data, filenamesData)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDPOSTDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDPOSTDATA, "", res, fiber.StatusOK)
}

func (uc *ProductControllerImpl) UpdateProductByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	productid := ctx.Params("id_product")
	if productid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	// VALIDATION REQUIRED IMAGE
	// mengambil context filename
	var filenamesData []string
	filenames := ctx.Locals("filenames")
	if filenames == nil {
		log.Println("no image uploaded")
	} else {
		filenamesData = filenames.([]string)
	}

	fmt.Println("filenamesData", filenamesData)

	data := new(productdto.ProductReqUpdate)
	if err := ctx.BodyParser(data); err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, err.Error(), nil, fiber.StatusBadRequest)
	}

	res, err := uc.productusecase.UpdateProductByID(c, productid, useridStr, *data, filenamesData)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}

func (uc *ProductControllerImpl) DeleteProductByID(ctx *fiber.Ctx) error {
	c := ctx.Context()
	productid := ctx.Params("id_product")
	if productid == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, "PARAM REQUIRED", nil, fiber.StatusBadRequest)
	}

	userid := ctx.Locals("userid")
	useridStr := userid.(string)
	if useridStr == "" {
		return helper.BuildResponse(ctx, false, helper.FAILEDUPDATEDATA, "UNAUTHORIZED", nil, fiber.StatusUnauthorized)
	}

	res, err := uc.productusecase.DeleteProductByID(c, productid, useridStr)
	if err != nil {
		return helper.BuildResponse(ctx, false, helper.FAILEDGETDATA, err.Err.Error(), nil, err.Code)
	}

	return helper.BuildResponse(ctx, true, helper.SUCCEEDGETDATA, "", res, fiber.StatusOK)
}
