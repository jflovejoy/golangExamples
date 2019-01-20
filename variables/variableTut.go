package main

import "fmt"

func main(){
	//Using the short declaration syntax, I can assign 3 variables with different type in one line
	a, b, c := 32, "James Bond", true
	fmt.Println(a, b, c)

	fmt.Printf("Hello, my name is %v and I am %v years old.\t This is all %v\n", b, a, c)
}