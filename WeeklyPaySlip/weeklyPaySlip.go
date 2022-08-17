package main

import (
	"errors"
	"fmt"
)

var (
	ErrHourlyRate  = errors.New("invalid hourly rate")
	ErrHoursWorked = errors.New("invalid hours worked per week")
)

func payDay(hoursWorked, hourlyRate int) (int, error) {
	if hoursWorked < 0 || hoursWorked > 80 {
		return 0, ErrHoursWorked
	}
	if hourlyRate < 10 || hourlyRate > 75 {
		return 0, ErrHourlyRate
	}
	if hoursWorked > 40 {
		overtime := (hoursWorked - 40) * hourlyRate * 2
		regularPay := 40 * hourlyRate
		return regularPay + overtime, nil
	} else {
		return hoursWorked * hourlyRate, nil
	}
}

func main() {
	pay, err := payDay(81, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pay)
	}
	pay, err = payDay(80, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pay)
	}
	pay, err = payDay(80, 50)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(pay)
	}
}
