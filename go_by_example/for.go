package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("J =", j)
	}

	for {
		fmt.Println("Loopin'")
		break
	}

	for n := 0; n <= 500; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("n =", n)
	}
}

//for is Go’s only looping construct. Here are some basic types of for loops.
