package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// sign dan welcome menggunakan JWT kedalam cookie
func main() {
	// start the server on port 8000
	fmt.Println("Starting Server at port :8000")
	log.Fatal(http.ListenAndServe(":8000", Routes()))
}

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials
		// Task: JSON body diconvert menjadi creditial struct & return bad request ketika terjadi kesalahan decoding json:

		// TODO: answer here
		body := r.Body
		defer body.Close()

		if err := json.NewDecoder(body).Decode(&creds); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		
		// Task: Ambil password dari username yang dipakai untuk login & return unauthorized jika password salah

		// TODO: answer here
		actualPassword := users[creds.Username]
		if creds.Password != actualPassword {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//Task: 1. Deklarasi expiry time untuk token jwt
		// 		2. Buat claim menggunakan variable yang sudah didefinisikan diatas
		//      3. Expiry time menggunakan time millisecond

		// TODO: answer here
		expiredTime := time.Now().Add(time.Millisecond * 15).Unix()
		claims := &Claims{
			Username: creds.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expiredTime,
			},
		}

		//Task: 1. Buat token menggunakan encoded claim dengan salah satu algoritma yang dipakai
		// 		2. Buat jwt string dari token yang sudah dibuat menggunakan JWT key yang telah dideklarasikan
		//      3. return internal error ketika ada kesalahan ketika pembuatan JWT string

		// TODO: answer here
		tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := tokenClaim.SignedString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		//Task: Set token string kedalam cookie response

		// TODO: answer here
		cookie := http.Cookie{
			Name: "token",
			Value: tokenString,
			Expires: time.Now().Add(time.Hour * 24),

		}
		http.SetCookie(w, &cookie)
	})

	mux.HandleFunc("/welcome", func(w http.ResponseWriter, r *http.Request) {
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//		 2. Buat return unauthorized ketika token kosong
		//		 3. Buat return bad request ketika field token tidak ada

		// TODO: answer here
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if cookie.Value == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Task: Ambil value dari cookie token

		// TODO: answer here
		token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		// Task: Deklarasi variable claim

		// TODO: answer here
		claim, ok := token.Claims.(*Claims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		//Task: parse JWT token ke dalam claim

		// TODO: answer here
		

		//Task: return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here

		// Task: return data dalam claim, seperti username yang telah didefinisikan

		// TODO: answer here
	})

	return mux
}
