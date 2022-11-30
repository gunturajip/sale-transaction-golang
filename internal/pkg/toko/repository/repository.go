package tokorepository

import (
	"context"
	"tugas_akhir/internal/dao"
	tokodto "tugas_akhir/internal/pkg/toko/dto"

	"gorm.io/gorm"
)

type TokoRepository interface {
	FindByIDOrUserID(ctx context.Context, data map[string]interface{}) (res dao.Toko, err error)
	GetAll(ctx context.Context, filter tokodto.TokoFilterRequest) (res []dao.Toko, err error)
	UpdateByID(ctx context.Context, IDToko string) (res string, err error)
	CreateToko(ctx context.Context, tx *gorm.DB, data dao.Toko) (id uint, err error)
}

type TokoRepositoryImpl struct {
	db *gorm.DB
}

func NewTokoRepository(db *gorm.DB) TokoRepository {
	return &TokoRepositoryImpl{
		db: db,
	}
}

func (tr *TokoRepositoryImpl) FindByIDOrUserID(ctx context.Context, data map[string]interface{}) (res dao.Toko, err error) {
	return res, err
}
func (tr *TokoRepositoryImpl) GetAll(ctx context.Context, filter tokodto.TokoFilterRequest) (res []dao.Toko, err error) {
	return res, err
}
func (tr *TokoRepositoryImpl) UpdateByID(ctx context.Context, IDToko string) (res string, err error) {
	return res, err
}
func (tr *TokoRepositoryImpl) CreateToko(ctx context.Context, tx *gorm.DB, data dao.Toko) (id uint, err error) {
	res := tx.Create(&data).WithContext(ctx)
	if res.Error != nil {
		return 0, err
	}
	return data.ID, res.Error
}
