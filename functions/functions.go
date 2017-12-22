package main

import "fmt"

func main() {
	sayHello()
	sayHelloTo("World")
	fmt.Println(getFullName("John", "Doe"))
	number1, number2 := swapNumbers(10, 20)
	fmt.Println(number1, number2)
	fmt.Println(getName())
	sum := calculateSum(1, 2, 3, 4, 5, 6)
	fmt.Println("Sum is", sum)
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("Sum is", calculateSum(numbers...))

	incrementBy5 := makeIncrementor(5)
	fmt.Println(incrementBy5()) /* prints 5 */
	fmt.Println(incrementBy5()) /* prints 10 */

	incrementBy10 := makeIncrementor(10)
	fmt.Println(incrementBy10()) /* prints 10 */
	fmt.Println(incrementBy10()) /* prints 20 */

	/* prints only even numbers */
	filter(numbers, func(n int) bool {
		return n%2 == 0
	})

	fmt.Println("Factorial of 5 is", factorial(5))

	defer sayHello() /* this function call is deferred */

	sayHelloTo("Rose")
}

/* function with no parameter and return value */

func sayHello() {
	fmt.Println("Hello")
}

/* function with one parameter and no retun value */

func sayHelloTo(name string) {
	fmt.Println("Hello,", name)
}

/* function with multiple parameters and one return value */

func getFullName(firstName, lastName string) string {
	return fmt.Sprint(firstName, " ", lastName)
}

/* function with multiple parameters and return values */

func swapNumbers(number1, number2 int) (int, int) {
	return number2, number1
}

/* function with named return values */

func getName() (name string) {
	name = "Jack"
	return
}

/* function with variadic parameters */

func calculateSum(numbers ...int) int {
	var sum int
	for _, value := range numbers {
		sum += value
	}
	return sum
}

/* function with functions as return type */

func makeIncrementor(incrementBy int) func() int {
	var value int
	incrementor := func() int {
		value += incrementBy
		return value
	}
	return incrementor
}

/* passing functions as function parameters */

func filter(numbers []int, filterPredicate func(int) bool) {
	for _, value := range numbers {
		if filterPredicate(value) {
			fmt.Println(value)
		}
	}
}

/* recursive function */

func factorial(number int) int {
	if number == 0 || number == 1 {
		return 1
	}
	return number * factorial(number-1)
}
