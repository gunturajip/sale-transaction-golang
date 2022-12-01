package alamatusecase

import (
	"context"
	"errors"
	"log"
	"strconv"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	alamatdto "tugas_akhir/internal/pkg/alamat/dto"
	alamatrepository "tugas_akhir/internal/pkg/alamat/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AlamatUseCase interface {
	GetAllAlamat(ctx context.Context, judul_alamat, userid string) (res []alamatdto.AlamatResp, err *helper.ErrorStruct)
	GetAlamatByID(ctx context.Context, alamatid string) (res alamatdto.AlamatResp, err *helper.ErrorStruct)
	CreateAlamat(ctx context.Context, userid string, data alamatdto.AlamatReqCreate) (res uint, err *helper.ErrorStruct)
	UpdateAlamatByID(ctx context.Context, alamatid, userid string, data alamatdto.AlamatReqUpdate) (res string, err *helper.ErrorStruct)
	DeleteAlamatByID(ctx context.Context, alamatid, userid string) (res string, err *helper.ErrorStruct)
}

type AlamatUseCaseImpl struct {
	alamatrepository alamatrepository.AlamatRepository
}

func NewAlamatUseCase(alamatrepository alamatrepository.AlamatRepository) AlamatUseCase {
	return &AlamatUseCaseImpl{
		alamatrepository: alamatrepository,
	}

}

func (alc *AlamatUseCaseImpl) GetAllAlamat(ctx context.Context, judul_alamat, userid string) (res []alamatdto.AlamatResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.alamatrepository.GetAllAlamat(ctx, judul_alamat, userid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Alamat"),
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
		res = append(res, alamatdto.AlamatResp{
			ID:           v.ID,
			JudulAlamat:  v.JudulAlamat,
			NamaPenerima: v.NamaPenerima,
			NoTelp:       v.NoTelp,
			DetailAlamat: v.DetailAlamat,
		})
	}

	return res, nil
}
func (alc *AlamatUseCaseImpl) GetAlamatByID(ctx context.Context, alamatid string) (res alamatdto.AlamatResp, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.alamatrepository.GetAlamatByID(ctx, alamatid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Alamat"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = alamatdto.AlamatResp{
		ID:           resRepo.ID,
		JudulAlamat:  resRepo.JudulAlamat,
		NamaPenerima: resRepo.NamaPenerima,
		NoTelp:       resRepo.NoTelp,
		DetailAlamat: resRepo.DetailAlamat,
	}

	return res, nil
}
func (alc *AlamatUseCaseImpl) CreateAlamat(ctx context.Context, userid string, data alamatdto.AlamatReqCreate) (res uint, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	userridParse, errParse := strconv.ParseUint(userid, 10, 64)
	if errParse != nil {
		log.Println(errParse)
		return res, &helper.ErrorStruct{
			Err:  errParse,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.alamatrepository.CreateAlamat(ctx, dao.Alamat{
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		NoTelp:       data.NamaPenerima,
		DetailAlamat: data.DetailAlamat,
		UserID:       uint(userridParse),
	})
	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *AlamatUseCaseImpl) UpdateAlamatByID(ctx context.Context, alamatid, userid string, data alamatdto.AlamatReqUpdate) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	userridParse, errParse := strconv.ParseUint(userid, 10, 64)
	if errParse != nil {
		log.Println(errParse)
		return res, &helper.ErrorStruct{
			Err:  errParse,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := alc.alamatrepository.UpdateAlamatByID(ctx, alamatid, userid, dao.Alamat{
		JudulAlamat:  data.JudulAlamat,
		NamaPenerima: data.NamaPenerima,
		NoTelp:       data.NamaPenerima,
		DetailAlamat: data.DetailAlamat,
		UserID:       uint(userridParse),
	})

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
func (alc *AlamatUseCaseImpl) DeleteAlamatByID(ctx context.Context, alamatid, userid string) (res string, err *helper.ErrorStruct) {
	resRepo, errRepo := alc.alamatrepository.DeleteAlamatByID(ctx, alamatid, userid)
	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}
