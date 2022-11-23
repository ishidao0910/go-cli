package main

import "fmt"

func test() {
	var testString string = "Go programming"	// OK
	var TestString2 string = "Go programming"	// Wrong
	var testInt int = 10						// OK
	var testInt2 int = 10						// OK
	var test_int3 int = 10						// Wrong

	fmt.Println(testString)
	fmt.Println(TestString2)
	fmt.Println(testInt)
	fmt.Println(testInt2)
	fmt.Println(test_int3)
}
