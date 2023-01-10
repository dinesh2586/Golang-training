package main

import (
	"fmt"
)

func main() {
	var a, b int
	fmt.Println("ENTER VALUE A ")
	fmt.Scanln(&a)
	fmt.Println("ENTER VALUE b ")
	fmt.Scanln(&b)

	fmt.Printf("sum of %v and %v \n", a, b, app.Sum(&a, &b))
	fmt.Printf("sub of %v and %v \n", a, b, app.Sub(&a, &b))
	fmt.Printf("multiply of %v and %v\n", a, b, app.Multiply(&a, &b))
	fmt.Println("quotient of %v and %v\n", a, b, app.Div(&a, &b))
	fmt.Println("remainder of %v and %v\n", a, b, app.Div(&a, &b))

}
