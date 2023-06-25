package main

import (
	"fmt"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		var command string
		for {
			fmt.Scan(&command)
			if command == "cancel" {
				fmt.Println("Entered cancel, aborting")
				break
			}
		}
		abort <- struct{}{}
	}()

	fmt.Println("Timer started")
	select {
	case <-time.After(10 * time.Second):
		fmt.Println("Time elapsed, ending the function")
	case <-abort:
		fmt.Println("Function was aborted")
	}
}
