package categorydto

type CategoryReq struct {
	NamaKategori string `json:"nama_category" validate:"required"`
}

type CategoryResp struct {
	ID           uint   `json:"id"`
	NamaKategori string `json:"nama_category" validate:"required"`
}
