package main

import "fmt"

type Logger interface {
	Log(message string)
}

type LogLevel string

const (
	Error LogLevel = "ERROR"
	Info  LogLevel = "INFO"
)

type Log struct {
	Level LogLevel
}

func (l *Log) Log(message string) {
	switch l.Level {
	case Error:
		fmt.Printf("%s: %s\n", l.Level, message)
	case Info:
		fmt.Printf("%s: %s\n", l.Level, message)
	default:
		panic("Unknow log level")
	}
}
