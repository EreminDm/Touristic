package main

import (
	"encoding/json"
	"net/http"
)

// type ProductAPI struct {
// }

func FindALL(response http.ResponseWriter, request *http.Request) {
	db, err := Connect()
	if err != nil {
		respondWithError(response, http.StatusBadRequest, err.Error())
		return
	}
	ProductModel := ProductModel{
		Db: db,
	}
	product, err2 := ProductModel.FindAll()
	if err2 != nil {
		respondWithError(response, http.StatusBadRequest, err2.Error())
		return
	}
	respondWithJson(response, http.StatusOK, product)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func FindByCost(responce http.ResponseWriter, request *http.Request) {

}

func FindByCountry(responce http.ResponseWriter, request *http.Request) {

}
