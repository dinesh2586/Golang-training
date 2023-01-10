package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 15}
	var target int
	fmt.Println("enter the target")
	fmt.Scanln(&target)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				fmt.Println([]int{i, j})
			}

		}
	}
	fmt.Println([]int{})
}
