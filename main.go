package main

import (
	"fmt"
)

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I'm a secret agent - this is my name:", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I'm a secret agent - this is my name:", sa.first, "| License to kill:", sa.ltk)
}

// any TYPE that has the methods specified by an interface is also of the interface type
// an interface says:
// "If you have these methods, then you're also my type"

type human interface {
	speak()
}


func main() {
	p1 := person{
		first: "Miss Moneypenny",
	}
	
	sa1 := secretAgent{
		person: person{
			first: "James",
		},
		ltk: true,
	}

	fmt.Printf("%T\n", p1)

	// in go a VALUE can be of more than one type
	// in this example, p1 is both TYPE person and TYPE human
	var x, y human
	x = p1
	y = sa1
	x.speak()
	y.speak()
	// fmt.Printf("%T\n", x)
}
