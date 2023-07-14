package programs1

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type Login struct {
	Username string
	Password string
}

var (
	User     = "dhanuja"
	Password = "Dhanu@123"
)

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func HandlerFuncReq(w http.ResponseWriter, r *http.Request) {
	data1 := Login{}
	json.NewDecoder(r.Body).Decode(&data1)

	r.Header.Add("Authorization", "Basic "+basicAuth(data1.Username, data1.Password))

	u, p, ok := r.BasicAuth()

	if !ok {
		fmt.Println("error in basic auth")
		w.WriteHeader(401)
		return
	}
	if u != User {
		fmt.Println("user name not matching")
		w.WriteHeader(401)
		return
	}
	if p != Password {
		fmt.Println("password not matching")
		w.WriteHeader(401)
		return
	}
	w.Write([]byte(u))
	w.WriteHeader(200)

}
