package middlewares

import (
	"log"
	"net/http"
)

func Logging(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func PaymentMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Logging Payment")
		next.ServeHTTP(w, r)
	}
}
