package dao

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	NamaKategori string    `json:"nama_category"`
	Product      []Product `json:"product"` // foreign key
}
