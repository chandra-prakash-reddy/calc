package math

import "fmt"

func init() {
	fmt.Println("This Project is using version math-module:0.1.1")
}

func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func init() {
	fmt.Println("initialized package math (local).")
}
