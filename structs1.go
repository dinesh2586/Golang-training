package main

import "fmt"

type person struct {
	dob    string
	gender string
	race   string
}

func (per *person) updatedob() string {
	per.dob = "07/04/1999"
	return per.dob
}
func unittesting(x person) bool {

	y := x.updatedob()

	if y == "07/04/1999" {
		return true
	}
	return false

}

func main() {
	dinesh := person{dob: "07/04/2000", gender: "male", race: "asian"}
	fmt.Println(dinesh)
	fmt.Println(dinesh.updatedob())
	fmt.Println(dinesh)
	fmt.Println("unittest is: ", unittesting(dinesh))

}
