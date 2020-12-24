/*Сгенерировать случайное число от 1 до 100 и сохранить его.
Предложить игроку угадать задуманное число и сохранить его ответ.
Если предположение игрока меньше загаданного числа, вывести сообщение (допустим, «Oops. Your guess was LOW»). Если оценка игрока больше загаданного числа, вывести другое сообщение («Oops. Your guess was HIGH»).
Разрешить игроку угадывать до 10 раз. Перед каждым предположением сообщить игроку, сколько попыток у него осталось.
Если предположение игрока равно загаданному числу, вывести сообщение об успехе и перестать запрашивать предположения.
Если у игрока кончились попытки, а число так и не было угадано, вывести сообщение: «Sorry. You didn’t guess my number. It was: [загаданное
число]».*/
// guess - игра, в которой игрок должен угадать случайное число.
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	reader := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		fmt.Print("Make a guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your guess was HIGH.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't guess my number. It was:", target)
	}
}
