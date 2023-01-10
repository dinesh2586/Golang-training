package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter the number ")
	fmt.Scanln(&num)
	if num < 1 {
		fmt.Println(" not a prime")
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			fmt.Println(num, "is  not prime number")
			return
		} else {
			fmt.Println(num, "is a prime number")
			return
		}

	}
}
