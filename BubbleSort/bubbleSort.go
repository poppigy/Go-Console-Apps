package main

import "fmt"

func swap(l *int, r *int) {
	a := *l
	*l = *r
	*r = a
}

func bubbleSort(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				swap(&nums[j], &nums[j+1])
			}
		}
	}
	return nums
}

func main() {
	a := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	fmt.Println(bubbleSort(a))
}
