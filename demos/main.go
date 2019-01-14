package main

import "fmt"

// Counter iterate from 0 to 999 and pass each value via a channel
func counter(out chan<- int) {
	for x := 0; x < 1000; x++ {
		out <- x
	}
	close(out)
}

// Squarer take values coming from a channel and pass its squared value in another
func squarer(out chan<- int, in <-chan int) {
	for sq := range in {
		out <- sq * sq
	}
	close(out)
}

// Printer logs every value passed down by a channel
func printer(in <-chan int) {
	for sq := range in {
		fmt.Println(sq)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
