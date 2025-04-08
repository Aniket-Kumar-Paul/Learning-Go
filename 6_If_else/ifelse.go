package main

func main() {
	// Go doesn't have a ternary operator like other languages

	age := 15
	if age >= 18 {
		println("You are an adult.")
	} else if age < 18 && age >= 13 {
		println("You are a teen.")
	} else {
		println("You are a child.")
	}

	// declaring a variable in the if statement
	if age := 20; age >= 18 {
		println("You are an adult.")
	}
}