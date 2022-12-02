package dao

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	NamaProduk    string          `json:"nama_produk"`
	Slug          string          `json:"slug"`
	HargaReseler  int             `json:"harga_reseler"`
	HargaKonsumen int             `json:"harga_konsumen"`
	Stok          int             `json:"stok"`
	Deskripsi     string          `json:"deskripsi" gorm:"type:text"`
	TokoID        uint            `json:"toko_id" gorm:"not null"`
	Toko          Toko            `json:"toko"` // foreign key
	CategoryID    uint            `json:"category_id" gorm:"not null"`
	Category      Category        `json:"category"` // foreign key
	Photos        []ProductPhotos `json:"photos"`
}

type LogProduct struct {
	gorm.Model
	NamaProduk    string          `json:"nama_produk"`
	Slug          string          `json:"slug"`
	HargaReseler  int             `json:"harga_reseler"`
	HargaKonsumen int             `json:"harga_konsumen"`
	Deskripsi     string          `json:"deskripsi" gorm:"type:text"`
	TokoID        uint            `json:"toko_id" gorm:"not null"`
	Toko          Toko            `json:"toko"` // foreign key
	CategoryID    uint            `json:"category_id" gorm:"not null"`
	Category      Category        `json:"category"` // foreign key
	Photos        []ProductPhotos `json:"ProductLogID" gorm:"foreignKey:ProductLogID"`
}

type ProductPhotos struct {
	gorm.Model
	ProductID    uint   `json:"product_id" gorm:"not null"`
	ProductLogID uint   `json:"product_log_id" gorm:"not null"`
	Url          string `json:"url"`
}
