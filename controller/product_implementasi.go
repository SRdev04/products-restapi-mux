package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/SRdev04/products-restapi-mux/helper"
	"github.com/SRdev04/products-restapi-mux/models"
	"github.com/SRdev04/products-restapi-mux/service"
	"github.com/gorilla/mux"
)

type ProductsCtrlImplementasi struct {
	Service service.ProductsRepository
}

func NewProductsCtrlImplementasi(service service.ProductsRepository) ProductsCtrlRepository {
	return &ProductsCtrlImplementasi{
		Service: service,
	}
}

func (ctrl *ProductsCtrlImplementasi) ShowAll(w http.ResponseWriter, r *http.Request) {
	//untuk memberikan informasi kalau client terkonek dengan function ini
	log.Println("connect to Products")
	products, err := ctrl.Service.GetAll(r.Context())
	if err != nil {
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return

	}

	//jika tidak ada error kita balikan products dan oke dengan helper yang dibuat
	helper.EncodeResponses(w, http.StatusOK, products)

}

func (ctrl *ProductsCtrlImplementasi) ShowById(w http.ResponseWriter, r *http.Request) {
	//untuk memberikan informasi kalau client terkonek dengan function ini
	log.Println("connect to Product Id")

	//kita akan tangkap paramater idnya
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		//terjadi error saat parsing parameter idnya
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	//jika dapat parameternya, maka kita panggil service
	product, err := ctrl.Service.GetByID(r.Context(), id)
	if err != nil {
		helper.EncodeResponError(w, http.StatusNotFound, err.Error())
		return
	}

	//jika tidak ada error kita balikan products dan oke dengan helper yang dibuat
	helper.EncodeResponses(w, http.StatusOK, product)

}

func (ctrl *ProductsCtrlImplementasi) CreateProduct(w http.ResponseWriter, r *http.Request) {
	//untuk memberikan informasi kalau client terkonek dengan function ini
	log.Println("connected create product")

	//kita buat penampung dari body request yang diterima
	var product models.Products

	//kita akan baca request bodynnya dengan package json
	err := helper.DecodeRequest(r, &product)
	if err != nil {
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	//jika tidak ada error kita panggil service create productnya
	result, err := ctrl.Service.Create(r.Context(), product)
	if err != nil {
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	//jika tidak ada error kita balikan products dan oke dengan helper yang dibuat
	helper.EncodeResponses(w, http.StatusOK, result)
}
func (ctrl *ProductsCtrlImplementasi) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	//untuk memberikan informasi kalau client terkonek dengan function ini
	log.Println("connect Update")

	//kita buat penampung dari body request yang diterima
	var product models.Products

	//kita tangkap dahulu query parameternya
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return
	}
	product.Id = int64(id)

	//kita akan baca request bodynya dengan package json
	//dan akan menulis isi body requestnya ke product penampungan
	dcoder := json.NewDecoder(r.Body)
	dcoder.Decode(&product)

	//jika tidak ada error kita panggil service create productnya
	result, err := ctrl.Service.Update(r.Context(), product)
	if err != nil {
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	//jika tidak ada error kita balikan products dan oke dengan helper yang dibuat
	helper.EncodeResponses(w, http.StatusOK, result)
}

func (ctrl *ProductsCtrlImplementasi) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	//untuk memberikan informasi kalau client terkonek dengan function ini
	log.Println("connect Delete")

	//kita tangkap paramaeternya
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return
	}

	//jika dapat parameternya, maka kita panggil service delete
	rowsAffected, err := ctrl.Service.DeleteById(r.Context(), id)
	//kita cek apakah ada error atau tidak
	if err != nil {
		//terdapat error saat menghapus
		helper.EncodeResponError(w, http.StatusBadRequest, err.Error())
		return

	} else if rowsAffected == 0 {
		//tidak ada yang terhapus, atau rowsnya tidak ada yang berubah
		//kita balikkan badrequest dan pesannya

		helper.EncodeResponError(w, http.StatusBadRequest, "cannot delete,cek request id")
		return

	}
	//jika berhasil kita akan balikan oke saja
	helper.EncodeResponses(w, http.StatusOK, nil)
}
