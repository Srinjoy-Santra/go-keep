package middleware

import (
	"log"
	"net/http"
	"strings"
)

func EnsureAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Checking if user is admin")
		if !strings.Contains(r.Cookie("Authorization"), "Admin") {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(http.StatusText(http.StatusUnauthorized)))
			return
		}
		next.ServeHTTP(w, r)
	})
}
