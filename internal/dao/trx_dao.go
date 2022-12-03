package dao

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	HargaTotal  int         `json:"harga_total"`
	KodeInvoice string      `json:"kode_invoice"`
	MethodBayar string      `json:"method_bayar"`
	UserID      uint        `json:"user_id"`
	User        User        `json:"user"` // foreign key
	AlamatID    uint        `json:"alamat_kirim"`
	Alamat      Alamat      `json:"alamat"` // foreign key
	DetailTrx   []DetailTrx `json:"detail_trx"`
}

type DetailTrx struct {
	gorm.Model
	TrxID        uint       `json:"trx_id"`
	Trx          Trx        `json:"Trx"` // foreign key
	LogProductID uint       `json:"log_product_id"`
	LogProduct   LogProduct `json:"log_product"` // foreign key
	TokoID       uint       `json:"toko_id"`
	Toko         Toko       `json:"toko"` // foreign key
	Kuantitas    int        `json:"kuantitas"`
	HagaTotal    int        `json:"harga_total"`
}
