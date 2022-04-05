package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)


func CSVToMap(data map[string]string, fileName string) (map[string]string, error) {

	f, err := os.Open(fileName)

	if err != nil {
		return data, err
	}

	defer f.Close()

	r := csv.NewReader(f)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return data, err
		}

		data[record[0]] = record[1]
	}

	return data, nil
}

func main() {

	data := make(map[string]string)

	CSVToMap(data, "backend/golang-io/assignment/csv-file/1-read-csv-cp/questions.csv")

	for key, value := range data {
		fmt.Println(key, value)
	}

}