package trxdto

import (
	alamatdto "tugas_akhir/internal/pkg/alamat/dto"
	productdto "tugas_akhir/internal/pkg/product/dto"
)

type TrxFilter struct {
	Limit int `quer:"limit"`
	Page  int `quer:"page"`
}

type TrxReq struct {
	// HargaTotal  int            `json:"harga_total"`
	MethodBayar string         `json:"method_bayar" validate:"required"`
	AlamatID    uint           `json:"alamat_kirim" validate:"required,numeric"`
	DetailTrx   []DetailTrxReq `json:"detail_trx" validate:"required"`
}

type DetailTrxReq struct {
	IDProduct uint `json:"product_id" validate:"required,numeric"`
	Kuantitas int  `json:"kuantitas" validate:"required,numeric"`
}

type TrxRes struct {
	ID          uint   `json:"id"`
	HargaTotal  int    `json:"harga_total"`
	KodeInvoice string `json:"kode_invoice"`
	MethodBayar string `json:"method_bayar"`
	// UserID      uint        `json:"user_id"`
	// User        User        `json:"user"` // foreign key
	AlamatID  uint                 `json:"alamat_kirim"`
	Alamat    alamatdto.AlamatResp `json:"alamat"` // foreign key
	DetailTrx []DetailTrxRes       `json:"detail_trx"`
}

type DetailTrxRes struct {
	TrxID uint `json:"trx_id"`
	// LogProductID uint                   `json:"product_id"`
	LogProduct productdto.ProductResp `json:"product"` // foreign key
	// IDProduct    uint                   `json:"product_id"`
	// TokoID    uint `json:"toko_id"`
	Kuantitas int `json:"kuantitas"`
	HagaTotal int `json:"harga_total"`
}

type TrxPagination struct {
	Data  []TrxRes `json:"data"`
	Page  int      `json:"page"`
	Limit int      `json:"limit"`
}
