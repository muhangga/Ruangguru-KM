package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// buat JSON string array nested seperti berikut
/*
{
	"ruangTamu": {
			"items": [
					{
							"nama": "Meja",
							"jumlah": 20,
							"warna": "Coklat",
							"ukuran": {
									"panjang": "50 cm",
									"tinggi": "25 cm"
							}
					},
					{
							"nama": "Kursi",
							"jumlah": 1,
							"warna": "Hitam",
							"ukuran": {
									"panjang": "70 cm",
									"tinggi": "30 cm"
							}
					}
			]
	}
}
*/

type Ruang struct {
	RuangTamu Items `json:"ruangTamu"`
}

// TODO: answer here
type Items struct {
	Items []Item `json:"items"`
}

type Item struct {
	Nama   string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna  string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here

	jsonRuang, err := json.Marshal(r)

	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}

	jsonData := string(jsonRuang)
	return jsonData
}

func NewRuang(r Ruang) Ruang {
	return r
}

func main() {

	ruangTamu := Ruang{
		Items{[]Item{
			{
				Nama:   "Meja",
				Jumlah: 20,
				Warna:  "Coklat",
				Ukuran: Ukuran{
					Panjang: "50 cm",
					Tinggi:  "25 cm",
				},
			},
			{
				Nama:   "Meja Lipat",
				Warna:  "Hitam",
				Jumlah: 1,
				Ukuran: Ukuran{
					Panjang: "70 cm",
					Tinggi:  "30 cm",
				},
			},
		}},
	}

	fmt.Println(ruangTamu.EncodeJSON())
}
