package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Millisecond)
	
	go func() {
		defer wg.Done() // decrement the counter
		fmt.Println("this is anonymous function, sleeping for 100ms ",id)
		time.Sleep(1 * time.Millisecond)	
	}()

	fmt.Printf("Work %d is going on\n", id)
}

func main() {

	var wg = new(sync.WaitGroup)

	for i := 1; i <= 10; i++ {
		wg.Add(2)
		go work(i, wg)
	}

	wg.Wait()
	fmt.Println("All work completed.")
}
