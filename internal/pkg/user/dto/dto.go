package userdto

import "tugas_akhir/internal/dao"

type UserUpdateReq struct {
	Nama         string `json:"nama,omitempty"`
	KataSandi    string `json:"kata_sandi,omitempty" validate:"required,min=6"`
	NoTelp       string `json:"no_telp,omitempty"`
	TanggalLahir string `json:"tanggal_Lahir" `
	Tentang      string `json:"tentang,omitempty"`
	Perkerjaan   string `json:"pekerjaan,omitempty"`
	Email        string `json:"email,omitempty" validate:"email"`
	IDProvinsi   string `json:"id_provinsi,omitempty"`
	IDKota       string `json:"id_kota,omitempty"`
}

type UserResp struct {
	ID           uint         `json:"id"`
	Nama         string       `json:"nama"`
	NoTelp       string       `json:"no_telp"`
	TanggalLahir string       `json:"tanggal_Lahir"`
	Tentang      string       `json:"tentang"`
	Perkerjaan   string       `json:"pekerjaan"`
	Email        string       `json:"email"`
	IDProvinsi   dao.Province `json:"id_provinsi"`
	IDKota       dao.City     `json:"id_kota"`
	// Alamat       []Alamat  `json:"alamat"`
	// Toko         Toko      `gorm:"-"`
	// Trx          []Trx     `json:"trx"`
}
