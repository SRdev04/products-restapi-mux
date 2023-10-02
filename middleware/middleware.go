package middleware

import (
	"net/http"

	"github.com/SRdev04/products-restapi-mux/helper"
)

type AuthMiddelware struct {
	Handler http.Handler
}

// kita akan buat function dependencnya
func NewAuthMiddelware(handler http.Handler) *AuthMiddelware {
	return &AuthMiddelware{
		Handler: handler,
	}
}

func (auth *AuthMiddelware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//kita tidak akan ribet2, hanya menggunakan ApiKey Dahulu
	value := r.Header.Get("X-ApiKey-X")

	//kita lakukan pengecekan value dari ApiKey
	if value == "Sahrul" {
		//Value sesuai dengan apikey yang kita mau
		auth.Handler.ServeHTTP(w, r)
	} else {

		//ApiKey tidak sesuai yang kita mau
		helper.EncodeResponError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}
}
