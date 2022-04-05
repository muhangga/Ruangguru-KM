package main

import (
	"encoding/csv"
	"fmt"
	// "io"
	"log"
	"os"
)

type Person struct {
	FirstName string
	LastName  string
	Domicile  string
}

func main() {
	// CSVRead()
	data, err := CSVReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, record := range data {
		user := Person{
			FirstName: record[0],
			LastName:  record[1],
			Domicile:  record[2],
		}
		fmt.Println(user)
	}
}

// func CSVRead() {
// 	f, err := os.Open("backend/golang-io/assignment/csv-file/1-read-csv/numbers.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	r := csv.NewReader(f)

// 	for {
// 		record, err := r.Read()
// 		if err == io.EOF {
// 			break
// 		}

// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		for value := range record {
// 			fmt.Printf("%s\n", record[value])
// 		}
// 	}
// }

func CSVReadAll() ([][]string, error) {
	f, err := os.Open("backend/golang-io/assignment/csv-file/1-read-csv/people.csv")
	if err != nil {
		return [][]string{}, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	// skip baris pertama (nama kolom)
	if _, err := r.Read(); err != nil {
		return [][]string{}, err
	}

	data, err := r.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return data, nil
}