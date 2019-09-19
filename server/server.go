package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		type HealthResponse struct {
			message string
		}

		json, err := json.Marshal(&HealthResponse{message: "I am healthy"})
		if err != nil {
			log.Fatal(err)
		}
		w.Write(json)
		return
	})

	http.HandleFunc("/api/v1.0/charge", func(w http.ResponseWriter, r *http.Request) {
		keys, _ := r.URL.Query()["userD"]
		userID := keys[0]

		results, err := db.Query("SELECT * FROM Users ( 2, 'TEST' )", userID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(nil)
			return
		}

		var cusID string

		for results.Next() {
			err = results.Scan(&cusID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write(nil)
				return
			}
		}

		type PaymentResponse struct {
			status int
		}

		chargeParams := &stripe.ChargeParams{
			Amount:   stripe.Int64(1000),
			Currency: stripe.String("jpy"),
			Customer: stripe.String(cusID),
		}

		charge.New(chargeParams)

		json, _ := json.Marshal(&PaymentResponse{status: 200})
		w.Write(json)
		return
	})

	log.Printf("server started on port 5050")
	if err := http.ListenAndServe(":5050", nil); err != nil {
		log.Fatalf("start server error: %v", err)
	}
}
