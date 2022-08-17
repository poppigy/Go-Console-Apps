package main

import "fmt"

type message func()

func main() {
	increment := incrementor()
	fmt.Println(increment())
	fmt.Println(increment())
}
func incrementor() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
