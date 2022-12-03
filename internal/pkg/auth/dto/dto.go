package authdto

import "tugas_akhir/internal/dao"

type LoginRequest struct {
	NoTelp   string `json:"no_telp" validate:"required"`
	Password string `json:"kata_sandi" validate:"required"`
}

type RegisterRequest struct {
	Nama         string `json:"nama,omitempty" validate:"required"`
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
	Nama         string       `json:"nama,omitempty"`
	NoTelp       string       `json:"no_telp,omitempty"`
	TanggalLahir string       `json:"tanggal_Lahir,omitempty"`
	Tentang      string       `json:"tentang,omitempty"`
	Perkerjaan   string       `json:"pekerjaan,omitempty"`
	Email        string       `json:"email,omitempty"`
	IDProvinsi   dao.Province `json:"id_provinsi,omitempty"`
	IDKota       dao.City     `json:"id_kota,omitempty"`
	Token        string       `json:"token,omitempty"`
	// Alamat       []Alamat  `json:"alamat"`
	// Toko         Toko      `gorm:"-"`
	// Trx          []Trx     `json:"trx"`
}
