package trxusecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"
	"tugas_akhir/internal/dao"
	"tugas_akhir/internal/helper"
	alamatrepository "tugas_akhir/internal/pkg/alamat/repository"
	productrepository "tugas_akhir/internal/pkg/product/repository"
	trxdto "tugas_akhir/internal/pkg/trx/dto"
	trxrepository "tugas_akhir/internal/pkg/trx/repository"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TrxUseCase interface {
	GetAllTrxs(ctx context.Context, filter trxdto.TrxFilter) (res trxdto.TrxPagination, err *helper.ErrorStruct)
	GetTrxByID(ctx context.Context, trxid string) (res []trxdto.TrxRes, err *helper.ErrorStruct)
	CreateTrx(ctx context.Context, userid string, data trxdto.TrxReq) (res interface{}, err *helper.ErrorStruct)
}

type TrxUseCaseImpl struct {
	trxrepository     trxrepository.TrxRepository
	productrepository productrepository.ProductRepository
	alamatrepository  alamatrepository.AlamatRepository
	db                *gorm.DB
}

func NewTrxUseCase(trxrepository trxrepository.TrxRepository, productrepository productrepository.ProductRepository, alamatrepository alamatrepository.AlamatRepository, db *gorm.DB) TrxUseCase {
	return &TrxUseCaseImpl{
		trxrepository:     trxrepository,
		productrepository: productrepository,
		db:                db,
		alamatrepository:  alamatrepository,
	}

}

func (tu *TrxUseCaseImpl) GetAllTrxs(ctx context.Context, filter trxdto.TrxFilter) (res trxdto.TrxPagination, err *helper.ErrorStruct) {
	if filter.Limit == 0 {
		filter.Limit = 10
	}
	if filter.Page == 0 {
		filter.Page = 1
	}

	_, errRepo := tu.trxrepository.GetAllTrxs(ctx, filter)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Trx"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	// for _, v := range resRepo {
	// 	result := trxdto.TrxResp{
	// 		ID:            v.ID,
	// 		NamaProduk:    v.NamaProduk,
	// 		Slug:          v.Slug,
	// 		HargaReseler:  v.HargaReseler,
	// 		HargaKonsumen: v.HargaKonsumen,
	// 		Stok:          v.Stok,
	// 		Deskripsi:     v.Deskripsi,
	// 		Toko: tokodto.TokoResp{
	// 			ID:       v.TokoID,
	// 			NamaToko: v.Toko.NamaToko,
	// 			UrlFoto:  v.Toko.UrlFoto,
	// 		},
	// 		Category: categorydto.CategoryResp{
	// 			ID:           v.Category.ID,
	// 			NamaKategori: v.Category.NamaKategori,
	// 		},
	// 	}

	// 	for _, photo := range v.Photos {
	// 		result.Photos = append(result.Photos, trxdto.TrxPhotos{
	// 			ID:    photo.ID,
	// 			TrxID: photo.TrxID,
	// 			Url:   photo.Url,
	// 		})
	// 	}

	// 	res.Data = append(res.Data, result)
	// }

	res.Limit = filter.Limit
	res.Page = filter.Page

	return res, nil
}

func (tu *TrxUseCaseImpl) GetTrxByID(ctx context.Context, trxid string) (res []trxdto.TrxRes, err *helper.ErrorStruct) {
	_, errRepo := tu.trxrepository.GetTrxByID(ctx, trxid)
	if errors.Is(errRepo, gorm.ErrRecordNotFound) {
		return res, &helper.ErrorStruct{
			Code: fiber.StatusNotFound,
			Err:  errors.New("No Data Trx"),
		}
	}

	if errRepo != nil {
		log.Println(errRepo)
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errRepo,
		}
	}

	// res = trxdto.TrxResp{
	// 	ID:            resRepo.ID,
	// 	NamaProduk:    resRepo.NamaProduk,
	// 	Slug:          resRepo.Slug,
	// 	HargaReseler:  resRepo.HargaReseler,
	// 	HargaKonsumen: resRepo.HargaKonsumen,
	// 	Stok:          resRepo.Stok,
	// 	Deskripsi:     resRepo.Deskripsi,
	// 	Toko: tokodto.TokoResp{
	// 		ID:       resRepo.TokoID,
	// 		NamaToko: resRepo.Toko.NamaToko,
	// 		UrlFoto:  resRepo.Toko.UrlFoto,
	// 	},
	// 	Category: categorydto.CategoryResp{
	// 		ID:           resRepo.Category.ID,
	// 		NamaKategori: resRepo.Category.NamaKategori,
	// 	},
	// }

	// for _, photo := range resRepo.Photos {
	// 	res.Photos = append(res.Photos, trxdto.TrxPhotos{
	// 		ID:    photo.ID,
	// 		TrxID: photo.TrxID,
	// 		Url:   photo.Url,
	// 	})
	// }

	return res, nil
}

