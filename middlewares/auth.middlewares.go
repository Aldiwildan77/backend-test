package middlewares

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JwtVerify to verif token
func JwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			responseAuthWithError(w, http.StatusBadRequest, "Authorization required")
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			responseAuthWithError(w, http.StatusBadRequest, "Authorization missing")
			return
		}

		token, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Something wrong")
			}

			return []byte("rahasiayaahahaha"), nil
		})

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok && !token.Valid {
			fmt.Println(err)
			responseAuthWithError(w, http.StatusUnauthorized, "Unauthorized")
		}

		ctx := context.WithValue(r.Context(), "userToken", claims["id"])
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func responseAuthWithError(w http.ResponseWriter, code int, message interface{}) {
	payload := map[string]interface{}{"errors": message}
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
