package dao

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	NamaKategori string       `json:"nama_category"`
	ProductID    uint         `json:"product_id" gorm:"not null"`
	Product      []Product    `json:"product"` // foreign key
	ProductLogID uint         `json:"product_log_id" gorm:"not null"`
	ProductLog   []LogProduct `json:"product_log"` // foreign key
}
