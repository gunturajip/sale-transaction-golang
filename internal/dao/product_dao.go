package dao

import (
	"gorm.io/gorm"
)

type PhotoProduct struct {
	gorm.Model
	Product Product `json:"product"`
}

type Product struct {
	gorm.Model
	NamaProduk    string          `json:"nama_produk"`
	Slug          string          `json:"slug"`
	HargaReseler  string          `json:"harga_reseler"`
	HargaKonsumen string          `json:"harga_konsumen"`
	Stok          int             `json:"stok"`
	Deskripsi     string          `json:"deskripsi" gorm:"type:text"`
	Toko          Toko            `json:"toko"`
	Category      Category        `json:"category"`
	Photos        []ProductPhotos `json:"photos"`
}

type LogProduct struct {
	gorm.Model
	NamaProduk    string          `json:"nama_produk"`
	Slug          string          `json:"slug"`
	HargaReseler  string          `json:"harga_reseler"`
	HargaKonsumen string          `json:"harga_konsumen"`
	Deskripsi     string          `json:"deskripsi" gorm:"type:text"`
	Toko          Toko            `json:"toko"`
	Category      Category        `json:"category"`
	Photos        []ProductPhotos `json:"photos"`
}

type ProductPhotos struct {
	gorm.Model
	Product Product `json:"product"`
	url     string  `json:"url"`
}
