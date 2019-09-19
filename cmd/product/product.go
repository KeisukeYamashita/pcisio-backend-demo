package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/product"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")

	params := &stripe.ProductParams{
		Name: stripe.String("My SaaS Platform"),
		Type: stripe.String(string(stripe.ProductTypeService)),
	}
	prod, _ := product.New(params)
	fmt.Printf("Product is %s", prod.ID)
}
