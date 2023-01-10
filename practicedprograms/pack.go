package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, a := os.Create("text.txt")
	if a != nil {
		fmt.Println("error", a)
	}
	defer f.Close()
	fmt.Println(f.Name())
	fmt.Println("enter the string")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	b := scanner.Text()
	_, a = f.WriteString(b)
	if a != nil {
		fmt.Println("error", a)
	}

	data, _ := ioutil.ReadFile("text.txt")
	fmt.Printf("%s", data)
}
