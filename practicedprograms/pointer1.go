package main

import "fmt"

func main() {

	states := [4]string{"Arizona", "florida", "Atlanta", "chicago"}
	fmt.Println("states are", states)
	updateThirdElement(&states[2])

	fmt.Println("states are", states)

}

func updateThirdElement(states *string) {
	*states = "texas"
	fmt.Println(*states)

}
