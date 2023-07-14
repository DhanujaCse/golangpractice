package controller

import (
	"encoding/json"
	"fmt"
	"golangprg/model"
	"golangprg/views"
	"net/http"
)

func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Responssss{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}

	}
}
func ping1() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := views.Responssss1{
			Code:  http.StatusCreated,
			Data1: "done",
		}
		json.NewEncoder(w).Encode(data)
	}
}
func dataPost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			data1 := views.PostData{}
			fmt.Println(data1)
			json.NewDecoder(r.Body).Decode(&data1)
			if err := model.CreatePost(data1.FirstName, data1.LastName); err != nil {
				w.Write([]byte("error"))
				return

			}
			w.WriteHeader((http.StatusCreated))

			json.NewEncoder(w).Encode(data1)

		} else if r.Method == http.MethodGet {
			data, err := model.ReadData()
			if err != nil {
				w.Write([]byte(err.Error()))

			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			return
		}

		fmt.Println("data inserted")

	}
}
