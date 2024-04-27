package main

import "fmt"

//cant use :=, not visible outside of the package
var thing string = "some letters have been added"

//can be used outside the package because capital letter
var I int = 420

var someNumbers float32 = 420.69

//complex numbers exist
var aye complex64 = 2i

func main() {
	//standard hello world using vars that are declared differently
	var hi string = "hello"
	world := " world"

	//printing the strings
	fmt.Println(hi + world)
	fmt.Println(thing)

	//type conversion loses data sometimes
	fmt.Println(someNumbers)
	fmt.Println(int32(someNumbers))
	
	fmt.Printf("%v,  %T\n", aye, aye)

}
