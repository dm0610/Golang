package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Tell me your name")
	reader := bufio.NewReader(os.Stdin)
	myname, _ := reader.ReadString('\n')
	fmt.Println("My name is", myname)

}
