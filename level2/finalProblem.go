package main

import (
	"errors"
	"strings"
	"time"
)

func currentDayOfTheWeek() string {
	now := time.Time{}
	switch now.Weekday() {
	case time.Sunday:
		return "Sunday"
	case time.Monday:
		return "Понедельник"
	case time.Tuesday:
		return "Вторник"
	case time.Wednesday:
		return "Среда"
	case time.Thursday:
		return "Четверг"
	case time.Friday:
		return "Пятница"
	case time.Saturday:
		return "Суббота"
	}
	return "Sunday"
}

func dayOrNight() string {
	now := time.Time{}
	if now.Hour() < 10 || now.Hour() > 22 {
		return "Ночь"
	}
	return "День"
}

func nextFriday() int {
	var now = time.Time{}.Weekday()
	var friday int
	switch now {
	case time.Saturday:
		friday = 7
	case time.Sunday:
		friday = 6
	case time.Monday:
		friday = 5
	case time.Tuesday:
		friday = 4
	case time.Wednesday:
		friday = 3
	case time.Thursday:
		friday = 2
	case time.Friday:
		friday = 1
	}
	return friday - 1
}

func CheckCurrentDayOfTheWeek(answer string) bool {
	now := time.Time{}.Weekday()
	var str string
	switch now {
	case time.Saturday:
		str = "суббота"
	case time.Sunday:
		str = "воскресенье"
	case time.Monday:
		str = "понедельник"
	case time.Tuesday:
		str = "вторник"
	case time.Wednesday:
		str = "среда"
	case time.Thursday:
		str = "четверг"
	case time.Friday:
		str = "пятница"
	}
	if strings.ToLower(answer) == str {
		return true
	}
	return false
}

func CheckNowDayOrNight(answer string) (bool, error) {
	if len(answer) != 8 || (strings.ToLower(answer) != "день" && strings.ToLower(answer) != "ночь") {
		return false, errors.New("исправь свой ответ, а лучше ложись поспать")
	}
	if strings.ToLower(dayOrNight()) == strings.ToLower(answer) {
		return true, nil
	}
	return false, nil
}
