package config

import (
	"time"

	"github.com/SRdev04/products-restapi-mux/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionMysql() *gorm.DB {
	//kita buat koneksi ke database
	db, err := gorm.Open(mysql.Open("root:sahrulramadhan@tcp(localhost:3306)/products_restapi_mux?parseTime=true"))
	if err != nil {
		panic(err)
	}
	//kita atur poolnya secara manual sesuai yang kebutuhan
	//agar kita tidak membuka tutup koneksi
	Pool, _ := db.DB()

	//saat traffic koneksi ke database naik, ini akan membuka hingga 110 koneksi
	Pool.SetMaxOpenConns(110)
	//saat traffic koneksi ke database mulai turun, ini akan menjaga hingga 100 koneksi agar bisa digunakan kembali
	//namun tidak menutupnya, dan 10 koneksi akan terus aktif
	Pool.SetMaxIdleConns(100)
	//ini akan mengatur jika dalam 5 menit koneksi tidak digunakan, akan dibuang atau ditutup
	//namun saat semua koneksi digunakan itu akan tetap aktif
	Pool.SetConnMaxIdleTime(5 * time.Minute)
	//ini untuk memperbarui koneksi setiap 60 menit,
	Pool.SetConnMaxLifetime(60 * time.Minute)

	db.AutoMigrate(&models.Products{})

	DB = db
	return DB
}
