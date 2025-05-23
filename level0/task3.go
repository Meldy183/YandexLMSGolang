package main

import "fmt"

func main() {
	var name string
	var flat int
	var pass int
	var duration float64
	fmt.Scanln(&name)
	fmt.Scanln(&flat)
	fmt.Scanln(&pass)
	fmt.Scanln(&duration)
	fmt.Printf("Привет, %s! Приглашаю тебя на соревнование по программированию, которое пройдёт, как всегда, в квартире %d. Оно будет длиться примерно %.1f часа. Не забудь секретный пароль для входа: %d.", name, flat, duration, pass)
}
