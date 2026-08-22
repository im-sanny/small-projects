package main

import "fmt"

var nums = []string{"one", "two", "skip three", "four", "five"}

func hof(a, b int, hf func(x, y int)) {
	hf(a, b)
}

func sum(a, b int) {
	c := a + b
	fmt.Println(c)
}

func main() {
	hof(2, 3, sum) // callback, func passed as parameter

	func(a, b int) { //IIFE can't be declared outside of a func
		c := a - b
		fmt.Println(c)
	}(5, 3)

	// nums = append(nums[:2], nums[2+1:]...)
	// fmt.Println(nums)

	// a := []int{1, 2}
	// b := []int{3, 4, 5}
	// a = append(a, b...)
	// fmt.Println(a)
}
