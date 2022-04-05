package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Hello World")
}

func ScanToArray(arr *[]string, fileName string) error {

	file, err := os.Open(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		*arr = append(*arr, scanner.Text())
	}

	return nil // TODO: replace this
}

func ScanToMap(dataMap map[string]string, fileName string) error {

	file, err := os.Open(fileName)

	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		// fmt.Println(strings.Split(line, ","))
		// fmt.Println(strings.Split(line, ",")[0])
		// fmt.Println(strings.Split(line, ",")[1])
		dataMap[strings.Split(line, ",")[0]] = strings.Split(line, ",")[1]
	}

	return nil // TODO: replace this
}
