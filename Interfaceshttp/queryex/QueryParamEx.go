package queryex

import (
	"fmt"
	"net/http"
)

func HttpQueryEx(w http.ResponseWriter, r *http.Request) {
	//query := r.URL.Query()
	check := r.Header
	// filters := query["filters"]
	// if len(filters) == 0 {
	// 	fmt.Println("filters are not present")
	// 	w.Write([]byte("no filters found"))
	// }
	//checking if key is present or not in  the header
	present, ok := check["Content-Type"]
	if ok {

		fmt.Println("its present")
		fmt.Println(present)

	} else {
		fmt.Println("its not found")
	}

	headerValues := r.Header.Values("content-type")
	fmt.Printf("header valuesssss is %v\n", headerValues)
	headerGet := r.Header.Get("name")
	fmt.Printf("header gettts is %v\n", headerGet)
	headeronly := r.Header
	fmt.Println("headeronly\n", headeronly)
	w.Header().Add("name", "dhanu")
	w.Header().Add("name2", "mamatha")
	w.Header().Set("name2", "swetha")
	w.Header().Add("name", "thanu")

	fmt.Println(w.Header().Values("name"))
	w.WriteHeader(200)

	w.Write([]byte("done"))
}
