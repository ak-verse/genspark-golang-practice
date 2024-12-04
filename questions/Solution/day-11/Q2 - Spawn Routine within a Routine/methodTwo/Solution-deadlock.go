package main

import (
	"fmt"
	"sync"
	"time"
)

func workDup(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(1 * time.Millisecond)
	

	wg.Add(1)
	go func() {
		defer wg.Done() // decrement the counter
		fmt.Println("this is anonymous function, sleeping for 100ms ",id)
		time.Sleep(1 * time.Millisecond)	
	}()

	fmt.Printf("Work %d is going on\n", id)
	wg.Wait();
}

func mainDup() {

	var wg = new(sync.WaitGroup)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workDup(i, wg)
	}

	wg.Wait()
	fmt.Println("All work completed.")
}
