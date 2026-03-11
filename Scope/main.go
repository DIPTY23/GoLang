package main

//the imported package "fmt" is in "file block" scope

import (
	furtherexplored "basics/Scope/furtherexplored"
	"fmt"
)

// x is in package block
var x int = 89

func main() {
	fmt.Println(x) //x can be accessed here

	printMe() //x can be accessed here also

	//y does not exist here yet so this won't work
	//fmt.Println(y)

	//y is in "block scope"
	y := 67
	fmt.Println(y)

	p1 := person{
		"James",
	}
	p1.sayHello()

	//variable shadowing
	x := 34
	fmt.Println(x)

	printMe() // package block x is still the same
	furtherexplored.Fascinating()

	if z := 82; z >= 50 {
		fmt.Println(z)
	}
	//you can't express z here , look at the "comma ok idiom"

}
func printMe() {
	fmt.Println(x)
}

type person struct {
	first string
}

// the method sayHello which is attached to values of TYPE PERSON
// is in package block scope
func (p person) sayHello() {
	fmt.Println("hi my name is", p.first)
}
