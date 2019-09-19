package main

import (
	"context"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	kms "cloud.google.com/go/kms/apiv1"
	kmspb "google.golang.org/genproto/googleapis/cloud/kms/v1"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	var err error
	if stripe.Key, err = os.Getenv("ENCODED_STRIPE_SECRET_KEY"); err != nil {
		log.Fatal(err)
	}

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

func decodeKey(ctx context.Context, encodedKey string) (string, error) {

	client, err := kms.NewKeyManagementClient(ctx)
	if err != nil {
		log.Fatalf("client error failed: %v", err)
	}

	t, err := base64.StdEncoding.DecodeString(encodedKey)
	if err != nil {
		log.Fatal("cannot base64")
	}

	// Build the request.
	req := &kmspb.DecryptRequest{
		Name:       os.Getenv("CLOUD_KMS_KEY"),
		Ciphertext: t,
	}
	// Call the API.
	resp, err := client.Decrypt(ctx, req)
	if err != nil {
		log.Fatalf("resp failed: %v", err)
	}

	fmt.Println(string(resp.Plaintext))
}
