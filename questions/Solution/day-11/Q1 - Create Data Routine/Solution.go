package main

import (
	"fmt"
	"sync"
	"time"
)

func work(id int, wg *sync.WaitGroup) {
	defer wg.Done()             // execute workGroup done at the end of the work function and reduce the counter
	time.Sleep(1 * time.Second) // wait for the spawned routines to execute print statement
	fmt.Printf("Work %d is going on\n", id)
}

func main() {

	var wg = new(sync.WaitGroup)

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go work(i, wg)
	}

	wg.Wait()
	fmt.Println("All work completed.")
}
