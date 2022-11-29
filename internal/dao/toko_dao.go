package dao

import (
	"time"

	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	KataSandi    string    `json:"kata_sandi"`
	NoTelp       string    `json:"no_telp" gorm:"unique"`
	TanggalLahir time.Time `json:"tanggal_Lahir" `
	Tentang      string    `json:"tentang" gorm:"type:text"`
	Perkerjaan   string    `json:"pekerjaan"`
	Email        string    `json:"email" gorm:"unique;not null"`
	IDProvinsi   string    `json:"id_provinsi"`
	IDKota       string    `json:"id_kota"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	User         User      `json:"user"` // foreign key
	Product      []Product `json:"product"`
}
