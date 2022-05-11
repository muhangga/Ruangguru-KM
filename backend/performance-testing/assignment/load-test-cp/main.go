package main

import (
	// "encoding/json"
	// TODO: answer here
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

type Movie struct {
	ID      int    `json:"id"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
}

//Baca README untuk tau jumlah request yang perlu dilakukan dan targetnya
//untuk durasi cukup gunakan satu detik

//menambahkan movie baru
//untuk data yang dikirim adalah JSON
//gunakan struct Movie diatas, cukup gunakan field episode dan name
//ID sudah auto increment
func addMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	frequency := 10
	duration := time.Second * 1

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "POST",
		URL:    target + "/movie",
		Body:   []byte(`{"episode": 1, "name": "The Phantom Menace"}`),
	})

	metrics = vegetaAttack(targeter, frequency, duration)
	return metrics
}

//mendapatkan informasi movie dengan ID 1-25
//vegeta.NewStaticTargeter() adalah variadic function
//kita bisa menggunakannya untuk menentukan multiple target vegeta attack
func getMovieTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	frequency := 25
	duration := time.Second * 1

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target + "/movie/1",
	})

	metrics = vegetaAttack(targeter, frequency, duration)
	return metrics
}

//mendapatkan semua informasi movie
func getMoviesTest(target string) *vegeta.Metrics {
	metrics := &vegeta.Metrics{}
	// TODO: answer here
	frequency := 20
	duration := time.Second * 1

	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    target + "/movies",
	})

	metrics = vegetaAttack(targeter, frequency, duration)
	return metrics
}

func vegetaAttack(targeter vegeta.Targeter, frequency int, duration time.Duration) *vegeta.Metrics {
	rate := vegeta.Rate{Freq: frequency, Per: time.Second}
	attacker := vegeta.NewAttacker()
	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Example") {
		metrics.Add(res)
	}
	metrics.Close()
	return &metrics
}
