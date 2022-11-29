package dao

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	HargaTotal  int         `json:"harga_total"`
	KodeInvoice string      `json:"kode_invoice"`
	MethodBayar string      `json:"method_bayar"`
	UserID      uint        `json:"user_id" gorm:"not null"`
	User        User        `json:"user"` // foreign key
	AlamatID    uint        `json:"alamat_kirim" gorm:"not null;column:alamat_kirim"`
	Alamat      Alamat      `json:"alamat"` // foreign key
	DetailTrx   []DetailTrx `json:"detail_trx"`
}

type DetailTrx struct {
	gorm.Model
	TrxID        uint       `json:"trx_id" gorm:"not null"`
	Trx          Trx        `json:"Trx"` // foreign key
	LogProductID uint       `json:"log_product_id" gorm:"not null"`
	LogProduct   LogProduct `json:"log_product"` // foreign key
	TokoID       uint       `json:"toko_id" gorm:"not null"`
	Toko         Toko       `json:"toko"` // foreign key
	Kuantitas    int        `json:"kuantitas"`
	HagaTotal    int        `json:"harga_total"`
}
