package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	go FileServer()
	r := mux.NewRouter()
	//fileserver

	//routing
	r.HandleFunc("/search", FindALL).Methods("GET")
	r.HandleFunc("/search/{cost}", FindByCost).Methods("GET")
	r.HandleFunc("/search/{country}", FindByCountry).Methods("GET")
	err := http.ListenAndServe(":3001", r)
	if err != nil {
		fmt.Println(err)
	}
}

// FileServer function to activate file server
func FileServer() {
	fs := http.FileServer(http.Dir("../temp/fs"))
	http.Handle("/", fs)

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
