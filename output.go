//go has three funtions to output text

//Print() - Prints output with their default formats. We have to \n to print arguments in new line. Spaces = " ". It adds spaces automatically if both are not string
//PrintLn() - Similar to above but spaces are added between arguments and new line at the end of the line.
//Printf - Formats arccoding to a given formatting verb and prints out the statements.

package main
import ("fmt")

func main(){

	var i string = "Hello"
	var j = 15
	var k string = "Hello"

	fmt.Print(i)
	fmt.Print(j)

	fmt.PrintLn(k)

	fmt.Printf("j has a value of: %v and is type: %t \n",j ,j)


}