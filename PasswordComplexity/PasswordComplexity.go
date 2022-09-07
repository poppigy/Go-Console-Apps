package main

import (
	"fmt"
	"unicode"
)

func passCheck(pw string) bool {
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	hasUpper, hasLower, hasNumber, hasSymbol :=
		false, false, false, false
	for _, v := range pwR {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsDigit(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	a := string("myN@wfw42t2gq27kl8$&*@555")
	b := [5]rune{'a'}
	for _, val := range b {
		fmt.Println("val: |" + string(val) + "|")
	}
	fmt.Println(passCheck(a))
	fmt.Println()
}
