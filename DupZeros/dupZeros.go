package main

import "fmt"

func duplicateZeros(arr []int) {
	if len(arr) < 2 {
		return
	}
	for i := len(arr) - 2; i >= 0; i-- {
		if arr[i] == 0 {
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			arr[i+1] = 0
		}
	}
}

func main() {
	arr := []int{1, 0, 2, 3, 0, 4, 5, 0}
	duplicateZeros(arr)
	fmt.Println(arr)
}
