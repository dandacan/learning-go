package main

import "fmt"

//cant use :=, not visible outside of the package
var thing string = "some letters have been added"

//can be used outside the package because capital letter
var I int = 420

var someNumbers float32 = 420.69

//complex numbers exist
var aye complex64 = 2i

//rune type = UTF-32 char

//constants are unchangeable values that behave like vars but can't be changed
const theMeaningOfLife = 42

const(
	//ignore the first value here
	_ = iota
	number = 1 + iota //1 + 1 = 2
	anotherNumber // auto assigned 3

)
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

	fmt.Printf("%v, %v", number, anotherNumber)
}
