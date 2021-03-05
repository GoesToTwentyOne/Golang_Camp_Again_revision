package main

import "fmt"

func main() {
	a := 2 // 0000 0010
	b := 5 //0000 0101
	fmt.Println("a&b bitwise operator : ", a&b)
	//0010
	//0101
	//answer 0
	fmt.Println("a | b bitwise operator : ", a|b)
	//0010
	//0101
	//answer 7
	fmt.Println("a ^ b bitwise operator : ", a^b)
	//0010
	//0101
	//answer 7

}
