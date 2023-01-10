package main

import "fmt"

func main() {
	nums := []int{2, 1, 3, 3}
	val := 1
	removeelement(nums, val)
	fmt.Println(nums)

}

func removeelement(nums []int, val int) int {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j++
		}

	}
	return j
}