func (tu *TrxUseCaseImpl) CreateTrx(ctx context.Context, userid string, data trxdto.TrxReq) (res interface{}, err *helper.ErrorStruct) {
	if errValidate := helper.Validate.Struct(data); errValidate != nil {
		log.Println(errValidate)
		return res, &helper.ErrorStruct{
			Err:  errValidate,
			Code: fiber.StatusBadRequest,
		}
	}

	// INIT
	var hargaTotal int
	kodeInvoice := fmt.Sprint("INV-", time.Now().Unix())

	// TRANSACTION
	tx := tu.db.Begin()

	// CHECK IF HIS/HIM ALAMAT
	isAlamatValid := tu.isMyAlamat(ctx, tx, userid, data.AlamatID)
	if !isAlamatValid {
		tx.Rollback()
		return res, &helper.ErrorStruct{
			Code: fiber.StatusUnauthorized,
			Err:  errors.New("ALAMAT INVALID"),
		}
	}

	// GET DATA PRODUCT
	var sliceProductID []uint
	var mapQty = make(map[uint]int) // MAPPING QTY PRODUCT TO ID PRODUCT (PRODUCT ID AS KEY ; QTY AS VALUE)
	for _, detailTrx := range data.DetailTrx {
		sliceProductID = append(sliceProductID, detailTrx.IDProduct)
		mapQty[detailTrx.IDProduct] = detailTrx.Kuantitas
	}

	resProduct, errProduct := tu.productrepository.GetProductsBySliceID(ctx, tx, sliceProductID)
	if errProduct != nil {
		tx.Rollback()
		log.Println(errProduct)
		return res, &helper.ErrorStruct{
			Err:  errProduct,
			Code: fiber.StatusBadRequest,
		}
	}

	// INSERT LOG PRODUCT
	var dataLogProduct []dao.LogProduct

	for _, v := range resProduct {
		dataLogProduct = append(dataLogProduct, dao.LogProduct{
			ProductID:     v.ID,
			NamaProduk:    v.NamaProduk,
			Slug:          v.Slug,
			HargaReseler:  v.HargaReseler,
			HargaKonsumen: v.HargaKonsumen,
			Deskripsi:     v.Deskripsi,
			TokoID:        v.TokoID,
			CategoryID:    v.CategoryID,
		})
	}

	resProductLog, errProductLog := tu.productrepository.CreateProductLog(ctx, tx, dataLogProduct)
	if errProductLog != nil {
		tx.Rollback()
		return res, &helper.ErrorStruct{
			Code: fiber.StatusBadRequest,
			Err:  errProduct,
		}
	}

	// MAP RESULT PRODUCTLOG
	var mapProductLog = make(map[uint]int) // PRODUCTD ID AS KEY, PRODUCT LOG AS VALUE
	for _, v := range resProductLog {
		mapProductLog[v.ProductID] = int(v.ID)
	}

	for i, v := range resProduct {
		resProduct[i].Totalprice = v.HargaKonsumen * mapQty[v.ID]
		resProduct[i].Qty = mapQty[v.ID]
		hargaTotal += v.HargaKonsumen * mapQty[v.ID]

		resProduct[i].LogProductID = mapProductLog[v.ID]
	}

	log.Println("INSERT PRODUCT LOG SUCCEED ")

	// PREPARE INSERT DATA TO TABLE TRX
	userIduint, errConv := strconv.ParseUint(userid, 10, 64)
	if errConv != nil {
		log.Println(errConv)
		tx.Rollback()
		return res, &helper.ErrorStruct{
			Err:  errConv,
			Code: fiber.StatusBadRequest,
		}
	}

	dataTrx := dao.Trx{
		HargaTotal:  hargaTotal,
		KodeInvoice: kodeInvoice,
		MethodBayar: data.MethodBayar,
		UserID:      uint(userIduint),
		AlamatID:    data.AlamatID,
		// DetailTrx:   []dao.DetailTrx{},
	}

	// Detail TRX
	for _, product := range resProduct {
		dataTrx.DetailTrx = append(dataTrx.DetailTrx, dao.DetailTrx{
			LogProductID: uint(mapProductLog[product.ID]),
			TokoID:       product.TokoID,
			Kuantitas:    product.Qty,
			HagaTotal:    product.Totalprice,
		})
	}

	// INSERT DATA TO TABLE TRX
	resTrx, errTrx := tu.trxrepository.CreateTrx(ctx, tx, dataTrx)
	if errTrx != nil {
		tx.Rollback()
		log.Println(errConv)
		return res, &helper.ErrorStruct{
			Err:  errConv,
			Code: fiber.StatusBadRequest,
		}
	}

	log.Println("INSERT PRODUCT LOG SUCCEED ")

	// UDPATE STOCK PRODUCT
	for _, v := range resProduct {
		resUpdateStock, errUpdateStock := tu.productrepository.UpdateProductStock(ctx, tx, v.ID, v.Qty)
		if errUpdateStock != nil {
			tx.Rollback()
			log.Println(errUpdateStock)
			return res, &helper.ErrorStruct{
				Err:  errUpdateStock,
				Code: fiber.StatusBadRequest,
			}
		}
		log.Println("resUpdateStock : ", resUpdateStock)
	}

	tx.Commit()

	return resTrx, nil
}

func (tu *TrxUseCaseImpl) isMyAlamat(ctx context.Context, tx *gorm.DB, userid string, alamatid uint) bool {
	res, err := tu.alamatrepository.GetAlamatByID(ctx, fmt.Sprint(alamatid))
	if err != nil {
		log.Println("isMyAlamat err : ", err)
		return false
	}

	userID, err := strconv.ParseUint(userid, 10, 64)
	if err != nil {
		log.Println(err)
		return false
	}

	return res.UserID == uint(userID)
}
