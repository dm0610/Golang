package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello to webservice")
	port := 3000
	isStarted := startWebService(port, 1)
	fmt.Println(isStarted)
	//u := models.User{
	//	ID:        2,
	//	FirstName: "Dmitriy",
	//	LastName:  "Strelnikov",
	//}
	//fmt.Println(u)

	//variables
	//var i int = 1
	//fmt.Println("print i: ", i)
	//j := "String"
	//fmt.Println("print j:", j)

	//pointers

	//var firstName *string = new(string)
	//*firstName = "Dmitriy"
	//fmt.Println(*firstName)
	//lastName := "Strelnikov"
	//lastNamePtr := &lastName
	//fmt.Println("*lastNamePtr: ", *lastNamePtr)
}

func startWebService(port, numberOfRetries int) error {
	fmt.Println("Starting server ...")

	// do important things

	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)

	return errors.New("Something goes wrong")

}
