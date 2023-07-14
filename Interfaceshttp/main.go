package main

import (
	"fmt"
	"golangprg/Interfaceshttp/queryex"
	"net/http"
	"strings"
)

func main() {
	getProducts := http.HandlerFunc(queryex.HttpQueryEx)
	http.Handle("/products", getProducts)
	http.ListenAndServe(":8080", nil)
}
func httpQueryEx1(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	filters, present := query["filters"]
	if !present || len(filters) == 0 {
		fmt.Println("filters are not present")
	}
	w.WriteHeader(200)
	
	w.Write([]byte(strings.Join(filters, ",")))

}
