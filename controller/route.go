package controller

import "net/http"

func Register() *http.ServeMux {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/ping", ping())
	mux1.HandleFunc("/ping1", ping1())
	mux1.HandleFunc("/post", dataPost())
	mux1.HandleFunc("/getdata/", getData())
	mux1.HandleFunc("/deletedata/", deleteData())

	return mux1
}
