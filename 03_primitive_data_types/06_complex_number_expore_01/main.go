package main

import "fmt"

func main() {
	var a complex64 = 2 + 1i
	var b complex128 = 2 + 1i
	fmt.Printf("%v %T \n", real(a), real(a))
	fmt.Printf("%v %T \n", imag(a), imag(a))
	// complex64 -> float32
	// complex128 -> float64
	fmt.Printf("%v %T \n", real(b), real(b))
	fmt.Printf("%v %T \n", imag(b), imag(b))
}
