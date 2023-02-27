// All about go Variables

package main

import ("fmt")

func main(){
	var student1 string = "John"
	var student2 = "Jane"
	x := 2

	fmt.Println(student1)
	fmt.Println(student2)
	fmt.Println(x)

	// Assigning empty variables. It assigns values within

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// := is only used in a function and must be assigned value once declared

	// Multiple Variables

	var a, b, c, d, int = 1, 2, 3 , 4

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	// Of different types

	var a, b = 7, "Michael"

	fmt.Println(a)
	fmt.Println(b)

	//Declaration in a block

	var (
		a int
		b int = 1
		c string = "hello"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}