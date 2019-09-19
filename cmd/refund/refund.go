package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/refund"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	refundParam := &stripe.RefundParams{
		Charge: stripe.String("CHARGE_ID"),
		Reason: stripe.String("requested_by_customer"),
		Amount: stripe.Int64(1000),
	}

	refund, err := refund.New(refundParam)
	if err != nil {
		fmt.Printf("refund error :%v", err)
	}
	fmt.Printf("refund id: %s", refund.ID)
}
