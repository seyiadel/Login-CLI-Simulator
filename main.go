package main

import (
	"fmt"
	"learning-go/validation"
)

var modeOfEntry string
var userEmail string
var userPassword string
var database = make([]map[string]string, 0) //Empty List of Maps with 0 as the size of the list (argument)


func login(){
	for{
		fmt.Println("Hi User, Login?(Yes/No): ")
		fmt.Scanln(&modeOfEntry)

		if modeOfEntry == "Yes"{
			fmt.Println("Fill in your details..")
			fmt.Println("Enter a Valid Email: "); fmt.Scan(&userEmail)
			fmt.Println("Enter Password: "); fmt.Scan(&userPassword)
			
			validEmail, validPassword := validation.ValidateUserInput(userEmail, userPassword)

			if validEmail && validPassword{
				fmt.Printf("User Logged In > %v \n", userEmail)	
	
				var userDatabase = make(map[string]string) //Empty Map

				userDatabase["Email"] = userEmail
				userDatabase["Password"] = userPassword
		
				database = append(database, userDatabase)

				fmt.Printf("List of LoggedIn Users %v \n", database)
			}

			
		} else{
			fmt.Println("Bye!")
		}
		}
}



func main(){
	fmt.Println("Welcome to The Store")
	login()
}
	