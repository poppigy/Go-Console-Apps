package main

import "fmt"

func developerSalary(x, y int) int {
	return x * y
}

func managerSalary(x, y int) int {
	return x + y
}

func salary(x, y int, f func(int, int) int) int {
	pay := f(x, y)
	return pay
}

func main() {
	devSalary := salary(50, 2080, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)
	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Dev salary: %d\n", devSalary)
}
