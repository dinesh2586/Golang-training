package main

import "fmt"

func main() {
	a := [5]int{10, 20, 30, 40, 50}

	fmt.Printf("%T\n", a)

	slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("%T\n", slice)

	fmt.Println("slice is", slice[0:5])
	fmt.Println("slice is", slice[5:])
	fmt.Println("slice is", slice[2:7])
	fmt.Println("slice is", slice[1:6])

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println("x is \n", x)
	x = append(x, 53, 54, 55)
	fmt.Println("x is \n", x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)

	fmt.Println("x ,y is ", x)

}
