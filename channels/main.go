package main

import (
	"fmt"
	"golangprg/channels/channelsex"
	"time"
)

func main() {
	// ch := make(chan int)
	// //go sum(ch, 3)
	// go sum1(ch)
	// ch <- 3
	// ch <- 4
	// ch <- 4
	// close(ch)

	// go send(ch)
	// <-ch

	// channelsex.LenCapOfChannel()
	// channelsex.UnBufferedChan()
	// channelsex.NilChannelLen()

	ch := make(chan int)
	go channelsex.SendChan(ch)
	go channelsex.ReceiveChan(ch)
	time.Sleep(time.Second * 1)
}
func sum(ch chan int, len int) {
	sum := 0
	for i := 0; i < len; i++ {
		sum = sum + <-ch
	}

	fmt.Println("sum", sum)

}

func send(ch chan int) {
	fmt.Println("Sending value to channnel start")
	ch <- 1
	fmt.Println("Sending value to channnel finish")
}
func sum1(ch chan int) {
	sum := 0
	for i := range ch {
		sum = sum + i
	}
	fmt.Println(sum)
}
