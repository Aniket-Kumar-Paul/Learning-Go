package main

import (
	"fmt"

	"github.com/Aniket-Kumar-Paul/Learning-Go/auth"
	"github.com/Aniket-Kumar-Paul/Learning-Go/user"
)

func main() {
	// First create a module using:
	// > go mod init <module name(usually github repo)>
	// > go get <external package> // install external packages
	// > go mod tidy
			// It adds any missing modules you import in your code to your go.mod file.
			// It removes any modules you imported earlier but are no longer used from go.mod and go.sum
	auth.LoginWithCredentials("aniket", "pass")
	session := auth.GetSession()
	fmt.Println(session)

	user := user.User{
		Email: "ani@gmail.com",
		Name: "John",
	}
	fmt.Println(user)
}