package main

import "fmt"

const pi = 3.14

func main() {
	const name string = "John Doe"
	fmt.Println("Pi:", pi)

	const (
		host = "localhost"
		port = 8080
	)
	fmt.Println(host, port)
}