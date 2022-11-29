package dao

import (
	"time"

	"gorm.io/gorm"
)

type Toko struct {
	gorm.Model
	User         User      `json:"user"`
	KataSandi    string    `json:"kata_sandi"`
	NoTelp       string    `json:"no_telp" gorm:"unique"`
	TanggalLahir time.Time `json:"tanggal_Lahir" `
	Tentang      string    `json:"tentang" gorm:"type:text"`
	Perkerjaan   string    `json:"pekerjaan"`
	Email        string    `json:"email" gorm:"unique;not null"`
	IDProvinsi   string    `json:"id_provinsi"`
	IDKota       string    `json:"id_kota"`
	Product      []Product `json:"product"`
}
