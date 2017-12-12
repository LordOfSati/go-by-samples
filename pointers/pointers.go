package main

import "fmt"

func main() {
	name := "Hello" /* type string */
	addressOfName := &name
	valueInAddress := *addressOfName /* type *string */

	fmt.Println(name)
	fmt.Println(addressOfName)
	fmt.Println(valueInAddress)

	num1 := 10
	fmt.Println(num1) /* prints 10 */
	passByValueExample(num1)
	fmt.Println(num1) /* prints 10 */

	num2 := 10
	fmt.Println(num2) /* prints 10 */
	passByReferenceExample(&num2)
	fmt.Println(num2) /* prints 110 */
}

func passByValueExample(num int) {
	num = num + 100
}

func passByReferenceExample(num *int) {
	*num = *num + 100
}
