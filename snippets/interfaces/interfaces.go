package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name string
	Age  int
}

type PersonError struct {
	When time.Time
	What string
}

func (e *PersonError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func (p Person) String() (string, error) {
	if p.Name == "" || p.Age == 0 {
		return "", &PersonError{
			time.Now(),
			"Person missing details",
		}
	}
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age), nil
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	y := Person{"", 9001}
	fmt.Println(a, z, y)
	message, ok := a.String()
	if ok == nil {
		fmt.Println(message)
	}
	fmt.Println(y.String())
}
