package main

import (
	"net/http"

	"github.com/dm0610/Golang/webservice/controllers"
)

func main() {
	//fmt.Println("Hello from a module webservice, Gophers!")
	/*	u := models.User{
			ID:        2,
			FirstName: "Dmitriy",
			LastName:  "Strelnikov",
		}
		fmt.Println(u)
	*/
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
