package dao

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string    `json:"nama"`
	KataSandi    string    `json:"kata_sandi"`
	NoTelp       string    `json:"no_telp" gorm:"unique"`
	TanggalLahir time.Time `json:"tanggal_Lahir" `
	Tentang      string    `json:"tentang" gorm:"type:text"`
	Perkerjaan   string    `json:"pekerjaan"`
	Email        string    `json:"email" gorm:"unique;not null"`
	IDProvinsi   string    `json:"id_provinsi"`
	IDKota       string    `json:"id_kota"`
	Alamat       []Alamat  `json:"alamat"`
	Toko         Toko      `gorm:"foreignkey:UserID"`
	Trx          []Trx     `json:"trx"`
	IsAdmin      bool      `json:"is_admin" gorm:"default:false"`
}

type UserLogin struct {
	KataSandi string `json:"kata_sandi"`
	NoTelp    string `json:"no_telp" gorm:"unique"`
}
