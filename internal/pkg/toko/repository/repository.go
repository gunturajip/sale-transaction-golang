package tokorepository

import (
	"context"
	"fmt"
	"log"
	"tugas_akhir/internal/dao"
	tokodto "tugas_akhir/internal/pkg/toko/dto"
	"tugas_akhir/internal/utils"

	"gorm.io/gorm"
)

type TokoRepository interface {
	FindByUserID(ctx context.Context, UserID string) (res dao.Toko, err error)
	FindByID(ctx context.Context, IDToko string) (res dao.Toko, err error)
	GetAll(ctx context.Context, filter tokodto.TokoFilterRequest) (res []dao.Toko, err error)
	UpdateByID(ctx context.Context, UserID string, IDToko string, data dao.Toko) (res string, err error)
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

func (tr *TokoRepositoryImpl) FindByUserID(ctx context.Context, UserID string) (res dao.Toko, err error) {
	if err := tr.db.First(&res, "user_id = ?", UserID).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (tr *TokoRepositoryImpl) FindByID(ctx context.Context, IDToko string) (res dao.Toko, err error) {
	if err := tr.db.First(&res, IDToko).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (tr *TokoRepositoryImpl) GetAll(ctx context.Context, filter tokodto.TokoFilterRequest) (res []dao.Toko, err error) {
	fmt.Printf("filter %#v \n\n", filter)
	var db = tr.db
	if filter.Name != "" {
		db = db.Where("nama_toko LIKE ?", "%"+filter.Name+"%")
	}

	offset := (filter.Page - 1) * filter.Limit

	if err := db.Debug().Limit(filter.Limit).Offset(offset).Find(&res).WithContext(ctx).Error; err != nil {
		return nil, err
	}

	return res, nil
}
func (tr *TokoRepositoryImpl) UpdateByID(ctx context.Context, UserID string, IDToko string, data dao.Toko) (res string, err error) {
	var dataToko dao.Toko
	if err = tr.db.Where("id = ? AND user_id = ?", IDToko, UserID).First(&dataToko).WithContext(ctx).Error; err != nil {
		return "Update toko failed", gorm.ErrRecordNotFound
	}

	if dataToko.UrlFoto != data.UrlFoto && dataToko.UrlFoto != "" && data.UrlFoto != "" {
		log.Println("update photo")
		err := utils.HandleRemoveFile(dataToko.UrlFoto)
		if err != nil {
			return "Update toko failed (delete photo failed)", fmt.Errorf("Update toko failed : %s", err.Error())
		}
	}

	if err := tr.db.Model(dataToko).Updates(&data).Where("id = ? AND user_id = ?", IDToko, UserID).Error; err != nil {
		return "Update toko failed", err
	}

	return "Update toko succeed", nil
}
func (tr *TokoRepositoryImpl) CreateToko(ctx context.Context, tx *gorm.DB, data dao.Toko) (id uint, err error) {
	res := tx.Create(&data).WithContext(ctx)
	if res.Error != nil {
		return 0, err
	}
	return data.ID, res.Error
}
