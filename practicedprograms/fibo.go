package main

import "fmt"

var a, b, c, x int

func main() {

	fmt.Println("enter the number ")
	fmt.Scanf("%d", &x)
	a = 0
	b = 1
	fmt.Println(a)
	fmt.Println(b)

	c = a + b
	for i := 1; i < x-2; i++ {

		fmt.Println(c)
		a = b
		b = c
		c = a + b

	}
}
