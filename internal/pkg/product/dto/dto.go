package productdto

type ProductFilter struct {
	Limit         int    `query:"limit"`
	Page          int    `query:"page"`
	NamaProduk    string `query:"nama_produk"`
	Slug          string `query:"slug"`
	HargaReseler  int    `query:"harga_reseler"`
	HargaKonsumen int    `query:"harga_konsumen"`
	Stok          int    `query:"stok"`
	Deskripsi     string `query:"deskripsi"`
	TokoID        uint   `query:"toko_id"`
	CategoryID    uint   `query:"category_id"`
	// Photos        []ProductPhotos `json:"photos"`
}

type ProductReqCreate struct {
	NamaProduk    string `form:"nama_produk" json:"nama_produk" validate:"required"`
	HargaReseler  int    `form:"harga_reseller" json:"harga_reseller" validate:"required,numeric"`
	HargaKonsumen int    `form:"harga_konsumen" json:"harga_konsumen" validate:"required,numeric"`
	Stok          int    `form:"stok" json:"stok" validate:"required,numeric"`
	Deskripsi     string `form:"deskripsi" json:"deskripsi" validate:"required"`
	// TokoID        uint   `json:"toko_id" validate:"required,numeric"`
	CategoryID uint `form:"category_id" json:"category_id" validate:"required,numeric"`
	// Photos        []ProductPhotos `json:"photos"`
}

type ProductReqUpdate struct {
	NamaProduk    string `form:"nama_produk" json:"nama_produk"`
	HargaReseler  int    `form:"harga_reseller" json:"harga_reseller"`
	HargaKonsumen int    `form:"harga_konsumen" json:"harga_konsumen"`
	Stok          int    `form:"stok" json:"stok"`
	Deskripsi     string `form:"deskripsi" json:"deskripsi"`
	// TokoID        uint   `json:"toko_id"`
	CategoryID uint `form:"category_id" json:"category_id"`
	// Photos        []ProductPhotos `json:"photos"`
}

type ProductPhotos struct {
	ID        uint `json:"id"`
	ProductID uint `json:"product_id"`
	// ProductLogID uint   `json:"product_log_id"`
	Url string `json:"url"`
}

type ProductResp struct {
	ID            uint   `json:"id"`
	NamaProduk    string `json:"nama_produk"`
	Slug          string `json:"slug"`
	HargaReseler  int    `json:"harga_reseler"`
	HargaKonsumen int    `json:"harga_konsumen"`
	Stok          int    `json:"stok"`
	Deskripsi     string `json:"deskripsi" gorm:"type:text"`
	TokoID        uint   `json:"toko_id" gorm:"not null"`
	// Toko          Toko            `json:"toko"` // foreign key
	CategoryID uint `json:"category_id" gorm:"not null"`
	// Category      Category        `json:"category"` // foreign key
	Photos []ProductPhotos `json:"photos"`
}
