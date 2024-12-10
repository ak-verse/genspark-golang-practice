package main

import (
	"fmt"
	"sync"
//	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)
	ch := make(chan int)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			wgWorker.Add(1)
			// fan out pattern, spinning up n number of goroutines, for n number of task
			go func() {
				defer wgWorker.Done()
//				fmt.Println("sendin to channel",i)
				ch <- i
//				fmt.Println("sendin to channel",i)

			}()

		}
		// close signal range that no more values be sent and it can stop after receiving remaining values
		// close the channel when sending is finished

		// we can't send more values after a channel is closed
		//ch <- 6 // panic: send on closed channel // channel is closed
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
//		time.Sleep(1 * time.Second)                                                          //AK : stops deadlock
		wgWorker.Wait() // until workers are not finished, we would wait
		// close the channel if workers are done sending
//		fmt.Println("closing channel")
		close(ch)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("inside receiver")                                                       //AK: stops deadlock
		// range gives a guarantee everything would be received
		for v := range ch {
			fmt.Println(v)
		}
	}()

	wg.Wait()
	fmt.Println("Done")

}
