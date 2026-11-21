package main

import "fmt"

func main() {
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("1 + 2.5 =", 1+2.5)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("100 * 25 =", 100*25)
	fmt.Println("100 * 12,5", 100*12.5)
	fmt.Println("5 % 3 =", 5%3)
	fmt.Printf("Tipo de 100 * 12,5 = %T \n ", 100*12.5)
}
