package main

import (
	"fmt"
)

// Problem: You are developing a system to manage customer orders for an e-commerce website.
// Each customer order includes an order ID, customer ID, and a timestamp indicating when
// the order was placed. Your goal is to efficiently track and retrieve orders by their IDs
// and quickly retrieve the most recent order for a particular customer.

type Order struct {
	orderId    int
	customerId int
	timestamp  string
}

var orders = make(map[int]Order)       // Map for retrieving orders by ID
var recentOrders = make(map[int]Order) // Map for tracking the most recent order per customer

func main() {
	// Add orders
	addOrder(Order{orderId: 001, customerId: 001, timestamp: "2024-07-28 14:30:00"})
	addOrder(Order{orderId: 006, customerId: 001, timestamp: "2025-07-28 14:30:00"})
	addOrder(Order{orderId: 002, customerId: 002, timestamp: "2023-11-15 09:45:30"})
	addOrder(Order{orderId: 003, customerId: 003, timestamp: "2022-03-22 18:10:45"})
	addOrder(Order{orderId: 004, customerId: 004, timestamp: "2021-08-19 23:59:59"})
	addOrder(Order{orderId: 005, customerId: 005, timestamp: "2024-01-10 07:00:00"})

	// Test retrieval
	fmt.Print("Search By Order Id result:\n", searchByOrderId(002).toString())
	fmt.Print("Most Recent Order For Customer 001:\n", getMostRecentOrderForCustomer(001).toString())
}

// Function to add an order and update the recentOrders map
func addOrder(order Order) {
	// Update the orders map
	orders[order.orderId] = order

	// Check if this order is more recent than the current one for the customer
	if existingOrder, exists := recentOrders[order.customerId]; !exists || order.timestamp > existingOrder.timestamp {
		recentOrders[order.customerId] = order
	}
}

// Function to retrieve an order by its ID
func searchByOrderId(orderId int) Order {
	return orders[orderId]
}

// Function to retrieve the most recent order for a customer
func getMostRecentOrderForCustomer(customerId int) Order {
	return recentOrders[customerId]
}

// Method to format the order information for printing
func (o Order) toString() string {
	return "Order Id: " + fmt.Sprintf("%03d", o.orderId) +
		"\nCustomer Id: " + fmt.Sprintf("%03d", o.customerId) +
		"\nTimestamp: " + o.timestamp + "\n"
}
