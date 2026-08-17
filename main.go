package main

import "fmt"

var nums = []int{10, 20, 30, 40, 50, 60}

func main() {
	fmt.Println(nums[:2])
	fmt.Println(nums[2+1:])

	a := []int{10, 20}
	b := []int{30, 40, 50}
	a = append(a, b...)
}
