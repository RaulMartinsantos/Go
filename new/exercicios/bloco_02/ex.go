package main

import "fmt"

func main() {
	fmt.Println("true && true =", true && true)
	fmt.Println("true && true =", true && false)
	fmt.Println("true && true =", false && true)
	fmt.Println("true && true =", false && false)

	fmt.Println("true || true =", true || true)
	fmt.Println("true || true =", false || true)
	fmt.Println("true || true =", true || false)
	fmt.Println("true || true =", false || false)
}
