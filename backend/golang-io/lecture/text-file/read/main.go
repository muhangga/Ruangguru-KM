package main

// importing the packages
import (
	"fmt"
	"io/ioutil"
	"log"
)

func readFile() {
	fileName := "backend/golang-io/lecture/text-file/read/read.txt"

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file %s", err)
	}
	fmt.Printf("File Name: %s\n", fileName)
	fmt.Printf("Size %d bytes\n", len(data))
	fmt.Printf("Data: %s\n", data)
}

func main() {
	readFile()
}
//reference : https://www.geeksforgeeks.org/how-to-read-and-write-the-files-in-golang/
