package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var result int
var wg sync.WaitGroup

//fungsi ini digunakan untuk menerima angka lalu mengembalikan hasil pangkat 2 angka tersebut
func squareWorker(workerInput <-chan int, workerOutput chan<- int) {
	for {
		num := <-workerInput      // menerima angka
		workerOutput <- num * num // mengirim hasil
	}
}

func createRequest(workerInput chan<- int, workerOutput <-chan int, wg *sync.WaitGroup) {
	for i := 1; i < 100; i++ {
		// TODO: answer here
		wg.Add(1)		

		go func(i int) {
			// TODO: answer here
			defer wg.Done()

			var res int

			// TODO: answer here
			workerInput <- i // kirim angka ke worker
			res = <-workerOutput // menerima hasil dari worker

			mu.Lock()
			result += res // menjumlahkan hasil
			mu.Unlock()

			//tambahkan res ke result. Selain itu gunakan juga sesuatu yang menghindari data race
			// TODO: answer here
			

			fmt.Println(res)
		}(i)
	}
	wg.Done()
}