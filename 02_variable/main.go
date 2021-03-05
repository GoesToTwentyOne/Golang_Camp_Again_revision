package main

import (
	"fmt"
	"strconv"
)

//package level vriable
var i int = 25

func main() {
	fmt.Println(i)
	//redeclare the package level variable
	//note:defferent scope declare same name variable
	i := 27
	fmt.Println(i)
	//shadow this variable
	i = 52
	fmt.Println(i)

	//====================================================================================================
	//type conversion
	var j float64
	j = float64(i)
	fmt.Printf("%v %T \n", j, j)
	//when conversion int to string use strconv library as don't risk to value lose in compiler
	s := strconv.Itoa(i)
	fmt.Printf("%v %T \n", s, s)

}
