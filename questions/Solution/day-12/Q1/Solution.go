package main

import (
	"fmt"
	"sync"
)

var ch = make(chan any)

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(5)
	go Add(1, 2, wg)

	go Subtract(10, 2, wg)

	go Multiply(2, 3, wg)

	go Divide(6, 2, wg)

	go CollectResults(1, 2, wg)

	wg.Wait()
}

func Add(a int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := a + b
	op := fmt.Sprint("Sum = ", sum)
	ch <- op
}

func Subtract(a int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	difference := a - b
	op := fmt.Sprint("Difference = ", difference)
	ch <- op

}

func Multiply(a int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	product := a * b
	op := fmt.Sprint("Product = ", product)
	ch <- op
}

func Divide(a int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	quotient := a / b
	op := fmt.Sprint("quotient = ", quotient)
	ch <- op
}

func CollectResults(a int, b int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 3; i++ {
		fmt.Println("recv picked up ", <-ch)
	}
}
