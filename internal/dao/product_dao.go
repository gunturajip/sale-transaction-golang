package dao

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	NamaProduk    string           `json:"nama_produk"`
	Slug          string           `json:"slug"`
	HargaReseler  int              `json:"harga_reseller"`
	HargaKonsumen int              `json:"harga_konsumen"`
	Stok          int              `json:"stok"`
	Deskripsi     string           `json:"deskripsi" gorm:"type:text"`
	TokoID        uint             `json:"toko_id" gorm:"not null"`
	Toko          *Toko            `json:"toko"` // foreign key
	CategoryID    uint             `json:"category_id" gorm:"not null"`
	Category      *Category        `json:"category"` // foreign key
	Photos        []*ProductPhotos `json:"photos" gorm:"foreignKey:ProductID;"`
}

type LogProduct struct {
	gorm.Model
	ProductID     uint             `json:"product_id" gorm:"not null"`
	NamaProduk    string           `json:"nama_produk"`
	Slug          string           `json:"slug"`
	HargaReseler  int              `json:"harga_reseller"`
	HargaKonsumen int              `json:"harga_konsumen"`
	Deskripsi     string           `json:"deskripsi" gorm:"type:text"`
	TokoID        uint             `json:"toko_id" gorm:"not null"`
	Toko          *Toko            `json:"toko"` // foreign key
	CategoryID    uint             `json:"category_id" gorm:"not null"`
	Category      *Category        `json:"category"` // foreign key
	Photos        []*ProductPhotos `json:"photos" gorm:"foreignKey:ProductID;references:ProductID"`
}

type ProductPhotos struct {
	gorm.Model
	ProductID uint   `json:"product_id" gorm:"not null"`
	Url       string `json:"url"`
}

type ProductTotalPrice struct {
	Product
	Totalprice   int `gorm:"-"`
	Qty          int `gorm:"-"`
	LogProductID int `gorm:"-"`
}

// TableName overrides the table name used by ProductTotalPrice to `products`
func (ProductTotalPrice) TableName() string {
	return "products"
}
