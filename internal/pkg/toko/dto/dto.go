package tokodto

type TokoFilterRequest struct {
	Limit int    `json:"limit" query:"limit"`
	Page  int    `json:"page" query:"page"`
	Name  string `json:"name" query:"nama"`
}

type TokoResp struct {
	ID       uint   `json:"id,omitempty"`
	NamaToko string `json:"nama_toko"`
	UrlFoto  string `json:"url_foto"`
	UserID   uint   `json:"user_id,omitempty"`
	// User     User      `json:"user"` // foreign key
	// Product []dao.Product `json:"product"`
}

type TokoPagination struct {
	Page  int        `json:"page"`
	Limit int        `json:"limit"`
	Data  []TokoResp `json:"data"`
}

type TokoUpdateReq struct {
	NamaToko string `form:"nama_toko,omitempty"`
	UrlFoto  string `form:"photo,omitempty"`
	UserID   uint   `json:"user_id,omitempty"`
}
