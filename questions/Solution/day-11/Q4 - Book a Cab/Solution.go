package main

import (
	"fmt"
	"sync"
)

var cabs int = 5

func main() {
	wg := new(sync.WaitGroup)
	m := new(sync.Mutex)

	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go bookCab(i, wg, m)
	}
	wg.Wait()
}

func bookCab(i int, wg *sync.WaitGroup, m *sync.Mutex) {
	defer wg.Done()
	m.Lock()
	if cabs >= 1{
	fmt.Println("cab allotted to user %d",i)
	cabs = cabs -1
	}else{
		fmt.Println("No cabs available for user %d",i)
	}
	m.Unlock()
}
