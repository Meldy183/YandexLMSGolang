package main

import "fmt"

type Person struct {
	name    string
	age     int
	address string
}

func (p Person) PrettyPrint() {
	fmt.Printf("Name: %s\n", p.name)
	fmt.Printf("Age: %d\n", p.age)
	fmt.Printf("Address: %s\n", p.address)
}
