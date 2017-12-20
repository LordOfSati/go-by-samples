package main

import "fmt"

func main() {
	/* sample if else block */
	num1, num2 := 10, 20
	if num1 > num2 {
		fmt.Println("num1 is greater")
	} else {
		fmt.Println("num2 is greater")
	}

	/* if with initialization - used in error handling scenarios */
	if name := getName(); name == "Hello" {
		fmt.Println("name is", name)
	}
}

func getName() string {
	return "Hello"
}
