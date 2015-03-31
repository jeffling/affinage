package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
)

// globals
var graph *GraphData

// fetches customers and puts into graph
func storeCustomers(data *GraphData) {
	params := &stripe.CustomerListParams{}

	params.Filters.AddFilter("limit", "", "100")
	i := customer.List(params)
	for i.Next() {
		c := i.Customer()
		data.InsertCustomer(c)
	}
}

func init() {
	stripe.Key = AppConfig.StripeKey

	graph = &GraphData{}
	// storeCustomers(graph)
}

func customers(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// returns graph data
	go storeCustomers(graph)

	http.HandleFunc("/", customers)
	fmt.Println("Listening on port", AppConfig.Port)
	err := http.ListenAndServe(":"+AppConfig.Port, nil)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
