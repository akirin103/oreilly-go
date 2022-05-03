package main

import "fmt"

type Mystruct struct {
	Name string
	Address string
	Age int
}

func add(x int, y int, z int) int {
	return x + y + z
}

func main() {
	b := Mystruct{
		Name:    "",
		Address: "",
		Age:     23,
	}
	fmt.Println("学習用", b.Name)
	c := add(1, 2, 3)
	fmt.Println(c)
}
