package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	people := []Person{{name: "Bob", age: 21}, {name: "Sam", age: 28}, {name: "Ann", age: 21}, {name: "Joe", age: 22}, {name: "Ben", age: 28}}
	fmt.Println(AgeDistribution(people))
	fmt.Println(FilterByAge(people, 21))
}

func AgeDistribution(people []Person) map[int]int {

	var ageDistribution = make(map[int]int)

	for _, person := range people {
		ageDistribution[person.age]++
	}

	return ageDistribution
}

func FilterByAge(people []Person, age int) []Person {

	var filteredPeople []Person
	for _, person := range people {
		if person.age == age {
			filteredPeople = append(filteredPeople, person)
		}
	}

	return filteredPeople
}
