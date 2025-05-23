package main

import (
	"strings"
)

type Reader interface {
	Read() []byte
}

type Writer interface {
	Write(p []byte) int
}

type ReaderWriter interface {
	Reader
	Writer
}

type UpperReaderWriter struct {
	UpperString string
}

func (tp *UpperReaderWriter) Write(p []byte) int {
	tp.UpperString = strings.ToUpper(string(p))
	return len([]byte(tp.UpperString))
}

func (tp *UpperReaderWriter) Read() []byte {
	return []byte(tp.UpperString)
}
