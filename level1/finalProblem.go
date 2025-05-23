package main

import (
	"fmt"
)

func err(a string) {
	fmt.Printf("Запись на место номер %s невозможна: место уже занято\n", a)
}
func err2(a string) {
	fmt.Printf("Запись на место номер %s невозможна: некорректный ввод\n", a)
}
func main() {
	a, b, c, d, e := "-", "-", "-", "-", "-"
	free := 5
	for {
		str, num, check := "", "", ""
		fmt.Scanln(&str, &num, &check)
		if str == "" {
			continue
		}
		if str != "" && num == "" {
			switch str {
			case "конец":
				queue(a, b, c, d, e)
				return
			case "количество":
				fmt.Printf("Осталось свободных мест: %d\n", free)
				fmt.Printf("Всего человек в очереди: %d\n", 5-free)
			case "очередь":
				queue(a, b, c, d, e)
			default:
				break
			}
			continue
		}
		if check != "" {
			continue
		}
		if num != "1" && num != "2" && num != "3" && num != "4" && num != "5" {
			err2(num)
			continue
		}
		if free == 0 {
			fmt.Printf("Запись на место номер %s невозможна: очередь переполнена\n", num)
			continue
		}
		switch num {
		case "1":
			if a == "-" {
				a = str
				free--
			} else {
				err(num)
			}
		case "2":
			if b == "-" {
				b = str
				free--
			} else {
				err(num)
			}
		case "3":
			if c == "-" {
				c = str
				free--
			} else {
				err(num)
			}
		case "4":
			if d == "-" {
				d = str
				free--
			} else {
				err(num)
			}
		case "5":
			if e == "-" {
				e = str
				free--
			} else {
				err(num)
			}
		}
	}
}

func queue(a string, b string, c string, d string, e string) {
	fmt.Printf("1. %s\n", a)
	fmt.Printf("2. %s\n", b)
	fmt.Printf("3. %s\n", c)
	fmt.Printf("4. %s\n", d)
	fmt.Printf("5. %s\n", e)
}
