package tokousecase

import (
	"context"
	"errors"
	"log"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	tokodto "tugas_akhir/internal/pkg/toko/dto"
	tokorepository "tugas_akhir/internal/pkg/toko/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TokoUseCase interface {
	MyToko(ctx context.Context, UserID string) (res tokodto.TokoResp, err *helper.ErrorStruct)
	FindByID(ctx context.Context, IDToko string) (res tokodto.TokoResp, err *helper.ErrorStruct)
	GetAll(ctx context.Context, filter tokodto.TokoFilterRequest) (res tokodto.TokoPagination, err *helper.ErrorStruct)
	UpdateByID(ctx context.Context, UserID string, IDToko string, data tokodto.TokoUpdateReq) (res string, err *helper.ErrorStruct)
	// CreateToko(ctx context.Context, tx *gorm.DB, data dao.Toko) (id uint,err *helper.ErrorStruct)
}

type TokoUseCaseImpl struct {
	tokorepository tokorepository.TokoRepository
}

func NewTokoUseCase(tokorepository tokorepository.TokoRepository) TokoUseCase {
	return &TokoUseCaseImpl{
		tokorepository: tokorepository,
	}

}

func (ar *TokoUseCaseImpl) MyToko(ctx context.Context, UserID string) (res tokodto.TokoResp, err *helper.ErrorStruct) {
	resRepo, errRepo := ar.tokorepository.FindByUserID(ctx, UserID)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("Toko tidak ditemukan"),
		}
	}

	res = tokodto.TokoResp{
		ID:       resRepo.ID,
		NamaToko: resRepo.NamaToko,
		UrlFoto:  resRepo.UrlFoto,
		UserID:   resRepo.UserID,
	}

	log.Println("Toko Find My Toko Succeed")
	return res, nil
}

func (ar *TokoUseCaseImpl) FindByID(ctx context.Context, IDToko string) (res tokodto.TokoResp, err *helper.ErrorStruct) {
	resRepo, errRepo := ar.tokorepository.FindByID(ctx, IDToko)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("Toko tidak ditemukan"),
		}
	}

	res = tokodto.TokoResp{
		ID:       resRepo.ID,
		NamaToko: resRepo.NamaToko,
		UrlFoto:  resRepo.UrlFoto,
		// UserID:   resRepo.UserID,
	}

	log.Println("Toko FindBy ID Succeed")
	return res, nil
}

func (ar *TokoUseCaseImpl) GetAll(ctx context.Context, filter tokodto.TokoFilterRequest) (res tokodto.TokoPagination, err *helper.ErrorStruct) {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Page == 0 {
		filter.Page = 1
	}

	resRepo, errRepo := ar.tokorepository.GetAll(ctx, filter)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("Toko tidak ditemukan"),
		}
	}

	for _, v := range resRepo {
		res.Data = append(res.Data, tokodto.TokoResp{
			ID:       v.ID,
			NamaToko: v.NamaToko,
			UrlFoto:  v.UrlFoto,
		})
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	res.Limit = filter.Limit
	res.Page = filter.Page

	return res, nil
}
func (ar *TokoUseCaseImpl) UpdateByID(ctx context.Context, UserID string, IDToko string, data tokodto.TokoUpdateReq) (res string, err *helper.ErrorStruct) {
	dataToko := dao.Toko{
		NamaToko: data.NamaToko,
		UrlFoto:  data.UrlFoto,
		UserID:   data.UserID,
	}
	res, errRepo := ar.tokorepository.UpdateByID(ctx, UserID, IDToko, dataToko)
	if errRepo != nil {
		log.Println(errRepo)
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return res, nil
}
