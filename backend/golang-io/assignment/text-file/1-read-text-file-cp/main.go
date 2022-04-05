package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// fmt.Println("dummyCommit")

	fileData, err := ReadFile("backend/golang-io/assignment/text-file/1-read-text-file-cp/read.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileData.Data))
	fmt.Println(fileData.Name)
	fmt.Println(fileData.Size)
}

// Gunakan struct untuk menyimpan data file nya ya
type FileData struct {
	Name string
	Size int
	Data []byte
}

func ReadFile(name string) (FileData, error) {

	// membaca text file
	data, err := ioutil.ReadFile(name)
	if err != nil {
		return FileData{}, err
	}

	// membuat struct untuk menyimpan data file
	fileData := FileData{
		Name: name,
		Size: len(data),
		Data: data,
	}

	return fileData, nil
}
