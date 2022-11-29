package authdto

import "time"

type LoginRequest struct {
	NoTelp   string `json:"no_telp" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Nama         string `json:"nama" validate:"required"`
	KataSandi    string `json:"kata_sandi" validate:"required,min=6"`
	NoTelp       string `json:"no_telp" validate:"required"`
	TanggalLahir string `json:"tanggal_Lahir" validate:"required"`
	Tentang      string `json:"tentang"`
	Perkerjaan   string `json:"pekerjaan" validate:"required"`
	Email        string `json:"email"  validate:"required,email"`
	IDProvinsi   string `json:"id_provinsi" validate:"required"`
	IDKota       string `json:"id_kota" validate:"required"`
}

type LoginResp struct {
	Nama         string    `json:"nama"`
	NoTelp       string    `json:"no_telp"`
	TanggalLahir time.Time `json:"tanggal_Lahir" `
	Tentang      string    `json:"tentang"`
	Perkerjaan   string    `json:"pekerjaan"`
	Email        string    `json:"email"`
	IDProvinsi   string    `json:"id_provinsi"`
	IDKota       string    `json:"id_kota"`
	// Alamat       []Alamat  `json:"alamat"`
	// Toko         Toko      `gorm:"-"`
	// Trx          []Trx     `json:"trx"`
}
