package main

import (
	"log"
	"net/http"

	"github.com/SRdev04/products-restapi-mux/config"
	"github.com/SRdev04/products-restapi-mux/controller"
	"github.com/SRdev04/products-restapi-mux/middleware"
	"github.com/SRdev04/products-restapi-mux/service"
	"github.com/gorilla/mux"
)

func main() {
	//ini untuk mengkoneksi ke database
	db := config.ConnectionMysql()

	//untuk menghandle routers
	router := mux.NewRouter()

	//kita buat dependency yang dibutuhkan
	serviceProducts := service.NewProductsImplementasi(db)
	controllerProducts := controller.NewProductsCtrlImplementasi(serviceProducts)

	//ini untuk menampilkan request semua products
	router.HandleFunc("/products", controllerProducts.ShowAll).Methods("GET")

	//ini untuk menampilkan request products dengan Idnya
	router.HandleFunc("/product/{id}", controllerProducts.ShowById).Methods("GET")

	//ini untuk menampung request untuk membuat products
	router.HandleFunc("/product", controllerProducts.CreateProduct).Methods("POST")

	//ini untuk menampung request untuk mengubah nama,harga,quantity, dan deskripsi
	//lewat parameter idnya atau kita buat agar bisa dengan id bodynya
	//jika user lupa memasukkan id paramaeternya
	router.HandleFunc("/product/{id}", controllerProducts.UpdateProduct).Methods("PUT")

	//ini untuk mendelete product dengan id
	router.HandleFunc("/product/{id}", controllerProducts.DeleteProduct).Methods("DELETE")

	//untuk menampilkan server berjalan di port berapa
	log.Println("Running on port:7777")

	//menjalankan server dengan default localhost dan port 7777
	http.ListenAndServe(":7777", middleware.NewAuthMiddelware(router))

}
