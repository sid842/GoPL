package main

import (
	"fmt"
)

func main() {
	// pipe1()
	pipe2()
}

func pipe2() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

func counter(out chan<- int) {
	for x := 0; x < 10; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for x := range in {
		out <- x * x
	}
	close(out)
}

func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

/*
func pipe1() {
	naturals := make(chan int)
	squares := make(chan int)

	//counter
	go func() {
		for x := 0; x < 10; x++ {
			naturals <- x
			// time.Sleep(1 * time.Second)
		}
		close(naturals)
	}()

	//squarer
	go func() {
		for x := range naturals {
			// x, ok := <-naturals
			// if !ok {
			// 	break
			// }
			squares <- x * x
		}
		close(squares)
	}()

	//printer
	for x := range squares {
		fmt.Println(x)
	}
}
*/
