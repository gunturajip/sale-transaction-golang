package authusecase

import (
	"context"
	"log"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	authdto "tugas_akhir/internal/pkg/auth/dto"
	authrepository "tugas_akhir/internal/pkg/auth/repository"

	"tugas_akhir/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type AuthUseCase interface {
	LoginUC(ctx context.Context, data authdto.LoginRequest)
	RegisterUC(ctx context.Context, data authdto.RegisterRequest) (res string, err *helper.ErrorStruct)
}

type AuthUseCaseImpl struct {
	authrepository authrepository.AuthRepository
}

func NewAuthRepository(authrepository authrepository.AuthRepository) AuthUseCase {
	return &AuthUseCaseImpl{
		authrepository: authrepository,
	}
}

func (ar *AuthUseCaseImpl) LoginUC(ctx context.Context, data authdto.LoginRequest) {

}

func (ar *AuthUseCaseImpl) RegisterUC(ctx context.Context, data authdto.RegisterRequest) (res string, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return "", &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	hashPass, errHash := utils.HashPassword(data.KataSandi)
	if errHash != nil {
		log.Println(errHash)
		return "", &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errHash,
		}
	}

	tanggalLahir, errShortDate := utils.ShortDateFromString(data.TanggalLahir)
	if errShortDate != nil {
		log.Println(errShortDate)
		return "", &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errShortDate,
		}
	}

	dataRegis := dao.User{
		Nama:         data.Nama,
		KataSandi:    hashPass,
		NoTelp:       data.NoTelp,
		TanggalLahir: tanggalLahir,
		Tentang:      data.Tentang,
		Perkerjaan:   data.Perkerjaan,
		Email:        data.Email,
		IDProvinsi:   data.IDProvinsi,
		IDKota:       data.IDKota,
		// Alamat:       []dao.Alamat{},
	}

	res, errRepo := ar.authrepository.RegisterRepo(ctx, dataRegis)
	if helper.MysqlCheckErrDuplicateEntry(errRepo) {
		log.Println(errRepo)
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return "", &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	log.Println("Register Succedd")
	return res, err
}
