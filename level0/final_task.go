package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	layout := "02.01.2006"
	var name, surname, dadName string
	var num1, num2, num3 float64
	var str string
	fmt.Scanln(&str)
	date, _ := time.Parse(layout, str)
	date = date.AddDate(0, 0, 15)
	fmt.Scan(&name, &surname, &dadName, &num1, &num2, &num3)
	ans := num1 + num2 + num3
	frac := int(ans*100) % 100
	fmt.Printf("Уважаемый, %s %s %s, доводим до вашего сведения, что бухгалтерия сформировала документы по факту выполненной вами работы.\nДата подписания договора: %s. Просим вас подойти в офис в любое удобное для вас время в этот день.\nОбщая сумма выплат составит %d руб. %d коп.\n\nС уважением,\nГл. бух. Иванов А.Е. ", surname, name, dadName, date.Format(layout), int(math.Trunc(ans)), frac)
}
