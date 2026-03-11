package main

import "fmt"

func main() {
	a := 56
	fmt.Println(a)

	b, c, d, _, f := 0, 1, 2, 3, "happiness"
	fmt.Println(b, c, d, f)

	//this would not work...
	/*
		b,c,d,e := 0,1,2,3
		fmt.Println(b,c,d)
	*/
	//this does work....

	var g int
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	//Declare a variable to hold value of a certain type
	//then assign a value of that type to the variable
	//initializing a variable
	var h int = 44
	fmt.Println(h)

}
