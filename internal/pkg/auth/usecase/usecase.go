package authusecase

import (
	"context"
	"errors"
	"log"
	"strings"
	"time"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	authdto "tugas_akhir/internal/pkg/auth/dto"
	authrepository "tugas_akhir/internal/pkg/auth/repository"
	tokorepository "tugas_akhir/internal/pkg/toko/repository"

	"tugas_akhir/internal/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthUseCase interface {
	LoginUC(ctx context.Context, data authdto.LoginRequest) (res authdto.LoginResp, err *helper.ErrorStruct)
	RegisterUC(ctx context.Context, data authdto.RegisterRequest) (res string, err *helper.ErrorStruct)
}

type AuthUseCaseImpl struct {
	authrepository authrepository.AuthRepository
	tokorepository tokorepository.TokoRepository
	db             *gorm.DB
}

func NewAuthRepository(authrepository authrepository.AuthRepository, tokorepository tokorepository.TokoRepository, db *gorm.DB) AuthUseCase {
	return &AuthUseCaseImpl{
		authrepository: authrepository,
		tokorepository: tokorepository,
		db:             db,
	}

}

func (ar *AuthUseCaseImpl) LoginUC(ctx context.Context, data authdto.LoginRequest) (res authdto.LoginResp, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := ar.authrepository.LoginRepo(ctx, dao.UserLogin{KataSandi: data.Password, NoTelp: data.NoTelp})
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("No Telp atau kata sandi salah"),
		}
	}

	isValid := utils.CheckPasswordHash(data.Password, resRepo.KataSandi)
	if !isValid {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("No Telp atau kata sandi salah"),
		}
	}

	// GENERATE JWT TOKEN
	claims := jwt.MapClaims{}
	claims["id"] = resRepo.ID
	claims["email"] = resRepo.Email
	claims["exp"] = time.Now().Add(48 * time.Hour).Unix()

	token, errGenerateToken := utils.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = authdto.LoginResp{
		Nama:         resRepo.Nama,
		NoTelp:       resRepo.Nama,
		TanggalLahir: utils.ShortDateFromDate(resRepo.TanggalLahir),
		Tentang:      resRepo.Nama,
		Perkerjaan:   resRepo.Nama,
		Email:        resRepo.Nama,
		IDProvinsi:   resRepo.Nama,
		IDKota:       resRepo.Nama,
		Token:        token,
	}

	return res, err
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

	// TRANSACTION

	// CREATE USER
	tx := ar.db.Begin()
	resID, errRepo := ar.authrepository.RegisterRepo(ctx, tx, dataRegis)
	if helper.MysqlCheckErrDuplicateEntry(errRepo) {
		log.Println(errRepo)
		tx.Rollback()
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		tx.Rollback()
		return "", &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	// CREATE TOKO
	dataToko := dao.Toko{
		NamaToko: ar.createShopName(data.Nama),
		UserID:   resID,
	}
	_, errRepoToko := ar.tokorepository.CreateToko(ctx, tx, dataToko)
	if errRepoToko != nil {
		log.Println(errRepoToko)
		tx.Rollback()
		return "", &helper.ErrorStruct{
			Code: fiber.StatusInternalServerError,
			Err:  errRepo,
		}
	}

	tx.Commit()

	log.Println("Register Succedd")
	return "Register Succeed", err
}

func (ar *AuthUseCaseImpl) createShopName(nama string) string {
	arrStr := strings.Split(nama, " ")
	var newArrStr []string

	for _, v := range arrStr {
		if len(v) < 3 {
			newArrStr = append(newArrStr, v)
		} else {
			newArrStr = append(newArrStr, v[0:3])
		}
	}

	return strings.Join(newArrStr, "-")

}
