package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	source := &stripe.SourceParams{
		Token: stripe.String("tok_visa"),
	}

	chargeParams := &stripe.ChargeParams{
		Source:   source,
		Currency: stripe.String("jpy"),
		Amount:   stripe.Int64(1000),
	}

	charge, _ := charge.New(chargeParams)
	fmt.Printf("charge id: %s", charge.ID)
}
