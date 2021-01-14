// average вычисляет среднее значение.
package main

import (
	"fmt"
	"log"

	"github.com/dm0610/Golang/100-days-of-go/06-day/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("e:\\MyProject\\Golang\\100-days-of-go\\05-day\\readfile\\data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}
