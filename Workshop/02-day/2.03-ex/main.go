//Использование инструкции else if
package main

import "fmt"

func main() {
	input_var := -10
	if input_var < 0 {
		fmt.Println("input_var can't be a negative number")
	} else if input_var%2 == 0 {
		fmt.Println(input_var, "is even")
	} else {
		fmt.Println(input_var, "is odd")
	}
}
