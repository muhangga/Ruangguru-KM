package main

import (
	"bytes"
	"fmt"
	"html/template"
)

// Dari contoh yang telah diberikan, kamu dapat mencoba untuk menggunakan function pada template.
// Lengkapi function CalculateScore yang berfungsi untuk menjumlahkan total score dari users
// Lengkapi function ExecuteToByteBuffer dan textTemplate sehingga template dapat mencetak objek Leaderboard dengan ketentuan:
// Lakukan looping untuk setiap user
// Pada setiap loop, cetak "Nama: Score", contoh: "Roger: 1000"
// Pada bagian terakhir, cetak "Total Score: total", contoh: "Total Score: 50000"

type UserRank struct {
	Name  string
	Email string
	Rank  int
	Score int
}

type Leaderboard struct {
	Users []*UserRank
}

func CalculateScore(leaderboard Leaderboard) int {
	// TODO: answer here
	var total int

	for _, user := range leaderboard.Users {
		total += user.Score
	}

	return total
}

func ExecuteToByteBuffer(leaderboard Leaderboard) ([]byte, error) {
	var textTemplate string
	// TODO: answer here
	textTemplate = "{{range .Users}}" +
		"{{.Name}}: " +
		"{{.Score}}" +
		"{{end}}" +
		"Total Score: {{CalculateScore .}}"

	funcMap := template.FuncMap{
		"CalculateScore": CalculateScore,
	}

	tmpl, err := template.New("test").Funcs(funcMap).Parse(textTemplate)

	if err != nil {
		return nil, err
	}

	var b bytes.Buffer
	err = tmpl.Execute(&b, leaderboard)

	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func main() {
	leaderboard := Leaderboard{
		Users: []*UserRank{
			{
				Name:  "Roger",
				Email: "angga@gmail.com",
				Rank:  1,
				Score: 1000,
			},
			{
				Name:  "Sandy",
				Email: "angga@gmail.com",
				Rank:  4,
				Score: 4000,
			},
			{
				Name:  "Sandy",
				Email: "angga@gmail.com",
				Rank:  4,
				Score: 4000,
			},
		},
	}

	b, err := ExecuteToByteBuffer(leaderboard)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
