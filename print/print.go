package main

import "fmt"

func main() {
	/* prints Decimal - 99 */
	fmt.Printf("Decimal - %d \n", 99)

	/* prints Binary - 99 : 1100011 */
	fmt.Printf("Binary - %d : %b \n", 99, 99)

	/* prints  Oct - 99 : 143 : 0143 */
	fmt.Printf("Oct - %d : %o : %#o \n", 99, 99, 99)

	/* prints  Hex - 99 : 63 : 0x63 */
	fmt.Printf("Hex - %d : %x : %#x \n", 99, 99, 99)

	/* prints  Hex - 29 : 1D : 0X1D */
	fmt.Printf("Hex - %d : %X : %#X \n", 29, 29, 29)

	/* prints  Float - 99.990000 */
	fmt.Printf("Float - %f \n", 99.99)

	/* prints  Float - 9.999000e+01 */
	fmt.Printf("Float - %e \n", 99.99)

	/* prints Formatted float -    99.990 */
	fmt.Printf("Formatted float - %9.3f \n", 99.99)

	/* prints type of 99.99 is float64 */
	fmt.Printf("type of 99.99 is %T \n", 99.99)

	/* prints type of 99 is int */
	fmt.Printf("type of 99 is %T \n", 99)
}
