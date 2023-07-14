package main

import (
	"encoding/json"
	"errors"
	"golangprg/apiexxxx/programs1"
	"net/http"
)

type employee struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	createEmployeeHanlder := http.HandlerFunc(createEmployee)
	http.Handle("/employee", createEmployeeHanlder)
	http.HandleFunc("/basicAuth", programs1.HandlerFuncReq)
	http.ListenAndServe(":8080", nil)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/json" {
		errorResponse(w, "Content Type is not application/json", http.StatusUnsupportedMediaType)
		return
	}
	var e employee
	var unmarshalErr *json.UnmarshalTypeError

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&e)
	if err != nil {
		if errors.As(err, &unmarshalErr) {
			errorResponse(w, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field, http.StatusBadRequest)
		} else {
			errorResponse(w, "Bad Request "+err.Error(), http.StatusBadRequest)
		}
		return
	}
	errorResponse(w, "Success", http.StatusOK)
	return
}

func errorResponse(w http.ResponseWriter, message string, httpStatusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatusCode)
	resp := make(map[string]string)
	resp["message"] = message
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
}
