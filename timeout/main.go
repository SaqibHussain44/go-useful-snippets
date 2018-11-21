package main

import (
	"fmt"
	"time"
)

func sleep(seconds int, endSignal chan<- bool) {
	time.Sleep(2000 * time.Millisecond)
	endSignal <- true
}

func main() {
	endSignal := make(chan bool, 1)
	go sleep(5, endSignal)

	select {
	case <-endSignal:
		fmt.Println("The end!")
	case <-time.After(2 * time.Second):
		fmt.Println("There's no more time to this. Exiting!")
	}
}
