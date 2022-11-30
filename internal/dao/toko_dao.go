package dao

import (
	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	NamaToko string `json:"nama_toko,omitempty"`
	UrlFoto  string `json:"url_foto,omitempty"`
	UserID   uint   `json:"user_id,omitempty" gorm:"not null,omitempty"`
	// User     User      `json:"user"` // foreign key
	Product []Product `json:"product"`
}
