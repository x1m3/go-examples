package sequence_test

import (
	"fmt"
)


// intGenerator is a function closure that returns an int64 returning function. The closure "has internal memory", so
// it remembers the last value and generates a new one
func intGenerator() func() int64 {
	var current int64
	current = 0
	return func() int64 {
		current++
		return current
	}
}


func ExampleIntGenerator() {

	// We create an instance of the generator, called next
	next := intGenerator()

	for i:=0;i<10;i++ {
		// Everytime we call next() we got a new value
		fmt.Println(next())
	}
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10
}

func dayOfWeekGenerator() func() string {
	daysOfWeek := [7]string {"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	current := 0
	return func() string {
		day := daysOfWeek[current%7]
		current++
		return day
	}
}

func ExampleDaysOfWeekGenerator() {

	nextDay := dayOfWeekGenerator()
	for i:=0;i<10;i++ {
		fmt.Println(nextDay())
	}
	// Output:
	// Monday
	// Tuesday
	// Wednesday
	// Thursday
	// Friday
	// Saturday
	// Sunday
	// Monday
	// Tuesday
	// Wednesday
}


