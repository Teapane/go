package main

import "fmt"

// argument type, argument type, return value
func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func vals() (int, int) {
	return 3, 7
}

func plainFuncs() {

}

// variadic function
func sum(nums ...int) {
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)
}

// anonymous function/closures
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//recursion recursion recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	// functions
	add := add(5, 10)
	fmt.Println("addition 5 + 10", add)

	subtract := subtract(5, 10)
	fmt.Println("Subtraction 5 - 10=", subtract)

	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println("C?", c)

	//variadic
	nums := []int{1, 2, 3, 4}

	sum(1, 2)
	sum(1, 2, 3)
	sum(nums...)

	//closures
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(7))
}
