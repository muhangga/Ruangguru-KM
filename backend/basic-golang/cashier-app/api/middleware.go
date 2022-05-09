package api

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	// "strings"
	// "github.com/dgrijalva/jwt-go"
)

func (api *API) AllowOrigin(w http.ResponseWriter, req *http.Request) {
	// localhost:9000 origin mendapat ijin akses
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	// semua method diperbolehkan masuk
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// semua header diperbolehkan untuk disisipkan
	w.Header().Set("Access-Control-Allow-Headers", "*")
	// allow cookie
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	}
}

func (api *API) AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		api.AllowOrigin(w, r)
		encoder := json.NewEncoder(w)
		// Task: 1. Ambil token dari cookie yang dikirim ketika request
		//       2. return unauthorized ketika token kosong
		//       3. return bad request ketika field token tidak ada

		// TODO: answer here
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: "Unauthorized"})
			return
		}
		if authorizationHeader == "" {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: "Unauthorized"})
			return
		}

		// TODO: answer here
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		
		// Task: Deklarasi variable claim

		// TODO: answer here
		claims := Claims{}

		// Task: 1. parse JWT token ke dalam claim
		//       2. return unauthorized ketika signature invalid
		//       3. return bad request ketika field token tidak ada
		//       4. return unauthorized ketika token sudah tidak valid (biasanya karna token expired)

		// TODO: answer here
		token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: "Unauthorized"})
			} else if method != jwt.SigningMethodHS256 {
				w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: "Unauthorized"})
			}
			return jwtKey, nil
		})

		// Task: Validasi
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			encoder.Encode(AuthErrorResponse{Error: "Bad Request"})
			return
		}

		claimsToken, ok := token.Claims.(*jwt.MapClaims)
		if !ok || !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			encoder.Encode(AuthErrorResponse{Error: "Unauthorized"})
			return
		}

		ctx := context.WithValue(r.Context(), "token", claimsToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
