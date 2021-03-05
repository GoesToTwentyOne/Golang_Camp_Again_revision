package main

import "fmt"

func main() {
	b := 64             //      0010 0000
	fmt.Println(b << 2) // 1000 0000 =256
	fmt.Println(b >> 2) //0000 1000 =16

}
