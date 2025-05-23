package main

import (
	"errors"
	"fmt"
)

type UserX struct {
	Name     (string)
	Age      (int)
	IsActive (bool)
}

func NewUser(name string, age int) (*UserX, error) {
	if name == "" {
		name = "name"
		return nil, errors.New(fmt.Sprintf("%s is empty for user", name))
	}
	if age == 0 {
		age = 18
	}
	return &UserX{name, age, true}, nil
}
