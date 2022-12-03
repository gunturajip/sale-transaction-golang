package trxrepository

import (
	"context"
	"tugas_akhir/internal/dao"
	trxdto "tugas_akhir/internal/pkg/trx/dto"

	"gorm.io/gorm"
)

type TrxRepository interface {
	GetAllTrxs(ctx context.Context, filter trxdto.TrxFilter) (res []dao.Trx, err error)
	GetTrxByID(ctx context.Context, trxid string) (res dao.Trx, err error)
	CreateTrx(ctx context.Context, tx *gorm.DB, data dao.Trx) (res uint, err error)
}

type TrxRepositoryImpl struct {
	db *gorm.DB
}

func NewTrxRepository(db *gorm.DB) TrxRepository {
	return &TrxRepositoryImpl{
		db: db,
	}
}

// TODO
func (tr *TrxRepositoryImpl) GetAllTrxs(ctx context.Context, filter trxdto.TrxFilter) (res []dao.Trx, err error) {
	offset := (filter.Page - 1) * filter.Limit

	if err := tr.db.Debug().Preload("User").Preload("Alamat").Preload("DetailTrx").WithContext(ctx).Limit(filter.Limit).Offset(offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (tr *TrxRepositoryImpl) GetTrxByID(ctx context.Context, trxid string) (res dao.Trx, err error) {
	if err := tr.db.Debug().Preload("User").Preload("Alamat").Preload("DetailTrx").First(&res, trxid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (tr *TrxRepositoryImpl) CreateTrx(ctx context.Context, tx *gorm.DB, data dao.Trx) (res uint, err error) {
	result := tx.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}
