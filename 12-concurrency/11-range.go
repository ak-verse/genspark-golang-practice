package main

import (
	"fmt"
	"sync"
//	"time"
)

//******

// Don't use this solution in real life, this is wrong one

// ********
func main() {

	wg := new(sync.WaitGroup)
	wgWorker := new(sync.WaitGroup)
	ch := make(chan int, 0)
	//m := &sync.Mutex{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		for i := 1; i <= 5; i++ {
			wgWorker.Add(1)

			// fan out pattern, spinning up n number of goroutines, for n number of task
			go func(i int) {
				defer wgWorker.Done()
//				fmt.Println("sendin to channel",i)
				ch <- i
//				fmt.Println("sendin to channel",i)

			}()

			}(i)
			time.Sleep(1 * time.Second)
		}

	}()

	// separate goroutine in main should not be created to tack workers
	// possibility of wgworker value is 0
	// maybe workers never got chace to run
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

}
