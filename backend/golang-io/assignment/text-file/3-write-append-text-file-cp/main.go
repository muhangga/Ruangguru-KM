package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	AddString("backend/golang-io/assignment/text-file/3-write-append-text-file-cp/original.txt", "\nThis is a new string to a line")

}

func AddString(fileName string, stringToAdd string) error {

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	defer file.Close()

	if _, err = file.WriteString(stringToAdd); err != nil {
		log.Fatalf("failed writing to file: %s", err)
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("failed reading file: %s", err)
	}

	log.Printf("%s", data)

	return nil // TODO: replace this
}
