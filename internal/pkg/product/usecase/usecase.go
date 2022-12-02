package productusecase

import (
	"context"
	"errors"
	"log"
	"strings"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	productdto "tugas_akhir/internal/pkg/product/dto"
	productrepository "tugas_akhir/internal/pkg/product/repository"
	tokorepository "tugas_akhir/internal/pkg/toko/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ProductUseCase interface {
	GetAllProducts(ctx context.Context, filter productdto.ProductFilter) (res []productdto.ProductResp, err *helper.ErrorStruct)
	GetProductByID(ctx context.Context, productid string) (res productdto.ProductResp, err *helper.ErrorStruct)
	CreateProduct(ctx context.Context, userid string, data productdto.ProductReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateProductByID(ctx context.Context, productid, userid string, data productdto.ProductReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteProductByID(ctx context.Context, productid, userid string) (res string, err *helper.ErrorStruct)
}

type ProductUseCaseImpl struct {
	productrepository productrepository.ProductRepository
	tokorepository    tokorepository.TokoRepository
}

func NewProductUseCase(productrepository productrepository.ProductRepository, tokorepository tokorepository.TokoRepository) ProductUseCase {
	return &ProductUseCaseImpl{
		productrepository: productrepository,
		tokorepository:    tokorepository,
	}

}

func (pu *ProductUseCaseImpl) GetAllProducts(ctx context.Context, filter productdto.ProductFilter) (res []productdto.ProductResp, err *helper.ErrorStruct) {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Page == 0 {
		filter.Page = 1
	}

	resRepo, errRepo := pu.productrepository.GetAllProducts(ctx, filter)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Product"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	for _, v := range resRepo {
		result := productdto.ProductResp{
			ID:            v.ID,
			NamaProduk:    v.NamaProduk,
			Slug:          v.Slug,
			HargaReseler:  v.HargaReseler,
			HargaKonsumen: v.HargaKonsumen,
			Stok:          v.Stok,
			Deskripsi:     v.Deskripsi,
			TokoID:        v.TokoID,
			CategoryID:    v.CategoryID,
		}

		for _, photo := range v.Photos {
			result.Photos = append(result.Photos, productdto.ProductPhotos{
				ID:        photo.ID,
				ProductID: photo.ProductID,
				Url:       photo.Url,
			})
		}

		res = append(res, result)
	}

	return res, nil
}
func (pu *ProductUseCaseImpl) GetProductByID(ctx context.Context, productid string) (res productdto.ProductResp, err *helper.ErrorStruct) {
	resRepo, errRepo := pu.productrepository.GetProductByID(ctx, productid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Product"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = productdto.ProductResp{
		ID:            resRepo.ID,
		NamaProduk:    resRepo.NamaProduk,
		Slug:          resRepo.Slug,
		HargaReseler:  resRepo.HargaReseler,
		HargaKonsumen: resRepo.HargaKonsumen,
		Stok:          resRepo.Stok,
		Deskripsi:     resRepo.Deskripsi,
		TokoID:        resRepo.TokoID,
		CategoryID:    resRepo.CategoryID,
	}

	for _, photo := range resRepo.Photos {
		res.Photos = append(res.Photos, productdto.ProductPhotos{
			ID:        photo.ID,
			ProductID: photo.ProductID,
			Url:       photo.Url,
		})
	}

	return res, nil
}
func (pu *ProductUseCaseImpl) CreateProduct(ctx context.Context, userid string, data productdto.ProductReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// GET ID TOKO FROM THIS USER
	tokoid, err := pu.GetMyTokoID(ctx, userid)
	if err != nil {
		return res, err
	}

	// CREATE SLUG PRODUCT
	arrStr := strings.Split(data.NamaProduk, " ")
	slug := strings.Join(arrStr, "-")
	slug = strings.ToLower(slug)

	//
	dataProduct := dao.Product{
		NamaProduk:    data.NamaProduk,
		Slug:          slug,
		HargaReseler:  data.HargaReseler,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		CategoryID:    data.CategoryID,
		TokoID:        tokoid,
		// Photos:        []dao.ProductPhotos{},
	}

	resRepo, errRepo := pu.productrepository.CreateProduct(ctx, dataProduct)

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (pu *ProductUseCaseImpl) UpdateProductByID(ctx context.Context, productid, userid string, data productdto.ProductReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// GET ID TOKO FROM THIS USER
	tokoid, err := pu.GetMyTokoID(ctx, userid)
	if err != nil {
		return res, err
	}

	// CREATE SLUG PRODUCT
	var slug string
	if data.NamaProduk != "" {
		arrStr := strings.Split(data.NamaProduk, " ")
		slug = strings.Join(arrStr, "-")
		slug = strings.ToLower(slug)
	}

	//
	dataProduct := dao.Product{
		NamaProduk:    data.NamaProduk,
		Slug:          slug,
		HargaReseler:  data.HargaReseler,
		HargaKonsumen: data.HargaKonsumen,
		Stok:          data.Stok,
		Deskripsi:     data.Deskripsi,
		TokoID:        tokoid,
		CategoryID:    data.CategoryID,
		// Photos:        []dao.ProductPhotos{},
	}

	resRepo, errRepo := pu.productrepository.UpdateProductByID(ctx, productid, dataProduct)

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (pu *ProductUseCaseImpl) DeleteProductByID(ctx context.Context, productid, userid string) (res string, err *helper.ErrorStruct) {
	// GET ID TOKO FROM THIS USER
	tokoid, err := pu.GetMyTokoID(ctx, userid)
	if err != nil {
		return res, err
	}

	resRepo, errRepo := pu.productrepository.DeleteProductByID(ctx, productid, tokoid)
	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (pu *ProductUseCaseImpl) GetMyTokoID(ctx context.Context, userid string) (res uint, err *helper.ErrorStruct) {
	resToko, errToko := pu.tokorepository.FindByUserID(ctx, userid)
	if errToko != nil {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("Unauthorized"),
		}
	}

	return resToko.ID, nil
}
