package mapsexa

import (
	"encoding/json"
	"fmt"
)

type employee struct {
	Name string
}

func MapToJson() {

	emp := make(map[int]string)
	emp[1] = "Dhanu"
	emp1 := make(map[string]employee)
	emp1["1"] = employee{"Dhanuja"}

	j, err := json.Marshal(emp) //converting map to json

	k, err1 := json.Marshal(emp1) //converting map to string
	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(j))
	if err1 != nil {
		fmt.Println("error")
	}
	fmt.Println(string(k))

	//converting json to map

	Jsontomap := make(map[int]string)
	json.Unmarshal(j, &Jsontomap)

	fmt.Println(Jsontomap)
	Jsontomap1 := make(map[string]employee)
	json.Unmarshal(k, &Jsontomap1)

	fmt.Println(Jsontomap1)
}
