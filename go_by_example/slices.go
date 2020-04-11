package main

import "fmt"

func main() {
	s := make([]string, 10)
	fmt.Println("emp:", s)

	s[0] = "x"
	s[9] = "z"
	fmt.Println("Filled:", s)

	s = append(s, "tyty")
	fmt.Println("Appended:", s)
	s = append(s, "foo", "bar", "baz")
	fmt.Println("Appended:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
