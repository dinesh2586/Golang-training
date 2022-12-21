package main

import (
	"fmt"

	app "app/calc"
)

func main() {
	fmt.Println(app.Sum(10, 20))
	fmt.Println(app.Sub(40, 50))
	fmt.Println(app.Multiply(7, 8))
	fmt.Println(app.Div(4, 8))
}
