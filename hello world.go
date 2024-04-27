package main

import "fmt"

//cant use :=
var thing string = "some letters have been added"

func main() {
	//standard hello world using vars that are declared differently
	var hi string = "hello"
	world := " world"

	//printing the strings
	fmt.Println(hi + world)
	fmt.Println(thing)
}
