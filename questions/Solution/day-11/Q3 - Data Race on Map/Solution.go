package main

import (
	"fmt"
	"sync"
)

var user map[int]int = make(map[int]int, 5)

func main1() {
	wg := new(sync.WaitGroup)
	//m := new(sync.Mutex)
	//wg.Add(10)
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go updateMap(i, wg)
	}
	wg.Wait()
}

func updateMap(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	user[i] = i * i
	fmt.Println(user)
}
