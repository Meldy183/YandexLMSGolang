package main

import (
	"fmt"
)

type AnimalSecond interface {
	MakeSound() string
	GetName() string
	GetInfo() string
}

type animal struct {
	name    string
	age     int
	species string
	sound   string
	Animal
}

func (a animal) MakeSound() string {
	return a.sound
}
func (a animal) GetName() string {
	return a.name
}
func (a animal) GetInfo() string {
	return fmt.Sprintf("Имя: %s, Вид: %s, Возраст: %d", a.name, a.species, a.age)
}
func NewAnimal(name, species string, age int, sound string) AnimalSecond {
	return &animal{name: name, species: species, age: age, sound: sound}
}
func ZooShow(animals []AnimalSecond) {
	for _, animal := range animals {
		fmt.Println(animal.GetInfo())
		fmt.Println(animal.MakeSound())
	}

}

type ZooKeeper struct {
}

func (z ZooKeeper) Feed(a AnimalSecond) {
	fmt.Printf("Смотритель зоопарка кормит %s. %s!", a.GetName(), a.MakeSound())
}
