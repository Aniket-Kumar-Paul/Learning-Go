package main

import "fmt"

// Enumerated types (enums) in Go are typically implemented using constants and types combination.
type OrderStatus int
const (
	Received OrderStatus = iota // 0 (default), automatically increments below
	Confirmed // 1
	Prepared // 2
	Delivered // 3
)

// type OrderStatus string
// const (
// 	Received OrderStatus = "received"
// 	Confirmed = "confirmed"
// 	Prepared = "prepared"
// 	Delivered = "delivered"
// )

func changeOrderStatus(status OrderStatus) {
	fmt.Println("Changing order status to ", status)
}

func main() {
	changeOrderStatus(Prepared) // 2
}