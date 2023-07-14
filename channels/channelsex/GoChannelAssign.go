package channelsex

import (
	"fmt"
)

func SendChan(ch chan<- int) {
	for i := 1; i <= 10; i++ {
		ch <- i
	}
}
func ReceiveChan(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}
