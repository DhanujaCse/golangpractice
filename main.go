package main

import (
	"fmt"
	arrayws "golangprg/arraysssss"
	"golangprg/stringsfunct"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

func fun1(wq *sync.WaitGroup) {

	fmt.Println("Hello")
	wq.Done()
}

func fun2() {
	wg := &sync.WaitGroup{}
	fmt.Println("world")
	wg.Done()
}

func GetAddr(c arrayws.Address) {

	fmt.Println(c.Pincode)
}

func main() {
	// fmt.Println("Hi")
	// arrayws.Add()
	//var num int
	//fmt.Println("enter the number to check even or odd")
	//fmt.Scan(&num)

	//arrayws.EvenOrOdd(num)
	//arrayws.FabonacciSeries()
	// var n int
	// fmt.Println("enter the number to find sum of n numbers ")
	// fmt.Scan(&n)
	// sum := arrayws.SumOfNNumbers(n)
	// fmt.Println(sum)

	//// structures
	// fmt.Println("1 way of initializing structure")
	// a := arrayws.Address{"dhanuja", "bangalore", "karnataka", 560068}
	// fmt.Println(a)
	// fmt.Println("second way of initializing structure")

	// a1 := arrayws.Address{
	// 	Name:    "dhanu",
	// 	City:    "bangalore",
	// 	State:   "Karnataka",
	// 	Pincode: 765434,
	// }
	fmt.Println(arrayws)
	// GetAddr(a1)
	// var sum arrayws.MyInterface

	// sum = arrayws.CalculateAdd{10, 20}
	// fmt.Println(sum.Add())
	// fmt.Println(sum.Sub())

	// var digits arrayws.Multiply

	// digits = arrayws.DigitsNum{20, 20}
	// fmt.Println(digits.Divide())
	// fmt.Println(digits.Multiply1())

	//database

	// mux1 := controller.Register()

	// fmt.Println("done")

	// db := model.Connect()
	// //defer is used to make that line run at the end of the program//
	// defer db.Close()

	// log.Fatal(http.ListenAndServe("localhost:3000", mux1))

	// wq := &sync.WaitGroup{}

	// wq.Add(3)
	// // arrayws.Goto()
	// go fun1(wq)
	// go func() {

	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("happy")
	// 	wq.Done()

	// }()
	// go func() {
	// 	fmt.Println("world")
	// 	time.Sleep(4 * time.Second)
	// 	wq.Done()

	// }()

	// wq.Wait()
	// fmt.Println("done")

	// c := make(chan int, 2)

	// go func() {
	// 	c <- 1
	// 	c <- 4
	// 	c <- 6
	// 	c <- 9
	// 	c <- 8
	// 	close(c)
	// }()

	// for i := range c {
	// 	fmt.Println(i)
	// }

	///String functions

	//stringsfunct.StringFunc()
	stringsfunct.ReplaceWhiteSpaces()
	stringsfunct.MultilineStrings()
	stringsfunct.CompareStrings()
	stringsfunct.CheckStringIsPresentOrNot()

}
