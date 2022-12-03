package trxdto

import (
	alamatdto "tugas_akhir/internal/pkg/alamat/dto"
	productdto "tugas_akhir/internal/pkg/product/dto"
	tokodto "tugas_akhir/internal/pkg/toko/dto"
)

type TrxFilter struct {
	Limit int `quer:"limit"`
	Page  int `quer:"page"`
}

type TrxReq struct {
	MethodBayar string         `json:"method_bayar" validate:"required"`
	AlamatID    uint           `json:"alamat_kirim" validate:"required,numeric"`
	DetailTrx   []DetailTrxReq `json:"detail_trx" validate:"required"`
}

type DetailTrxReq struct {
	IDProduct uint `json:"product_id" validate:"required,numeric"`
	Kuantitas int  `json:"kuantitas" validate:"required,numeric"`
}

type TrxResp struct {
	ID          uint                 `json:"id"`
	HargaTotal  int                  `json:"harga_total"`
	KodeInvoice string               `json:"kode_invoice"`
	MethodBayar string               `json:"method_bayar"`
	Alamat      alamatdto.AlamatResp `json:"alamat_kirim"` // foreign key
	DetailTrx   []DetailTrxRes       `json:"detail_trx"`
}

type DetailTrxRes struct {
	LogProduct productdto.ProductResp `json:"product,omitempty"` // foreign key
	Toko       tokodto.TokoResp       `json:"toko"`
	Kuantitas  int                    `json:"kuantitas"`
	HagaTotal  int                    `json:"harga_total"`
}

type TrxPagination struct {
	Data  []TrxResp `json:"data"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
}
