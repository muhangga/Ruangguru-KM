package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Dari contoh sebelumnya
// buat JSON string array nested seperti berikut
/*
[
{
		"jenis": "Meja Makan",
		"warna": "Coklat",
		"jumlah": 20,
		"ukuran": {
				"panjang": "50 cm",
				"tinggi": "25 cm"
		}
},
{
		"jenis": "Meja Lipat",
		"warna": "Hitam",
		"jumlah": 1,
		"ukuran": {
				"panjang": "70 cm",
				"tinggi": "30 cm"
		}
}
]
*/

type Ukuran struct {
	// TODO: answer here
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

type Meja struct {
	// TODO: answer here
	Jenis string `json:"jenis"`
	Warna string `json:"warna"`
	Jumlah int    `json:"jumlah"`
	Ukuran Ukuran `json:"ukuran"`
}

type Items struct {
	MejaMeja []Meja
}

func (m Items) EncodeJSON() string {
	// TODO: answer here
	jsonMeja, err := json.Marshal(m)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}
	
	return string(jsonMeja)
}

func NewMeja(m Items) Items {
	return m
}

func main() {
	items := []Meja{
		{
			Jenis:  "Meja Makan",
			Warna:  "Coklat",
			Jumlah: 20,
			Ukuran: Ukuran{
				Panjang: "50 cm",
				Tinggi:  "25 cm",
			},
		},
		{
			Jenis:  "Meja Lipat",
			Warna:  "Hitam",
			Jumlah: 1,
			Ukuran: Ukuran{
				Panjang: "70 cm",
				Tinggi:  "30 cm",
			},
		},
	}

	mejaJSON := EncodeJSON(NewMeja(Items{MejaMeja: items}))
	fmt.Println(mejaJSON)

}