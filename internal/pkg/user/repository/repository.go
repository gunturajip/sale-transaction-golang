package userrepository

import (
	"context"
	"tugas_akhir/internal/dao"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetMyProfileRepo(ctx context.Context, userid string) (res dao.User, err error)
	UpdateMyProfileRepo(ctx context.Context, userid string, data dao.User) (res string, err error)
	IsAdminRepo(ctx context.Context, userid string) (res bool, err error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (ur *UserRepositoryImpl) GetMyProfileRepo(ctx context.Context, userid string) (res dao.User, err error) {
	if err := ur.db.First(&res, userid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (ur *UserRepositoryImpl) UpdateMyProfileRepo(ctx context.Context, userid string, data dao.User) (res string, err error) {
	var dataUser dao.User
	if err = ur.db.Where("id = ? ", userid).First(&dataUser).WithContext(ctx).Error; err != nil {
		return "Update user failed", gorm.ErrRecordNotFound
	}

	if err := ur.db.Model(dataUser).Debug().Updates(&data).Where("id = ? ", userid).Error; err != nil {
		return "Update user failed", err
	}

	return res, nil
}

func (ur *UserRepositoryImpl) IsAdminRepo(ctx context.Context, userid string) (res bool, err error) {
	var dataUser dao.User
	if err = ur.db.Where("id = ? ", userid).First(&dataUser).WithContext(ctx).Error; err != nil {
		return false, gorm.ErrRecordNotFound
	}

	if dataUser.IsAdmin {
		return true, nil
	} else {
		return false, nil
	}

}
