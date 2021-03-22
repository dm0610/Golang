package main

import "fmt"

func main() {
	// Main course
	var total float64 = 2 * 13
	fmt.Println("Main course:", total)
	// Drinks
	total = total + (4 * 2.25)
	fmt.Println("Drinks:", total)
	// Discount
	total = total - 5
	fmt.Println("Discount:", total)
	// 10% Tip
	tip := total * 0.1
	fmt.Println("10% Tip:", tip)
	total = total + tip
	fmt.Println("Total:", total)
	// Split bill
	split := total / 2
	fmt.Println("Split bill:", split)
	// Reward every 5th visit
	visitCount := 24
	visitCount = visitCount + 1
	remainder := visitCount % 5
	if remainder == 0 {
		fmt.Println("With this visit, you've earned a reward.")
	}

	// concat
	givenName := "John"
	familyName := "Smith"
	fullName := givenName + " " + familyName
	fmt.Println("Hello,", fullName)

	// short operators
	count := 5
	count += 5
	fmt.Println(count)
	count++
	fmt.Println(count)
	count--
	fmt.Println(count)
	count -= 5
	fmt.Println(count)
	name := "John"
	name += " Smith"
	fmt.Println("Hello,", name)

	// compare values
	visits := 15
	fmt.Println("First visit   :", visits == 1)
	fmt.Println("Return visit  :", visits != 1)
	fmt.Println("Silver member   :", visits >= 10 && visits < 21)
	fmt.Println("Gold member   :", visits > 20 && visits <= 30)
	fmt.Println("Platinum member :", visits > 30)
}
