package dao

import "gorm.io/gorm"

type Trx struct {
	gorm.Model
	HargaTotal  int         `json:"harga_total"`
	KodeInvoice string      `json:"kode_invoice"`
	MethodBayar string      `json:"method_bayar"`
	User        User        `json:"user"`
	AlamatKirim Alamat      `json:"alamat_kirim" gorm:"column:alamat_kirim"`
	DetailTrx   []DetailTrx `json:"detail_trx"`
}

type DetailTrx struct {
	gorm.Model
	Trx       Trx        `json:"id_trx"`
	Product   LogProduct `json:"product"`
	Toko      Toko       `json:"toko"`
	Kuantitas int        `json:"kuantitas"`
	HagaTotal int        `json:"harga_total"`
}
