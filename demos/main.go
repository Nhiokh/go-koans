package main

import "fmt"

// Counter
func counter(out chan<- int) {
	for x := 0; x < 1000; x++ {
		out <- x
	}
	close(out)
}

// Squarer
func squarer(out chan<- int, in <-chan int) {
	for sq := range in {
		out <- sq * sq
	}
	close(out)
}

// Printer
func printer(in <-chan int) {
	for sq := range in {
		fmt.Println(sq)
	}
}

func add(x int, y int) int {
	return x + y
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	var res int = add(6, 2)
	fmt.Println(res)

	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}
