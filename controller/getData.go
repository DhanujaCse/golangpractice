package controller

import (
	"encoding/json"
	"golangprg/model"
	"net/http"
)

func getData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.ReadSpecificData(name)
			if err != nil {
				w.Write([]byte(err.Error()))

			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			return

		} else if r.Method == http.MethodDelete {
			name := r.URL.Query().Get("name")
			if err := model.DeleteSpecificData1(name); err != nil {
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(struct {
				Status string `json:status`
			}{

				"deleted Successfully"})
		}

	}

}

func deleteData() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			name := r.URL.Query().Get("name")
			data, err := model.DeleteSpecificData3(name)
			if err != nil {
				w.Write([]byte(err.Error()))

			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
			return

		}
	}
}
