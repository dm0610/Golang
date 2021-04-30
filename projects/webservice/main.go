package main

import (
	"fmt"

	"github.com/dm0610/Golang/projects/webservice/models"
)

func main() {

	u := models.User{
		ID:        2,
		FirstName: "Dmitry",
		LastName:  "Strelnikov",
	}
	port := 3000
	fmt.Println(u)
	_, err := startWebServer(port)
	if err != nil {
		print(err)
	}
}

func startWebServer(port int) (int, error) {
	fmt.Println("Starting server...")
	//do important things
	fmt.Println("Server started on port", port)
	return port, nil
}
