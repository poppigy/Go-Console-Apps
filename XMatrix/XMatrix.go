package main

func checkXMatrix(grid [][]int) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if i == j {
				if grid[i][j] == 0 {
					return false
				}
			} else if i+j == len(grid)-1 {
				if grid[i][j] == 0 {
					return false
				}
			} else {
				if grid[i][j] != 0 {
					return false
				}
			}

		}
	}
	return true
}

func main() {
	grid := [][]int{{2, 0, 0, 0, 1}, {0, 4, 0, 1, 5}, {0, 0, 5, 0, 0}, {0, 5, 0, 2, 0}, {4, 0, 0, 0, 2}}
	print(checkXMatrix(grid))
}
