package main

import (
	"fmt"
	"time"
)

func main() {
	var i string = "hubba"
	fmt.Print("Write ", i, " as ")
	switch i {
	case "hello":
		fmt.Println("hello")
	case "hi":
		fmt.Println("two")
	case "hubba":
		fmt.Println("huyya")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")

	println()
}
