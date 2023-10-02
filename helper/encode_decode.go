package helper

import (
	"encoding/json"
	"net/http"
)

//kita akan membuat func yang jika sukkses

func EncodeResponses(w http.ResponseWriter, code int, load interface{}) {
	//kita encode atau mengubah dari golang ke json
	response, err := json.Marshal(load)
	if err != nil {
		//terjadi error dan kita return
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//kita buat headernya
	w.Header().Set("Content-Type", "application/json")

	//kita gunakan code agar response statusnya sesuai
	w.WriteHeader(code)

	//kita kirim responsenya
	w.Write(response)

}

//kita akan buat func yang jika terjadi error dalam requestnya
func EncodeResponError(w http.ResponseWriter, code int, msg string) {
	//kita akan buat func helper agar tidak bloated
	//dan digabung if err!=nil agar tidak banyak mengetikkan if err!=nil

	EncodeResponses(w, code, map[string]string{"message": msg})

}

//kita buat func untuk menerima request dari client, kita decode dahulu
func DecodeRequest(r *http.Request, product interface{}) error {
	//kita baca dan langsung tulis request bodynya
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&product)
	if err != nil {
		//terjadi err
		return err
	}
	return nil
}
