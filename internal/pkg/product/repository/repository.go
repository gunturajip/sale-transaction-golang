package productrepository

import (
	"context"
	"tugas_akhir/internal/dao"
	productdto "tugas_akhir/internal/pkg/product/dto"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context, filter productdto.ProductFilter) (res []dao.Product, err error)
	GetProductByID(ctx context.Context, productid string) (res dao.Product, err error)
	CreateProduct(ctx context.Context, data dao.Product) (res uint, err error)
	UpdateProductByID(ctx context.Context, productid string, data dao.Product) (res string, err error)
	DeleteProductByID(ctx context.Context, productid string, tokoid uint) (res string, err error)
}

type ProductRepositoryImpl struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &ProductRepositoryImpl{
		db: db,
	}
}

// TODO
func (cr *ProductRepositoryImpl) GetAllProducts(ctx context.Context, filter productdto.ProductFilter) (res []dao.Product, err error) {
	offset := (filter.Page - 1) * filter.Limit

	if err := cr.db.Find(&res).WithContext(ctx).Limit(filter.Limit).Offset(offset).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (cr *ProductRepositoryImpl) GetProductByID(ctx context.Context, productid string) (res dao.Product, err error) {
	if err := cr.db.Debug().First(&res, productid).WithContext(ctx).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (cr *ProductRepositoryImpl) CreateProduct(ctx context.Context, data dao.Product) (res uint, err error) {
	result := cr.db.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data.ID, nil
}

func (cr *ProductRepositoryImpl) UpdateProductByID(ctx context.Context, productid string, data dao.Product) (res string, err error) {
	var dataProduct dao.Product
	if err = cr.db.Where("id = ? AND toko_id ", productid, data.TokoID).First(&dataProduct).WithContext(ctx).Error; err != nil {
		return "Update product failed", gorm.ErrRecordNotFound
	}

	if err := cr.db.Model(dataProduct).Updates(&data).Where("id = ? AND toko_id ", productid, data.TokoID).Error; err != nil {
		return "Update product failed", err
	}

	return res, nil
}

func (cr *ProductRepositoryImpl) DeleteProductByID(ctx context.Context, productid string, tokoid uint) (res string, err error) {
	var dataProduct dao.Product
	if err = cr.db.Where("id = ? AND toko_id ", productid, tokoid).First(&dataProduct).WithContext(ctx).Error; err != nil {
		return "Delete product failed", gorm.ErrRecordNotFound
	}

	if err := cr.db.Model(dataProduct).Delete(&dataProduct).Error; err != nil {
		return "Delete product failed", err
	}

	return res, nil
}
