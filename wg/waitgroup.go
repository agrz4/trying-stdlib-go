package main

import (
	"fmt"
	"sync"
	"time"
)

// Let's imagine this is our combined Product model
type Product struct {
	ID           int
	Availability string
	Price        float64
}

func main() {
	wg := &sync.WaitGroup{}
	product := &Product{ID: 1}

	// We need to make two calls.
	// It's crucial to call Add for *each* goroutine before starting it.

	// Goroutine to fetch availability
	wg.Add(1)
	go func() {
		// Always defer Done() immediately inside the goroutine.
		defer wg.Done()
		// Simulate network call to a "Products" subgraph
		fmt.Println("Fetching availability for product 1...")
		// Pretend GraphQL query: query { product(id:1) { availability } }
		time.Sleep(100 * time.Millisecond) // Simulate network latency
		product.Availability = "In Stock"
		fmt.Println("Availability fetched.")
	}()

	// Goroutine to fetch price
	wg.Add(1)
	go func() {
		// Defer Done() right away!
		defer wg.Done()
		// Simulate network call to a "Pricing" subgraph
		fmt.Println("Fetching price for product 1...")
		// Pretend GraphQL query: query { product(id:1) { price } }
		time.Sleep(150 * time.Millisecond) // Simulate network latency
		product.Price = 29.99
		fmt.Println("Price fetched.")
	}()

	// Wait for both goroutines to complete
	fmt.Println("Router is waiting for subgraph responses...")
	wg.Wait() // This blocks until the counter hits zero (both Done() calls have happened)

	fmt.Printf("Successfully fetched data for Product ID %d: %+v\n", product.ID, *product)
}
