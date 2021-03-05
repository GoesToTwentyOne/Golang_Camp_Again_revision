package main

import "fmt"

var i = 25

func main() {
	//if_you_declare_variable_on_package_level_you_declare_that_variable_again_function_level_but_didn%27t_redeclare_in_same_
	//scope_twice
	fmt.Println(i)
	i := 20
	//doesn't redeclare here i:=9
	fmt.Println(i)

}
