package main

import "fmt"

func main() {
	var a complex64 = 1 + 2i
	var b complex64 = 2 + 5i
	fmt.Printf("%v  %T \n", a, a)
	fmt.Printf("%v  %T \n", b, b)
	fmt.Printf("sum %v \n", a+b)
	fmt.Printf("subtraction %v \n", a-b)
	fmt.Printf("Multiplication %v \n", a*b)
	fmt.Printf("division %v \n", a/b)

}
