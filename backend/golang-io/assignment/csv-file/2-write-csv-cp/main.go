package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Menu struct {
	Name  string
	Price int
}

func WriteToCSV(fileName string, records []Menu) error {
	csvFile, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	// TODO: answer here
	csvWriter := csv.NewWriter(csvFile)
	defer csvWriter.Flush()

	menus := []Menu{
		{"Pizza", 100},
		{"Burger", 200},
		{"Coffee", 300},
		{"Tea", 400},
		{"Sandwich", 500},
	}

	for _, menu := range menus {
		row := []string{menu.Name, strconv.Itoa(menu.Price)}
		if err := csvWriter.Write(row); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	err := WriteToCSV("menu.csv", nil)
	if err != nil {
		panic(err)
	}
}
