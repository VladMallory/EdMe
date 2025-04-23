package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//https://ru.stackoverflow.com/q/548852
	//rand.Seed(time.Now().UnixNano())
	//генерируем рандомное число
	secret := rand.Intn(10) + 1

	fmt.Println("Угадайте число от 1 до 10!")

	var consoleInput int
	for {
		fmt.Print("Ваш вариант: ")
		fmt.Scanln(&consoleInput)

		if consoleInput == secret {
			fmt.Printf("Вы угадали! Это %d.\n", secret)
			break //на всякий случай выходим из цикла
		} else if consoleInput < secret {
			fmt.Println("Загаданное число больше!")
		} else {
			fmt.Println("Загаданное число меньше!")
		}
	}
}
