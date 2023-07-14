package main

import "golangprg/MapsEx/mapsexa"

func main() {
	// fmt.Println("maps")

	// emp := map[string]int{}
	// emp["dhanu"] = 40000
	// emp["Mamatha"] = 40000
	// fmt.Println(emp)
	// emp["dhanu"] = 50000 //updating the exixting key
	// val := emp["dhanu"]
	// fmt.Print("value of key dhanu is ")
	// fmt.Println(val)

	// delete(emp, "Mamatha") //delete a key from  map
	// fmt.Println(emp)

	// val1, ok := emp["dhanuj"] //check if key is present in map
	// if ok {
	// 	fmt.Println("the key is present", val1)
	// } else {
	// 	fmt.Println("not present")
	// }
	// fmt.Println("length of the map emp is ", len(emp)) //finding length of the map

	// for k, v := range emp {
	// 	fmt.Println("key : ", k+" value", v) //iterating over key and value
	// }
	// for k := range emp {
	// 	fmt.Println("key : ", k) //iterating over only key
	// }
	// for _, v := range emp {
	// 	fmt.Println("value : ", v) //iterating over only value
	// }

	// arr := []int{10, 20, 30, 20, 10, 40}
	// count := mapsassignment.CountOccurences(arr)
	// fmt.Println("count of occurences: ", count)
	// deletedArray := mapsassignment.RemoveDuplicates(arr)
	// fmt.Println("after deleting duplicates values", deletedArray)

	mapsexa.MapToJson()

}
