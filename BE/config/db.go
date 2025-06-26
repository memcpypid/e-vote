package config

import (
	"E-vote/E-voteService/models"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:password@tcp(127.0.0.0:3306)/e_vote?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Koneksi ke database gagal: ", err)
	} else {
		fmt.Println("Berhasil terhubung ke database")
	}

	if err := db.AutoMigrate(
		&models.AkunMahasiswa{},
		&models.DataMahasiswa{},
		&models.WilayahTPS{},
		&models.DaerahTPS{},
		&models.TPS{},
		&models.ImageDataMahasiswa{},
		&models.SessionLogin{},
	); err != nil {
		log.Fatalf("Gagal melakukan migrasi: %v", err)
	}

	log.Println("Migrasi database berhasil!")
	DB = db
}
