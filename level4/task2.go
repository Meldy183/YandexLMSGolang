package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   float64
	bonus    float64
}

func (p Employee) CalculateTotalSalary() {
	fmt.Printf("Employee: %s\n", p.name)
	fmt.Printf("Position: %s\n", p.position)
	fmt.Printf("Total Salary: %.2f\n", p.salary+p.bonus)
}
