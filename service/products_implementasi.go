package service

import (
	"context"

	"github.com/SRdev04/products-restapi-mux/config"
	"github.com/SRdev04/products-restapi-mux/models"
	"gorm.io/gorm"
)

type ProductsImplementasi struct {
	DB *gorm.DB
}

//kita buat dependency yang dibutuhkan
func NewProductsImplementasi(DB *gorm.DB) ProductsRepository {
	return &ProductsImplementasi{
		DB: DB,
	}
}

func (service *ProductsImplementasi) GetAll(ctx context.Context) ([]models.Products, error) {
	//ini untuk menampung hasil query dan nnti akan direturn
	var products []models.Products

	//ini akan menghasilkan berhasil atau tidaknya query, bukan valuenya
	result := config.DB.Find(&products)
	//lalu kita cek apakah gagal atau berhasil
	if result.Error != nil {
		//dsini gagal melakukan query
		return products, result.Error
	}
	// dsni berhasil query dan kita balikan querynya, dan errornya nil
	return products, nil
}

func (service *ProductsImplementasi) GetByID(ctx context.Context, id int) (models.Products, error) {
	// membuat penampung untuk hasil query dan kita return productnya, dan errornya
	var product models.Products

	// ini akan menghasilkan berhasil atau tidaknya query, bukan valuenya
	result := config.DB.First(&product, id)
	if result.Error != nil {
		//tidak ada produknya

		return product, result.Error
	}
	// dsni berarti ada product dengan id tersebut
	return product, nil
}

func (service *ProductsImplementasi) Create(ctx context.Context, product models.Products) (models.Products, error) {

	// ini akan menghasilkan berhasil atau tidaknya menambahkan product, bukan valuenya
	result := config.DB.Create(&product)
	if result.Error != nil {
		return product, result.Error
	}
	// dsni berhasil query dan kita balikan querynya, dan errornya nil
	return product, nil
}

func (service *ProductsImplementasi) Update(ctx context.Context, product models.Products) (models.Products, error) {

	// ini akan menghasilkan berhasil atau tidaknya query, bukan valuenya
	result := config.DB.Updates(&product).Where("id = ?", product.Id)
	if result.Error != nil {
		// terjadi error saat mengubah

		return product, result.Error
	}

	return product, nil
}

func (service *ProductsImplementasi) DeleteById(ctx context.Context, id int) (int, error) {
	//kita akan delete item by id tanpa membalikan apa2
	result := config.DB.Delete(&models.Products{}, id)
	if result.Error != nil {
		// terjadi error saat menghapus data
		return int(result.RowsAffected), result.Error
	} else if result.RowsAffected == 0 {
		//tidak ada yang terhapus, atau rowsnya tidak ada yang berubah
		return int(result.RowsAffected), result.Error

	}
	//berhasil menghapus data
	return int(result.RowsAffected), nil
}
