package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// Dari contoh sebelumnya tambahkan filter untuk
// menampilkan meja berdasarkan total meja
// key input dari client menggunakan `total`
// contuh URL: http://localhost:8080/table?total=2

type Table struct {
	ID    string `json:"id"`
	Type  string `json:"type"`
	Color string `json:"color"`
	Total int    `json:"total"`
}

func TableHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		// validate methode request
		// logic handler untuk GET request
		if r.URL.Query().Get("total") != "" {
			total, err := strconv.Atoi(r.URL.Query().Get("total"))
			if err != nil {
				log.Fatalf("expected status bad request 400; got %v", http.StatusNotFound)
			}

			// filter data
			var result []Table
			for _, table := range data {
				if table.Total == total {
					result = append(result, table)
				}
			}

			// expected status bad request 400; got 404
			if len(result) == 0 {
				log.Fatalf("expected status bad request 400; got %v", http.StatusNotFound)
			}

			// encode data ke dalam format string JSON
			resultJSON, err := json.Marshal(result)
			if err != nil {
				
				return
			}

			// expected status bad request 400; got 404

	
			// untuk mendaftarkan result sebagai response
			w.Write(resultJSON)
			return
		}

		// TODO: answer here
		http.Error(w, `{"status":"table not found"}`, http.StatusNotFound)
		return
	}

	http.Error(w, "", http.StatusBadRequest)
}

var data = []Table{
	{ID: "T001", Type: "Meja Lipat", Color: "Coklat", Total: 3},
	{ID: "T002", Type: "Meja Belajar", Color: "Hitam", Total: 4},
	{ID: "T003", Type: "Meja Makan", Color: "Coklat", Total: 9},
	{ID: "T004", Type: "Meja Hejau", Color: "Hijau", Total: 3},
}

func main() {
	http.HandleFunc("/table", TableHandler)

	fmt.Println("starting web server at http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}