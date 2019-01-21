package main

import "fmt"

//Declaration of variabls outside a function have to follow this format
//You can not use the short format
var x = 42
var y = "James BondTwo"
var z = true

func main(){
	//Using the short declaration syntax, I can assign 3 variables with different type in one line
	a, b, c := 32, "James Bond", true
	fmt.Println(a, b, c)

	fmt.Printf("Hello, my name is %v and I am %v years old.\t This is all %v\n", b, a, c)

	//Used the  S print method which outputs the data as string
	//allowing us to save it as string inside the variable s
	s := fmt.Sprintf("Hello %v %v %v", x,y,z)
	fmt.Println(s)
}