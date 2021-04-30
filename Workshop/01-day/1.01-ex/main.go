//указатели
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	seeder := time.Now().UnixNano()
	rand.Seed(seeder)
	r := rand.Intn(100) + 1
	var stars string = strings.Repeat("*", r)
	fmt.Println(stars)
}
