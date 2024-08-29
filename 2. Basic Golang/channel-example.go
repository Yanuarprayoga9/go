package main

import (
	"fmt"
	"time"
)

// Order represents a customer order
type Order struct {
	ID     int
	Status string
}

// processOrder simulates order processing and sends status through a channel
func processOrder(orderID int, statusChannel chan Order) {
	// Simulating some processing time
	fmt.Println("processing order",orderID)
	time.Sleep(time.Second * 2)
	// Sending processed order status back through the channel
	statusChannel <- Order{ID: orderID, Status: "Completed"}
}

func main() {
	// Creating a channel for communicating order statuses
	statusChannel := make(chan Order)

	// Starting goroutines for processing orders
	for i := 1; i <= 5; i++ {
		go processOrder(i, statusChannel)
	}

	fmt.Println("test")
	// Receiving and printing order statuses
	for i := 1; i <= 5; i++ {
		processedOrder := <-statusChannel
		fmt.Printf("Order %d processed: %s\n", processedOrder.ID, processedOrder.Status)
	}
}
