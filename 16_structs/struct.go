package main

import (
	"fmt"
	"time"
)

// In go, we don't have classes, but we can use structs to create our own data types.
// By default, all fields are optional
type order struct {
	ID       int
	Amount   float64
	Status   string
	creaedAt time.Time
}

// constructor function (kind of)
func newOrder(id int, amount float64, status string) *order {
	return &order{
		ID:     id,
		Amount: amount,
		Status: status,
	}
}

// (o order) -> this attaches this function to the struct
func (o *order) changeStatus(status string) {
	o.Status = status
	fmt.Println("Order status changed to:", o.Status)
}

func main() {
	myorder := order{
		ID:       1,
		Amount:   100.50,
		Status:   "Pending",
	}
	myorder.creaedAt = time.Now()
	fmt.Println("Order status:", myorder.Status)
	myorder.changeStatus("Shipped")
	fmt.Println("Order status:", myorder.Status)

	order2 := newOrder(2, 200.75, "Processing")
	fmt.Println("Order 2 status:", order2.Status)

	// inline struct 
	language := struct {	
		Name string
		IsGood bool
	}{
		Name: "Go",
		IsGood: true,
	}
	fmt.Println("Language name:", language)
}