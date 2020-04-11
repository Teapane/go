package main

import "fmt"

func main() {
	// string concatenation
	fmt.Println("go" + "lang")

	// addition
	fmt.Println("1 + 1 =", 1+1)

	// division
	fmt.Println("7/3 =", 7.0/3.0)

	// booleans
	fmt.Println(true && false)

	fmt.Println(true || false)

	fmt.Println(!true)

	var a = "Initial"

	fmt.Println(a)
	// "Initial"

	var b, c = 1, 2
	fmt.Println(b, c)
	// 1, 2

	var d = true
	fmt.Println(d)
	// true

	var e int
	fmt.Println(e)
	// 0

	// shorthand for "var"
	f := "foo"

	fmt.Println(f)
	// foo
}
