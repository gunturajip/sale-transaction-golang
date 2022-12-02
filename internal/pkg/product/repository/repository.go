package productrepository

import (
	"context"
	"fmt"
	"log"
	"tugas_akhir/internal/dao"
	productdto "tugas_akhir/internal/pkg/product/dto"
	"tugas_akhir/internal/utils"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetAllProducts(ctx context.Context, filter productdto.ProductFilter) (res []dao.Product, err error)
	GetProductByID(ctx context.Context, productid string) (res dao.Product, err error)
	CreateProduct(ctx context.Context, data dao.Product) (res uint, err error)
	UpdateProductByID(ctx context.Context, productid string, data dao.Product, photos []dao.ProductPhotos) (res string, err error)
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

	if err := cr.db.Preload("Photos").Find(&res).WithContext(ctx).Limit(filter.Limit).Offset(offset).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (cr *ProductRepositoryImpl) GetProductByID(ctx context.Context, productid string) (res dao.Product, err error) {
	if err := cr.db.Preload("Photos").First(&res, productid).WithContext(ctx).Error; err != nil {
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

func (cr *ProductRepositoryImpl) UpdateProductByID(ctx context.Context, productid string, data dao.Product, photos []dao.ProductPhotos) (res string, err error) {
	var dataProduct dao.Product
	var productPhotos dao.ProductPhotos
	if err = cr.db.Preload("Photos").Where("id = ? AND toko_id = ? ", productid, data.TokoID).First(&dataProduct).WithContext(ctx).Error; err != nil {
		return "Update product failed", gorm.ErrRecordNotFound
	}

	// TRANSACTION, BECAUSE ISSUE
	// https://stackoverflow.com/questions/71129151/golang-gorm-error-on-insert-or-update-on-table-violates-foreign-key-contraint
	// CANT FIX !!!
	tx := cr.db.Begin()
	if err := tx.Debug().Where("id = ? AND toko_id = ?", productid, data.TokoID).Updates(&data).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return "Update product failed", err
	}

	if len(photos) > 0 {
		if err := tx.Debug().Where("product_id = ?", productid).Delete(&productPhotos).WithContext(ctx).Error; err != nil {
			tx.Rollback()
			log.Println(err)
			return "Delete old photo product failed", err
		}

		if err := tx.Debug().Save(&photos).WithContext(ctx).Error; err != nil {
			tx.Rollback()
			log.Println(err)
			return "Insert new photo product failed", err
		}

		// DELETE OLD PHOTO
		// if len(Photos) > 0 {
		log.Println("update photo product")
		if len(dataProduct.Photos) > 0 {
			for _, v := range dataProduct.Photos {
				err := utils.HandleRemoveFile(v.Url)
				if err != nil {
					tx.Rollback()
					return "Update product failed (delete photo failed)", fmt.Errorf("Update product failed : %s", err.Error())
				}
			}
			fmt.Println("photo deleted")
		}
		// }
	}

	tx.Commit()
	return res, nil
}

func (cr *ProductRepositoryImpl) DeleteProductByID(ctx context.Context, productid string, tokoid uint) (res string, err error) {
	var dataProduct dao.Product
	var productPhotos dao.ProductPhotos

	tx := cr.db.Begin()

	fmt.Println("GET DATA PRODUCT")
	if err = tx.Preload("Photos").Where("id = ? AND toko_id = ?", productid, tokoid).First(&dataProduct).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		return "Delete product failed", gorm.ErrRecordNotFound
	}

	fmt.Println("DELETE PRODUCT PHOTO")
	if err := tx.Debug().Where("product_id = ?", productid).Delete(&productPhotos).WithContext(ctx).Error; err != nil {
		tx.Rollback()
		log.Println(err)
		return "Delete old photo product failed", err
	}

	fmt.Println("DELETE PRODUCT")
	if err := tx.Model(dataProduct).Delete(&dataProduct).Error; err != nil {
		return "Delete product failed", err
	}

	fmt.Println("DELETE OLD PHOTO")
	if len(dataProduct.Photos) > 0 {
		log.Println("update photo product")
		for _, v := range dataProduct.Photos {
			err := utils.HandleRemoveFile(v.Url)
			if err != nil {
				tx.Rollback()
				return "Delete product failed (delete photo failed)", fmt.Errorf("Delete product failed : %s", err.Error())
			}
		}
	}

	tx.Commit()

	return res, nil
}
