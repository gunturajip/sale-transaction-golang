package dao

import "gorm.io/gorm"

type Alamat struct {
	gorm.Model
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat" gorm:"type:text"`
	UserID       uint   `json:"user_id" gorm:"not null"`
}
