package main

import (
	"fmt"
	"github.com/kr/pretty"
)

type Person struct{
	firstName string
	lastName  string
	age       int
}

func (p Person) showNames() {
	fmt.Println(p.firstName+" "+p.lastName)
}

func (p *Person) incAge() {
	p.age += 1
}

func main() {
	p := Person{firstName: "Jose", lastName: "Abraham"}
	_, _ = pretty.Println(p)

	p.incAge()
	_, _ = pretty.Println(p)
}
