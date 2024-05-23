package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/anglesson/web-server/middlewares"
)

func main() {
	mainRouter := http.NewServeMux()

	userRouter := http.NewServeMux()
	userRouter.HandleFunc("GET /", getUsers)
	userRouter.HandleFunc("POST /", createUser)

	paymentRouter := http.NewServeMux()
	paymentRouter.HandleFunc("GET /", getPayments)
	paymentRouter.HandleFunc("POST /", createPayments)

	mainRouter.Handle("/user", userRouter)
	mainRouter.Handle("/payment", middlewares.PaymentMiddleware(paymentRouter))

	server := http.Server{
		Addr:    ":8080",
		Handler: middlewares.Logging(mainRouter),
	}

	fmt.Println("Server listening on port 8080")
	server.ListenAndServe()
}

func createUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Created!")
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("User's List!")
}

func createPayments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Payment Created!")
}

func getPayments(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("Payment's List!")
}
