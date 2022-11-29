package mysql

import (
	"fmt"
	"log"
	"tugas_akhir/internal/dao"

	"gorm.io/gorm"
)

func RunMigration(mysqlDB *gorm.DB) {
	err := mysqlDB.AutoMigrate(
		&dao.User{},
		&dao.Alamat{},
		&dao.Toko{},
		&dao.Product{},
		&dao.LogProduct{},
		&dao.ProductPhotos{},
		&dao.Category{},
		&dao.DetailTrx{},
		&dao.Trx{},
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
