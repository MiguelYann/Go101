package main

import (
	"fmt"

	"github.com/bitfield/script"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(add(1, 2))

	person := createPersonWithDefaultHobbies(18, "Yann")

	fmt.Println(person.describe())

	executeShellScript()
}

func add(a, b int) int {
	sum := a + b
	return sum
}

type person struct {
	name    string
	age     int
	hobbies [2]string
}

func createPersonWithDefaultHobbies(age int, name string) *person {
	person := person{age: age, name: name, hobbies: [2]string{"Foot", "Basket"}}
	return &person
}

func (p person) describe() string {
	return fmt.Sprintf("My name is %s and I have %d years old", p.name, p.age)
}

func executeShellScript() {
	script.Echo("Hello from Shell").Stdout()
}
