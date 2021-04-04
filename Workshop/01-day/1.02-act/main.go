package main

import "fmt"

func main() {
	a, b := 5, 10
	// call swap here
	swap(&a, &b)
	fmt.Println(a, b)
}
func swap(a *int, b *int) {
	*a, *b = *b, *a
}
