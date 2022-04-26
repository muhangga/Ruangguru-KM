package main

import "sync"

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {

	// TODO: answer here
	var wg sync.WaitGroup

	count := 0
	c := make(chan int)

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			//kirim 1 ke channel
			// TODO: answer here
			defer wg.Done()
			c <- 1
		}()
	}
	//mengubah nilai count menggunakan data dari channel

	// TODO: answer here
	go func() {
		for {
			count += <-c
		}
	}()

	wg.Wait() // menunggu seluruh goroutine selesai berjalan
	output <- count
}