package main

import "fmt"

var baz string = "qux" //declaring on package level

func main() {
	var foo string = "bar" //declaring on main fung level
	fmt.Println(foo, baz)
}
