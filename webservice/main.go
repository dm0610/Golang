package main

import (
	"fmt"

	"github.com/dm0610/Golang/webservice/models"
)

func main() {
	//fmt.Println("Hello from a module webservice, Gophers!")
	u := models.User{
		ID:        2,
		FirstName: "Dmitriy",
		LastName:  "Strelnikov",
	}
	fmt.Println(u)
}
