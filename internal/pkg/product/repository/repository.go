package productrepository

import (
	"context"
	"errors"
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

	GetProductsBySliceID(ctx context.Context, tx *gorm.DB, sliceid []uint) (res []dao.ProductTotalPrice, err error)
	CreateProductLog(ctx context.Context, tx *gorm.DB, data []dao.LogProduct) (res []dao.LogProduct, err error)
	UpdateProductStock(ctx context.Context, tx *gorm.DB, productid uint, qty int) (res string, err error)
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

	db := cr.db
	if filter.NamaProduk != "" {
		db = db.Where("nama_produk like ?", "%"+filter.NamaProduk+"%")
	}

	if filter.MaxHarga > 0 && filter.MinHarga > 0 {
		fmt.Println("filter.MaxHarga", filter.MaxHarga)
		fmt.Println("filter.MinHarga", filter.MinHarga)
		db = db.Where("harga_konsumen > ? AND harga_konsumen < ?", filter.MinHarga, filter.MaxHarga)
	}

	if filter.TokoID != 0 {
		db = db.Where("toko_id = ?", filter.TokoID)
	}

	if filter.CategoryID != 0 {
		db = db.Where("category_id = ?", filter.CategoryID)
	}

	if err := db.Debug().Preload("Category").Preload("Toko").Preload("Photos").WithContext(ctx).Limit(filter.Limit).Offset(offset).Find(&res).Error; err != nil {
		return res, err
	}
	return res, nil
}

func (cr *ProductRepositoryImpl) GetProductByID(ctx context.Context, productid string) (res dao.Product, err error) {
	if err := cr.db.Debug().Preload("Category").Preload("Toko").Preload("Photos").First(&res, productid).WithContext(ctx).Error; err != nil {
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

// /
func (cr *ProductRepositoryImpl) GetProductsBySliceID(ctx context.Context, tx *gorm.DB, sliceid []uint) (res []dao.ProductTotalPrice, err error) {
	db := tx

	if len(sliceid) > 0 {
		db = db.Where("id IN ?", sliceid)
	}

	if err := db.Debug().Preload("Category").Preload("Toko").Preload("Photos").WithContext(ctx).Find(&res).Error; err != nil {
		tx.Rollback()
		return res, err
	}
	return res, nil
}
func (cr *ProductRepositoryImpl) CreateProductLog(ctx context.Context, tx *gorm.DB, data []dao.LogProduct) (res []dao.LogProduct, err error) {
	result := tx.Create(&data).WithContext(ctx)
	if result.Error != nil {
		return res, result.Error
	}

	return data, nil
}

func (cr *ProductRepositoryImpl) UpdateProductStock(ctx context.Context, tx *gorm.DB, productid uint, qty int) (res string, err error) {
	var dataProd dao.Product
	if err := tx.Debug().First(&dataProd, productid).WithContext(ctx).Error; err != nil {
		return "", err
	}

	if dataProd.Stok < qty {
		log.Println("exceed stok product id ", dataProd.ID)
		return "", errors.New(fmt.Sprint("exceed stok product id", dataProd.ID))
	}

	if err := tx.Debug().Model(&dataProd).Where("id = ?", productid).Update("stok", dataProd.Stok-qty).WithContext(ctx).Error; err != nil {
		log.Println("Update product failed", err, " product id : ", productid)
		return "", errors.New("Update product failed")
	}

	log.Println("sukes update product stok, id product : ", dataProd.ID)
	return "sukes update product stok", nil
}
