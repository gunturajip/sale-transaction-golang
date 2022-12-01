package alamatrepository

import (
	"context"
	"tugas_akhir/internal/dao"

	"gorm.io/gorm"
)

type AlamatRepository interface {
	GetAllAlamat(ctx context.Context, judul_alamat, userid string) (res []dao.Alamat, err error)
	GetAlamatByID(ctx context.Context, alamatid string) (res dao.Alamat, err error)
	CreateAlamat(ctx context.Context, data dao.Alamat) (res uint, err error)
	UpdateAlamatByID(ctx context.Context, alamatid, userid string, data dao.Alamat) (res string, err error)
	DeleteAlamatByID(ctx context.Context, alamatid, userid string) (res string, err error)
}

type AlamatRepositoryImpl struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) AlamatRepository {
	return &AlamatRepositoryImpl{
		db: db,
	}
}
func (alr *AlamatRepositoryImpl) GetAllAlamat(ctx context.Context, judul_alamat, userid string) (res []dao.Alamat, err error) {
	db := alr.db

	if judul_alamat != "" {
		db = alr.db.Where("judul_alamat LIKE ?", "%"+judul_alamat+"%")
	}

	if err := db.Debug().Where("user_id = ?", userid).Find(&res).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *AlamatRepositoryImpl) GetAlamatByID(ctx context.Context, alamatid string) (res dao.Alamat, err error) {
	if err := alr.db.First(&res, alamatid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (alr *AlamatRepositoryImpl) CreateAlamat(ctx context.Context, data dao.Alamat) (res uint, err error) {
	result := alr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (alr *AlamatRepositoryImpl) UpdateAlamatByID(ctx context.Context, alamatid, userid string, data dao.Alamat) (res string, err error) {
	var dataAlamat dao.Alamat
	if err = alr.db.Where("id = ?  AND user_id = ?", alamatid, userid).First(&dataAlamat).WithContext(ctx).Error; err != nil {
		return "Update alamat failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataAlamat).Updates(&data).Where("id = ? ", alamatid).Error; err != nil {
		return "Update alamat failed", err
	}

	return res, nil
}

func (alr *AlamatRepositoryImpl) DeleteAlamatByID(ctx context.Context, alamatid, userid string) (res string, err error) {
	var dataAlamat dao.Alamat
	if err = alr.db.Where("id = ?  AND user_id = ?", alamatid, userid).First(&dataAlamat).WithContext(ctx).Error; err != nil {
		return "Delete alamat failed", gorm.ErrRecordNotFound
	}

	if err := alr.db.Model(dataAlamat).Delete(&dataAlamat).Error; err != nil {
		return "Delete alamat failed", err
	}

	return res, nil
}
