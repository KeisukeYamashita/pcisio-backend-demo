package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/plan"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	params := &stripe.PlanParams{
		Amount:    stripe.Int64(3000),
		Currency:  stripe.String(string(stripe.CurrencyJPY)),
		Interval:  stripe.String(string(stripe.PlanIntervalMonth)),
		Nickname:  stripe.String("Pro Plan"),
		ProductID: stripe.String("prod_CHxGUqw1dyKsDM"),
	}
	plan, _ := plan.New(params)
	fmt.Printf("Product is %s", plan.ID)
}
