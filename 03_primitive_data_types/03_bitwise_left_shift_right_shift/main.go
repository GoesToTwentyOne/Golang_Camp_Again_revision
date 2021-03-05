package main

import "fmt"

func main() {
	a := 2              //0000 0010
	fmt.Println(a << 3) //left shift           00010000 =16
	fmt.Println(a >> 1) //right shift          0000 0001 =1

}
