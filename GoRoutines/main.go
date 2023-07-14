package main

import (
	"fmt"
	"golangprg/GoRoutines/goexprg"
	"runtime"
	"sync"
	"time"
)

func func1(wq *sync.WaitGroup) {
	fmt.Println("go routine 1")
	wq.Done()
}

func main() {

	var wq sync.WaitGroup
	wq.Add(2)

	fmt.Println("go routines")

	go func1(&wq)
	go goexprg.PrintStat()
	fmt.Println(runtime.NumGoroutine())
	time.Sleep(2 * time.Second)

	fmt.Println("completed")

	fmt.Println(runtime.NumCPU()) //used to get number of logical processors available to go program
	fmt.Println(runtime.NumGoroutine())

	go goexprg.WaitGroupEx(&wq)

	wq.Wait()
	fmt.Println("done with go routines execution")
}
