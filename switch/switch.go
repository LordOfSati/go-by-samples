package main

import "fmt"

func main() {
	/* switch statement in Go doesn't fallthrough by default */
	num := 11
	switch num {
	case 1:
		fmt.Println("Num is 1")
	case 10, 11:
		fmt.Println("Num is 10 or 11")
	case 100:
		fmt.Println("Num is 100")
	}
	/* switch on string + fallthrough */
	str := "hello"
	switch str {
	case "hello":
		fmt.Println("Hello")
		fallthrough
	case "world":
		fmt.Println("World")
	default:
		fmt.Println("Default case!!")
	}

	/* switch without input */
	number := 100
	switch {
	case number < 10:
		fmt.Println("Number less than 100")
	case number >= 10 && number < 100:
		fmt.Println("Number is >= 10 and < 100")
	default:
		fmt.Println("Number is >= 100")
	}

	/* type switch */
	var flag interface{}
	flag = true
	switch flag.(type) {
	case bool:
		fmt.Println("flag is boolean")
	case int:
		fmt.Println("flag is int")
	}
}
