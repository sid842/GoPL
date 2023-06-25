package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

var worker_names = map[int]string{
	0: "Alex",
	1: "Bob",
	2: "Cindy",
	3: "Dina",
	4: "Emma",
}

type worker_output struct {
	result      int
	worker_name string
}

func main() {
	// Create channels
	numbers := make(chan int)
	result := make(chan worker_output)
	stop := make(chan struct{})
	input_received := make(chan struct{})

	// Initialise workers
	init_workers(numbers, result, input_received)

	// Open file to store result
	f := open_file()
	defer f.Close()

	// Separate thread to gather input and send to workers
	go get_input(numbers, stop, input_received)

	for {
		stop_listen := false
		select {
		case r := <-result:
			f.WriteString(fmt.Sprintf("Work done by %s Result %d\n", r.worker_name, r.result))
		case <-stop:
			fmt.Printf("Stopping\n")
			stop_listen = true
		}
		if stop_listen {
			break
		}
	}
}

func init_workers(in <-chan int, out chan<- worker_output, input_received chan<- struct{}) {
	num_of_workers := 5

	for i := 0; i < num_of_workers; i++ {
		go func(worker_id int) {
			for {
				num := <-in
				fmt.Printf("Worker %s, input num %d\n", worker_names[worker_id], num)
				input_received <- struct{}{}

				res := perform_work(num)

				out <- worker_output{
					worker_name: worker_names[worker_id],
					result:      res,
				}
			}
		}(i)
	}
	fmt.Println("Workers Initialised")
}

func get_input(numbers chan<- int, stop chan<- struct{}, input_received <-chan struct{}) {
	for {
		var command string
		fmt.Printf("Enter command: ")
		fmt.Scan(&command)
		if command == "stop" {
			fmt.Println("Entered stop, aborting")
			stop <- struct{}{}
			break
		}
		num, err := strconv.Atoi(command)
		if err != nil {
			fmt.Println("Please enter a number")
		} else {
			fmt.Println("Sending Input")
			numbers <- num
			<-input_received
		}
	}
}

func perform_work(num int) int {
	// Perform some heavy work
	time.Sleep(5 * time.Second)
	return num * num
}

func open_file() *os.File {
	f, err := os.OpenFile("./ch8/producer_consumer/text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	return f
}
