package main

import (
	"fmt"
	"sync"
)

/*
The program uses the `go` keyword to run two goroutines
concurrently and print a sequence of numbers.

`sync.WaitGroup` to make your program wait for those
goroutines to finish before existing the program.
*/
func main() {
	// Note: if a WaitGroup is explicitly passed
	// into functions, it should be done by pointer.
	var wg sync.WaitGroup
	numberChan := make(chan int)

	for idx := 1; idx <= 3; idx++ {
		wg.Add(1)
		go printNumbers(idx, numberChan, &wg)
	}

	generateNumbers(5, numberChan, &wg)

	// onces a channel is no longer being used it can be
	// closed using built-in `close` function.
	close(numberChan)

	fmt.Println("Waiting for goroutines to finish ...")
	wg.Wait()
	fmt.Println("Done!")
}

// ch chan<- int
// ch is write-only
func generateNumbers(total int, ch chan<- int, wg *sync.WaitGroup) {
	// send data to channel to another goroutine
	for idx := 1; idx <= total; idx++ {
		fmt.Printf("Sending %d to channel\n", idx)
		ch <- idx
	}
}

// ch <-chan int
// ch is read-only
func printNumbers(idx int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	// read data from channel from another goroutine
	for num := range ch {
		fmt.Printf("%d: Read %d from channel\n", idx, num)
	}
}
