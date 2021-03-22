package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var foo string = "bar"
	var baz string = "qux"
	fmt.Println(foo, baz)
	var (
		var1 string  = "Hello!"
		var2 bool    = false
		var3 float64 = 1.23
	)
	fmt.Println(var1, var2, var3)
	var (
		var4 int = 1
		var5     = time.Now()
	)
	fmt.Println(var4, var5)

	var (
		var6 int
	)
	fmt.Println(var6)
	var6 = 1
	fmt.Println(var6)
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
	Debug1, LogLevel1, startUpTime1 := getConfig()
	fmt.Println(Debug1, LogLevel1, startUpTime1)

}

func getConfig() (bool, string, time.Time) {
	return false, "info123", time.Now()
}
