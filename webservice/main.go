package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello to webservice")
	startWebService()
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

func startWebService() {
	fmt.Println("Starting server ...")

	// do important things

	fmt.Println("Server started.")

}
