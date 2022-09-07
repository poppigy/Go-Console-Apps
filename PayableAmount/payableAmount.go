package main

import "fmt"

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func (d *Developer) LogHours(w Weekday, a int) {
	d.WorkWeek[w] = a
}

func (d *Developer) HoursWorked() int {
	a := 0
	for i := range d.WorkWeek {
		a += d.WorkWeek[i]
	}
	return a
}

func nonLoggedHours() func(int) int {
	total := 0
	x := func(i int) int {
		total += i
		return i
	}
	return x
}

func (d *Developer) PayDay() (int, bool) {
	paidHours := d.HoursWorked()
	if paidHours <= 40 {
		return paidHours * d.HourlyRate, false
	} else {
		overtime := (paidHours - 40) * d.HourlyRate * 2
		return 40*d.HourlyRate + overtime, true
	}
}

func (d *Developer) PayDetails() {
	fmt.Printf("Sunday hours: %d\n", d.WorkWeek[0])
	fmt.Printf("Monday hours: %d\n", d.WorkWeek[1])
	fmt.Printf("Tuesday hours: %d\n", d.WorkWeek[2])
	fmt.Printf("Wednesday hours: %d\n", d.WorkWeek[3])
	fmt.Printf("Thursday hours: %d\n", d.WorkWeek[4])
	fmt.Printf("Friday hours: %d\n", d.WorkWeek[5])
	fmt.Printf("Saturday hours: %d\n\n", d.WorkWeek[6])

	fmt.Printf("Hours worked this week: %d\n", d.HoursWorked())
	pay, overtime := d.PayDay()
	fmt.Printf("Pay for the week: %d\n", pay)
	fmt.Printf("Is this overtime pay: %v\n", overtime)
}

func main() {
	d := Developer{
		Individual: Employee{
			Id: 0, FirstName: "Name", LastName: "LastName",
		},
		HourlyRate: 10,
	}
	defer d.PayDetails()
	x := nonLoggedHours()
	fmt.Println("Tracking hours worked thus far today: ", x(2))
	fmt.Println("Tracking hours worked thus far today: ", x(3))
	fmt.Println("Tracking hours worked thus far today: ", x(5))
	fmt.Println()
	d.LogHours(Monday, 8)
	d.LogHours(Tuesday, 10)
	d.LogHours(Wednesday, 10)
	d.LogHours(Thursday, 10)
	d.LogHours(Friday, 6)
	d.LogHours(Saturday, 8)
	//d.PayDetails()
}
