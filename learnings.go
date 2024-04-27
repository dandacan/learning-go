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
	
	//real and imaginary components to the imaginary numbers
	fmt.Printf("%v,  %T\n", real(aye), real(aye))
	fmt.Printf("%v,  %T\n", imag(aye), imag(aye))

	//creating an imaginary number
	var n complex128 = complex(5,12)
	fmt.Printf("%v,  %T\n", n, n)
}
