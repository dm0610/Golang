package main

import (
	"fmt"
	"os"
)

var users = map[string]string{
	"301": "Alex",
	"204": "Maya",
	"631": "Nikolay",
	"013": "Kate",
}

func getName(id string) (string, bool) {
	name, exists := users[id]
	return name, exists
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User ID not passed")
		os.Exit(1)
	}
	name, exists := getName(os.Args[1])
	if !exists {
		fmt.Printf("error: user (%v) not found", os.Args[1])
		os.Exit(1)
	}
	fmt.Println("Hi,", name)
}
