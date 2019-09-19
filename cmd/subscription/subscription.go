package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/sub"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.SubscriptionParams{
		Customer: stripe.String("YOUR_CUSTOMER_ID"),
		Items: []*stripe.SubscriptionItemsParams{
			{
				Plan: stripe.String("gold"),
			},
		},
	}

	s, _ := sub.New(params)
	fmt.Printf("SubscriptionID is %s", s.ID)
}
