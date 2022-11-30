package dao

import (
	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	NamaToko string    `json:"nama_toko"`
	UrlFoto  string    `json:"url_foto"`
	UserID   uint      `json:"user_id" gorm:"not null"`
	User     User      `json:"user"` // foreign key
	Product  []Product `json:"product"`
}
