package alamatdto

type AlamatFilter struct {
	JudulAlamat string `query:"judul_alamat"`
}
type AlamatReqCreate struct {
	JudulAlamat  string `json:"judul_alamat" validate:"required"`
	NamaPenerima string `json:"nama_penerima" validate:"required"`
	NoTelp       string `json:"no_telp" validate:"required"`
	DetailAlamat string `json:"detail_alamat" validate:"required"`
}

type AlamatReqUpdate struct {
	JudulAlamat  string `json:"judul_alamat,omitempty"`
	NamaPenerima string `json:"nama_penerima,omitempty"`
	NoTelp       string `json:"no_telp,omitempty"`
	DetailAlamat string `json:"detail_alamat,omitempty"`
}

type AlamatResp struct {
	ID           uint   `json:"id"`
	JudulAlamat  string `json:"judul_alamat"`
	NamaPenerima string `json:"nama_penerima"`
	NoTelp       string `json:"no_telp"`
	DetailAlamat string `json:"detail_alamat"`
}
