package main

import "fmt"

type Person struct{
	name string
	age int
 }

 func (p Person) speak(){
	fmt.Printf("My name is %v and I AM SPEAKING\n", p.name)
 }

func main(){
apple := Person {
	name: "John",
	age: 32,
}
// apple.name = "John"
// apple.age = 32
apple.speak()
}