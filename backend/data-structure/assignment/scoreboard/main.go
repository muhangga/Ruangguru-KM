package main

import (
	"fmt"
	"sort"
)

type Score struct {
	Name    string
	Correct int
	Wrong   int
	Empty   int
}
type Scores []Score

func (s Scores) Len() int {
	return len(s)
}

func (s Scores) Less(i, j int) bool {

	if s[i].Correct*4-s[i].Wrong == s[j].Correct*4-s[j].Wrong {
		if s[i].Empty < s[j].Empty {
			return true
		} else {
			return false
		}
	} 
	
	if s[i].Correct*4-s[i].Wrong > s[j].Correct*4-s[j].Wrong {
		return true
	} else {
		return false
	}
}

func (s Scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Scores) TopStudents() []string {
	sort.Sort(s)
	var result []string

	for _, score := range s {
		result = append(result, score.Name)
	}
	return result
}

func main() {
	scores := Scores([]Score{
		{"Levi", 3, 2, 2},
		{"Agus", 3, 4, 0},
		{"Doni", 3, 0, 7},
		{"Ega", 3, 0, 7},
		{"Anton", 2, 0, 5},
	})

	sort.Sort(scores)
	fmt.Println(scores.TopStudents())
	fmt.Println(scores)
}
