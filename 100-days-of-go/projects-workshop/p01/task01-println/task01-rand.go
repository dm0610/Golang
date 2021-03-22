package main

import (
	//не надо сразу добавлять вручную импорты - они добавляются нормально при go format. И вообще
	// при компиляции потом может всплыть ошибка, что два импорта вызываются для одной штуки
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) //засеиваем рандомайзер
	r := rand.Intn(5) + 1            //вызов генератора случайных чисел
	stars := strings.Repeat("*", r)  //создаём репитер строк
	fmt.Println(stars)
}
