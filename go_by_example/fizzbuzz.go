package main

import "fmt"

func main() {

	const (
		FIZZ     = 3
		BUZZ     = 5
		FIZZBUZZ = 15
	)

	// fizzzbuzzzzinnn \0/
	for i := 1; i <= 1000; i++ {
		switch {
		case i%FIZZBUZZ == 0:
			fmt.Printf("%d fizzbuzz\n", i)
		case i%FIZZ == 0:
			fmt.Printf("%d fizz\n", i)
		case i%BUZZ == 0:
			fmt.Printf("%d buzz\n", i)
		default:
			fmt.Printf("%d\n", i)
		}
	}
}
