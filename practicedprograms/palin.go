package main

import "fmt"

func main() {
	var a, b, c int
	var reverse int = 0

	fmt.Println("enter the number")
	fmt.Scan(&a)

	c = a
	for {
		b = a % 10
		reverse = reverse*10 + b
		a /= 10

		if a == 0 {
			break

		}
	}

	if c == reverse {
		fmt.Println(c, "is a Palindrome")
	} else {
		fmt.Println(c, "is not a Palindrome")
	}

}
