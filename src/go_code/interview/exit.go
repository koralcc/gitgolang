package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int, 1000)
	//endCh := make(chan struct{}, 1)
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	close(ch)
	wg.Add(4)
	for i := 0; i < 4; i++ {

		go func(i int) {
			for {
				val, ok := <-ch
				if !ok {
					//endCh <- struct{}{}
					//return
					break
				}
				fmt.Println("[", strconv.Itoa(i), "]val=", val)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	// for i := 0; i < 4; i++ {
	// 	<-endCh
	// }

}
