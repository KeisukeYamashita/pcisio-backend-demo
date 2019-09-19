  
package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	customerParams := &stripe.CustomerParams{
		Email: stripe.String("piosiki@gmail.com"),
	}
	cus, _ := customer.New(customerParams)
	fmt.Printf("CustomerID is %s", cus.ID)
}