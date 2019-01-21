package main

import "fmt"

type apple int 
var y int


func main(){
var num apple
fmt.Println(num)
fmt.Printf("The type is %T\n", num)
num = 5
fmt.Println(num)
y = int(num)
fmt.Printf("The type is %T\n", y)
fmt.Println(y)


}

