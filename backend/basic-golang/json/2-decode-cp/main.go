package main

import (
	"encoding/json"
	"fmt"
)

// Dari contoh yang telah diberikan, cobalah untuk melakukan decode dari json menjadi objek dari struct.
// Mengambil kasus pada encode, lengkapi function DecodeToLeaderboard agar json dapat di decode menjadi objek Leaderboard

// type UserRank struct {
// 	Name  string
// 	Email string
// 	Rank  int
// }

type UserRank struct {
	// TODO: answer here
	Name  string `json: "name"`
	Email string `json: "-"`
	Rank  int    `json: "rank"`
}

type Leaderboard struct {
	Users []*UserRank
}

func DecodeToLeaderboard(jsonData []byte) (Leaderboard, error) {
	// TODO: answer here

	leaderBoard := Leaderboard{}
	err := json.Unmarshal(jsonData, &leaderBoard)

	if err != nil {
		return leaderBoard, err
	}

	
	return leaderBoard, nil
}

func main() {
	jsonData := []byte(`{
		"Users":[
		   {
			  "Name":"Roger",
			  "Email":"roger@ruangguru.com",
			  "Rank":1
		   },
		   {
			  "Name":"Tony",
			  "Email":"tony@ruangguru.com",
			  "Rank":2
		   }
		]
	 }`)

	leaderBoard, err := DecodeToLeaderboard(jsonData)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(leaderBoard)
}
