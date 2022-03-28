package main

import (
	"fmt"
)

// Birthday Cake Candles
// - Problem Solving (Basic)
// - Difficulty: Easy

// Full Problem: https://www.hackerrank.com/challenges/birthday-cake-candles/problem

func birthdayCakeCandles(candles []int32) int32 {
	// TODO: answer here
	var n int32
    max := candles[0]
    
    for _, num := range candles {
        if num > max {
            max = num
        }
    }
    
    for _, num := range candles {
        if num == max {
            n++
        }
    }
    return n
}

func main() {
	var arr = []int32{3, 2, 1, 3}
	fmt.Println(birthdayCakeCandles(arr))
}
