package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]string) // map[key type]value type
	m["name"] = "John"
	m["age"] = "30"
	fmt.Println(m["name"]) // John
	fmt.Println(m["otherkey"])  // IMP: if key doesn't exist, it returns zero value of the value type
	fmt.Println(len(m)) // 2

	delete(m, "name") // delete key
	fmt.Println(m)

	clear(m) // clear map
	fmt.Println(m) // map[]

	myMap := map[string]int{"price":40, "phones": 3}
	fmt.Println(myMap) // map[phones:3 price:40]

	// Iterating over a map
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	// Check if key exists
	value , ok := myMap["price"]
	if ok {
		fmt.Println("Key exists: ", value)
	} else {
		fmt.Println("Key does not exist")
	}

	// Checking if two maps are equal using the maps package
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"a": 1, "b": 2}
	fmt.Println("Are maps equal? ", maps.Equal(m1, m2)) // true
}