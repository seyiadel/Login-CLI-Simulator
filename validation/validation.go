package validation

import (
	"fmt"
	"strings"
)

func ValidateUserInput(email string, password string) (bool, bool){
	var validEmail = strings.Contains(email, "@")
	var validPassword = len(password) >= 8

	if !validEmail && !validPassword{
		fmt.Println("############# ?>")
		fmt.Println("Validation Error: Password must equal to or more than 8 characters ..")
		fmt.Println("Validation Error: Invalid Email ..")
	}

	return validEmail, validPassword
}