package userusecase

import (
	"context"
	"errors"
	"log"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	"tugas_akhir/internal/utils"

	provincecityrepository "tugas_akhir/internal/pkg/provincecity/repository"
	userdto "tugas_akhir/internal/pkg/user/dto"
	userrepository "tugas_akhir/internal/pkg/user/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserUseCase interface {
	MyProfile(ctx context.Context, UserID string) (res userdto.UserResp, err *helper.ErrorStruct)
	UpdateMyProfile(ctx context.Context, UserID string, data userdto.UserUpdateReq) (res string, err *helper.ErrorStruct)
}

type UserUseCaseImpl struct {
	userrepository         userrepository.UserRepository
	provincecityrepository provincecityrepository.ProviceCityRepository
}

func NewUserUseCase(userrepository userrepository.UserRepository, provincecityrepository provincecityrepository.ProviceCityRepository) UserUseCase {
	return &UserUseCaseImpl{
		userrepository:         userrepository,
		provincecityrepository: provincecityrepository,
	}

}

func (ar *UserUseCaseImpl) MyProfile(ctx context.Context, UserID string) (res userdto.UserResp, err *helper.ErrorStruct) {
	resRepo, errRepo := ar.userrepository.GetMyProfileRepo(ctx, UserID)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("Unauthorized"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	// GET DATA PROVINCE
	resProvince, errProvince := ar.provincecityrepository.GetDetailProvince(resRepo.IDProvinsi)
	if err != nil {
		log.Println("get province", errProvince)
	}

	// GET DATA CITY
	resCity, errCity := ar.provincecityrepository.GetDetailCity(resRepo.IDKota)
	if err != nil {
		log.Println("err get city", errCity)
	}

	res = userdto.UserResp{
		ID:           resRepo.ID,
		Nama:         resRepo.Nama,
		NoTelp:       resRepo.NoTelp,
		TanggalLahir: utils.ShortDateFromDate(resRepo.TanggalLahir),
		Tentang:      resRepo.Tentang,
		Perkerjaan:   resRepo.Perkerjaan,
		Email:        resRepo.Email,
		IDProvinsi: dao.Province{
			ID:   resProvince.ID,
			Name: resProvince.Name,
		},
		IDKota: dao.City{
			ID:         resCity.ID,
			ProvinceID: resCity.ProvinceID,
			Name:       resCity.Name,
		},
	}

	log.Println("Get My Profile Succeed")
	return res, nil
}

func (ar *UserUseCaseImpl) UpdateMyProfile(ctx context.Context, UserID string, data userdto.UserUpdateReq) (res string, err *helper.ErrorStruct) {
	tanggalLahir, errTgl := utils.ShortDateFromString(data.TanggalLahir)
	if errTgl != nil {
		log.Println(errTgl)
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errTgl,
		}
	}

	if data.KataSandi != "" {
		hashPass, errHash := utils.HashPassword(data.KataSandi)
		if errHash != nil {
			log.Println(errHash)
			return res, &helper.ErrorStruct{
				Code: fiber.StatusInternalServerError,
				Err:  errHash,
			}
		}

		data.KataSandi = hashPass
	}

	dataUser := dao.User{
		Nama:         data.Nama,
		KataSandi:    data.KataSandi,
		NoTelp:       data.NoTelp,
		TanggalLahir: tanggalLahir,
		Tentang:      data.Tentang,
		Perkerjaan:   data.Perkerjaan,
		Email:        data.Email,
		IDProvinsi:   data.IDProvinsi,
		IDKota:       data.IDKota,
	}

	resRepo, errRepo := ar.userrepository.UpdateMyProfileRepo(ctx, UserID, dataUser)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("Unauthorized"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return "", &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	log.Println("Update My User Succeed")
	return resRepo, nil
}
