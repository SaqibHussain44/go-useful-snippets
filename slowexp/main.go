package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Program started...")
	start := time.Now()

	// Initialise top-level context
	ctx := context.Background()

	operationWithTimeout(ctx)

	elapsed := time.Since(start)
	fmt.Printf("Program finished... It took %s \n", elapsed)
}

func operationWithTimeout(ctx context.Context) {
	fmt.Println("operationWithTimeout started...")

	// Create a channel for signal handling
	c := make(chan bool)

	// Define a cancellation after 1s in the context
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	// Run slowOperation via a goroutine
	go func() {
		slowOperation(c)
	}()

	// Listening to signals
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	case <-c:
		fmt.Println("Unexpected success!")
	}

	fmt.Println("operationWithTimeout finished...")
}

func slowOperation(c chan bool) {
	fmt.Println("slowOperation started...")

	time.Sleep(3 * time.Second)

	c <- true

	fmt.Println("slowOperation finished...")
}
