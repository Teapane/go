package main

import "fmt"

func main() {

	// empty
	m := make(map[string]int)
	fmt.Println("new map:", m)

	m["key 1"] = 12
	m["key 2"] = 13

	fmt.Println("new filled map:", m)

	v1 := m["key 1"]
	fmt.Println("value:", v1)

	fmt.Println("len:", len(m))

	_, prs := m["key 2"]
	fmt.Println("prs:", prs)

	// prebuilt
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}
