package main

import (
	"fmt"
)

const constString = "this is Const String OutSide the function."

func main() {
	fmt.Println("Hello from Go!")

	var aString string = "this is Golang!"
	fmt.Println(aString)
	fmt.Printf("The Variable type is %T.\n", aString)

	var aInteger int = 62
	fmt.Println(aInteger)
	fmt.Printf("The Variable type is %T.\n", aInteger)

	var defaultString string
	fmt.Println(defaultString)
	fmt.Printf("The Variable type is %T.\n", defaultString)

	var defaultInteger int
	fmt.Println(defaultInteger)
	fmt.Printf("The Variable type is %T.\n", defaultInteger)

	var anotherString = "This is another String."
	fmt.Println(anotherString)
	fmt.Printf("The Variable type is %T.\n", anotherString)

	var anotherInteger = 100
	fmt.Println(anotherInteger)
	fmt.Printf("The Variable type is %T.\n", anotherInteger)

	myString := "This is my String."
	fmt.Println(myString)
	fmt.Printf("The Variable type is %T.\n", myString)

	fmt.Println(constString)
	fmt.Printf("The Variable type is %T.\n", constString)
}
