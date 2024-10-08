package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	sum := 0
	var wg sync.WaitGroup
	//lock := make(chan bool)
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			mutex.Lock()
			sum = sum + 1
			mutex.Unlock()
			wg.Done()
			//lock <- true
		}()
		//_ = <-lock
	}

	wg.Wait()
	fmt.Println(sum)
}
