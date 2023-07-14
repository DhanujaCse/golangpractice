package goexprg

import (
	"fmt"
	"sync"
)

func WaitGroupEx(wq *sync.WaitGroup) {
	fmt.Print("waiting")
	wq.Done()
}
