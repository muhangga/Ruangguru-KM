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

// TODO: answer here
type Ruang struct {
	Items []Meja
}

type Meja struct {
	Nama  string `json:"nama"`
	Jumlah int    `json:"jumlah"`
	Warna string `json:"warna"`
	Ukuran Ukuran `json:"ukuran"`
}

type Ukuran struct {
	Panjang string `json:"panjang"`
	Tinggi  string `json:"tinggi"`
}

func (r Ruang) EncodeJSON() string {
	// TODO: answer here
	jsonRuang, err := json.Marshal(r.Items)
	if err != nil {
		log.Fatal("JSON Marshal error: ", err)
	}

	return string(jsonRuang)
}

func NewRuang(r Ruang) Ruang {
	return r
}

func main() {
	items := `{
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
	}`

	res := items.EncodeJSON()
	fmt.Println(res)
}