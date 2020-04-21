package main

import(
	"fmt"
)

func main(){
	fmt.Println("Hello")
	fmt.Println(ReturnFive("world"))
}
////
func ReturnFive(s string) int{
	return 5
}
