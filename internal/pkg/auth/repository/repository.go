package authrepository

import (
	"context"
	"tugas_akhir/internal/dao"

	"gorm.io/gorm"
)

type AuthRepository interface {
	LoginRepo(ctx context.Context, user dao.UserLogin) (data dao.User, err error)
	RegisterRepo(ctx context.Context, tx *gorm.DB, user dao.User) (id uint, err error)
}

type AuthRepositoryImpl struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) AuthRepository {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (ar *AuthRepositoryImpl) LoginRepo(ctx context.Context, user dao.UserLogin) (data dao.User, err error) {
	if err := ar.db.Debug().Preload("Toko").First(&data, "no_telp = ?", user.NoTelp).WithContext(ctx).Error; err != nil {
		return data, err
	}

	return data, nil
}

func (ar *AuthRepositoryImpl) RegisterRepo(ctx context.Context, tx *gorm.DB, user dao.User) (id uint, err error) {
	result := tx.Create(&user).WithContext(ctx)
	if result.Error != nil {
		return id, result.Error
	}

	return user.ID, nil
}
