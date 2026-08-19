package main

import "fmt"

var nums = []string{"one", "two", "skip three", "four", "five"}

func main() {
	nums = append(nums[:2], nums[2+1:]...)
	fmt.Println(nums)

	a := []int{1, 2}
	b := []int{3, 4, 5}
	a = append(a, b...)
	fmt.Println(a)
}
