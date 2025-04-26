package auth

import "fmt"

// NOTE: the function name must start with capital letter (automatically denotes export) to be able to use it in other packages
// If it starts with small letter, eg. loginWithCredentials, it can only be used in auth package
func LoginWithCredentials(username string, password string) {
	fmt.Println("Logging in using :", username, password)
}