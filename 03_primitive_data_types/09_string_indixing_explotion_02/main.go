package main

import "fmt"

func main() {
	s := "Hello! this is me"
	s1 := []byte(s)
	fmt.Println(s1)
	fmt.Println(string(s1))
}
