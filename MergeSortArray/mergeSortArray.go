package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n < 1 {
		return
	}
	x, y := m-1, n-1
	i := m + n - 1
	for x >= 0 && y >= 0 {
		if nums2[y] > nums1[x] {
			nums1[i] = nums2[y]
			y--
		} else {
			nums1[i] = nums1[x]
			x--
		}
		i--
	}
	for y >= 0 {
		nums1[i] = nums2[y]
		y--
		i--
	}
	for x >= 0 {
		nums1[i] = nums1[x]
		x--
		i--
	}
	nums1 = nums1[i+1:]
}

func main() {
	//arr1 := []int{1, 2, 3, 0, 0, 0}
	//arr2 := []int{2, 5, 6}

	//arr1 := []int{1}
	//arr2 := []int{}

	//arr1 := []int{2, 0}
	//arr2 := []int{1}

	arr1 := []int{4, 5, 6, 0, 0, 0}
	arr2 := []int{1, 2, 3}
	merge(arr1, 3, arr2, 3)
	fmt.Println(arr1)
}
