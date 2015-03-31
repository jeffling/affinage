package main

import "github.com/stripe/stripe-go"

type GraphData struct {
	data           GraphPoints
	lastCustomerId string // use this ID to prevent duplication of data
}

type GraphPoints []GraphPoint

type GraphPoint struct {
	X           int64  `json "x"`
	Delta       int64  `json "delta"`
	Id          string `json "id"`
	Trial_end   int64  `json "trial_end"`
	Description string `json "description"`
}

// AddPoint adds a GraphPoint into the Graph
func (graph *GraphData) AddPoint(g GraphPoint) {
	graph.data = append(graph.data, g)
}

// InsertCustomer takes a customer and inserts points into the graph
// based on subscription information
func (graph *GraphData) InsertCustomer(c *stripe.Customer) {
	if c.Subs.Count > 0 {
		subscription := c.Subs.Values[0]
		graph.AddPoint(GraphPoint{
			c.Created,
			1,
			c.ID,
			subscription.TrialEnd,
			c.Desc,
		})

		if subscription.Canceled > 0 {
			graph.AddPoint(GraphPoint{
				subscription.Canceled,
				-1,
				c.ID,
				subscription.TrialEnd,
				c.Desc,
			})
		}
	}
}
