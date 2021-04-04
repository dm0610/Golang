package main

import (
	"fmt"
	"time"
)

var (
	LogLevel    = "info"
	startUpTime = time.Now()
	Debug       bool
)

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
