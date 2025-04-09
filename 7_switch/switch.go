package main

import "time"

func main() {
	i := 2

	// single switch
	switch i { // no need to write break after every case like in C/C++
	case 1:
		println("i is 1")
	case 2:
		println("i is 2")
	case 3:
		println("i is 3")
	default:
		println("i is not 1, 2 or 3")
	}

	// multiple cases
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		println("It's the weekend")

	default:
		println("It's a weekday")
	}

	// type switch
	whatsMyType := func(i interface{}) { // interface{} is the empty interface, which can hold any type
		switch i.(type) {
		case int:
			println("int")
		case string:
			println("string")
		default:
			println("unknown type")
		}
	}
	whatsMyType("1")
}
