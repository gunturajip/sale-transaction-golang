package categoryusecase

import (
	"context"
	"errors"
	"log"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	categorydto "tugas_akhir/internal/pkg/category/dto"
	categoryrepository "tugas_akhir/internal/pkg/category/repository"
	userrepository "tugas_akhir/internal/pkg/user/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryUseCase interface {
	GetAllCategories(ctx context.Context) (res []categorydto.CategoryResp, err *helper.ErrorStruct)
	GetCategoryByID(ctx context.Context, categoryid, userid string) (res categorydto.CategoryResp, err *helper.ErrorStruct)
	CreateCategory(ctx context.Context, userid string, data categorydto.CategoryReq) (res uint, err *helper.ErrorStruct)
	UpdateCategoryByID(ctx context.Context, categoryid, userid string, data categorydto.CategoryReq) (res string, err *helper.ErrorStruct)
	DeleteCategoryByID(ctx context.Context, categoryid, userid string) (res string, err *helper.ErrorStruct)
}

type CategoryUseCaseImpl struct {
	categoryrepository categoryrepository.CategoryRepository
	userrepository     userrepository.UserRepository
}

func NewCategoryUseCase(categoryrepository categoryrepository.CategoryRepository, userrepository userrepository.UserRepository) CategoryUseCase {
	return &CategoryUseCaseImpl{
		categoryrepository: categoryrepository,
		userrepository:     userrepository,
	}

}

func (cu *CategoryUseCaseImpl) GetAllCategories(ctx context.Context) (res []categorydto.CategoryResp, err *helper.ErrorStruct) {
	resRepo, errRepo := cu.categoryrepository.GetAllCategories(ctx)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Category"),
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
		res = append(res, categorydto.CategoryResp{
			ID:           v.ID,
			NamaKategori: v.NamaKategori,
		})
	}

	return res, nil
}
func (cu *CategoryUseCaseImpl) GetCategoryByID(ctx context.Context, categoryid, userid string) (res categorydto.CategoryResp, err *helper.ErrorStruct) {
	isAdmin := cu.IsAdmin(ctx, userid)
	if !isAdmin {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("Unauthorized"),
		}
	}

	resRepo, errRepo := cu.categoryrepository.GetCategoryByID(ctx, categoryid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Category"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	res = categorydto.CategoryResp{
		ID:           resRepo.ID,
		NamaKategori: resRepo.NamaKategori,
	}

	return res, nil
}
func (cu *CategoryUseCaseImpl) CreateCategory(ctx context.Context, userid string, data categorydto.CategoryReq) (res uint, err *helper.ErrorStruct) {
	isAdmin := cu.IsAdmin(ctx, userid)
	if !isAdmin {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("Unauthorized"),
		}
	}

	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := cu.categoryrepository.CreateCategory(ctx, dao.Category{
		NamaKategori: data.NamaKategori,
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
func (cu *CategoryUseCaseImpl) UpdateCategoryByID(ctx context.Context, categoryid, userid string, data categorydto.CategoryReq) (res string, err *helper.ErrorStruct) {
	isAdmin := cu.IsAdmin(ctx, userid)
	if !isAdmin {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("Unauthorized"),
		}
	}

	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	resRepo, errRepo := cu.categoryrepository.UpdateCategoryByID(ctx, categoryid, dao.Category{
		NamaKategori: data.NamaKategori,
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
func (cu *CategoryUseCaseImpl) DeleteCategoryByID(ctx context.Context, categoryid, userid string) (res string, err *helper.ErrorStruct) {
	isAdmin := cu.IsAdmin(ctx, userid)
	if !isAdmin {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("Unauthorized"),
		}
	}

	resRepo, errRepo := cu.categoryrepository.DeleteCategoryByID(ctx, categoryid)
	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	return resRepo, nil
}

func (cu *CategoryUseCaseImpl) IsAdmin(ctx context.Context, userid string) bool {
	isAdmin, errIsAdmin := cu.userrepository.IsAdminRepo(ctx, userid)
	if !isAdmin || errIsAdmin != nil {
		log.Println("errIsAdmin : ", errIsAdmin)
		return false
	}

	return true
}
