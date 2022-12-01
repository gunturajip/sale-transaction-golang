package categoryrepository

import (
	"context"
	"tugas_akhir/internal/dao"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	GetAllCategories(ctx context.Context) (res []dao.Category, err error)
	GetCategoryByID(ctx context.Context, categoryid string) (res dao.Category, err error)
	CreateCategory(ctx context.Context, data dao.Category) (res uint, err error)
	UpdateCategoryByID(ctx context.Context, categoryid string, data dao.Category) (res string, err error)
	DeleteCategoryByID(ctx context.Context, categoryid string) (res string, err error)
}

type CategoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		db: db,
	}
}
func (cr *CategoryRepositoryImpl) GetAllCategories(ctx context.Context) (res []dao.Category, err error) {
	if err := cr.db.Find(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (cr *CategoryRepositoryImpl) GetCategoryByID(ctx context.Context, categoryid string) (res dao.Category, err error) {
	if err := cr.db.First(&res, categoryid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (cr *CategoryRepositoryImpl) CreateCategory(ctx context.Context, data dao.Category) (res uint, err error) {
	result := cr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (cr *CategoryRepositoryImpl) UpdateCategoryByID(ctx context.Context, categoryid string, data dao.Category) (res string, err error) {
	var dataCategory dao.Category
	if err = cr.db.Where("id = ? ", categoryid).First(&dataCategory).WithContext(ctx).Error; err != nil {
		return "Update category failed", gorm.ErrRecordNotFound
	}

	if err := cr.db.Model(dataCategory).Updates(&data).Where("id = ? ", categoryid).Error; err != nil {
		return "Update category failed", err
	}

	return res, nil
}

func (cr *CategoryRepositoryImpl) DeleteCategoryByID(ctx context.Context, categoryid string) (res string, err error) {
	var dataCategory dao.Category
	if err = cr.db.Where("id = ? ", categoryid).First(&dataCategory).WithContext(ctx).Error; err != nil {
		return "Delete category failed", gorm.ErrRecordNotFound
	}

	if err := cr.db.Model(dataCategory).Delete(&dataCategory).Error; err != nil {
		return "Delete category failed", err
	}

	return res, nil
}
