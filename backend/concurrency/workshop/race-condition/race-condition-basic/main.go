package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mu := &sync.Mutex{}
	counter := 0

	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			counter += 1
			mu.Unlock()
		}()

		go func() {
			mu.Lock()
			counter += 2
			mu.Unlock()
		}()
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println(counter)
}
