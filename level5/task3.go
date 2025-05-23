package main

import (
	"errors"
	"fmt"
	"sort"
)

type CompanyInterface interface {
	AddWorkerInfo(name, position string, salary, experience uint) error
	SortWorkers() ([]string, error)
}

type Company struct {
	name       []string
	position   []string
	salary     []uint
	experience []uint
}

type Interface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

func (c *Company) Len() int {
	return len(c.name)
}
func (c *Company) Less(i, j int) bool {
	mapX := map[string]int{
		"директор":       1,
		"зам. директора": 2,
		"начальник цеха": 3,
		"мастер":         4,
		"рабочий":        5,
	}
	if c.salary[i]*c.experience[i] != c.salary[j]*c.experience[j] {
		return c.salary[i]*c.experience[i] > c.salary[j]*c.experience[j]
	}
	return mapX[c.name[i]] < mapX[c.name[j]]
}
func (c *Company) Swap(i, j int) {
	c.name[i], c.name[j] = c.name[j], c.name[i]
	c.position[i], c.position[j] = c.position[j], c.position[i]
	c.salary[i], c.salary[j] = c.salary[j], c.salary[i]
	c.experience[i], c.experience[j] = c.experience[j], c.experience[i]
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
	if salary < 0 || experience < 0 || name == "" || position == "" {
		return errors.New("salary or experience is invalid")
	}
	c.name = append(c.name, name)
	c.position = append(c.position, position)
	c.salary = append(c.salary, salary)
	c.experience = append(c.experience, experience)
	return nil
}
func (c *Company) SortWorkers() ([]string, error) {
	if len(c.name) < 1 {
		return nil, errors.New("name is empty")
	}
	sort.Sort(c)
	ans := make([]string, 0, len(c.name))
	for i := 0; i < len(c.name); i++ {
		x := fmt.Sprintf("%s — %d — %s", c.name[i], c.salary[i]*c.experience[i], c.position[i])
		ans = append(ans, x)
	}
	return ans, nil
}
