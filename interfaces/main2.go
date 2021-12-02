package main

import (
	"fmt"
)

// If we dont know the type of the function and type of argument of the function during compile time then we can use empty interfaces,
// here someArgument can be int, float, etc.
func SomeFunction(someArgument interface{}) {
	fmt.Println(someArgument)
}

func main() {
	fmt.Println("empty interfaces")
	SomeFunction("hello world")
	SomeFunction(97)
	// Instantiated empty struct
	SomeFunction(struct{}{})
	// Instantiate struct with name
	// struct{Name string}{"Tom"}

	//  map string to interface means key of the map should be string and value can be anything
	myMap := make(map[string]interface{})
	myMap["Name"] = "Rizwana"
	myMap["number"] = 10
	fmt.Println(myMap)
}
