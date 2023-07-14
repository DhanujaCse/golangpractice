package channelsex

import "fmt"

func LenCapOfChannel() {
	ch := make(chan int, 3)

	ch <- 3
	fmt.Println("len of channel ", len(ch))
	fmt.Println("capacity of channel", cap(ch))
	ch <- 3
	fmt.Println("len of channel ", len(ch))
	fmt.Println("capacity of channel", cap(ch))
	ch <- 3
	fmt.Println("len of channel ", len(ch))
	fmt.Println("capacity of channel", cap(ch))

}
func UnBufferedChan() {
	ch := make(chan int)
	fmt.Println("len of channel of unbuffered channel  ", len(ch))
	fmt.Println("capacity of channel of unbuffered channel", cap(ch))

}
func NilChannelLen() {
	var ch chan int
	fmt.Println("len of channel of nil channel", len(ch))
	fmt.Println("capacity of channel of nil channel", cap(ch))

}
