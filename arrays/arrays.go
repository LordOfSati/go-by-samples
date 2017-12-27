package main

import "fmt"

func main() {
	/* arrays are of fixed length */
	/* arrays are value types */
	/* array items are initialized with default values */
	var intArray [5]int
	fmt.Println(intArray)
	intArray[0] = 10
	intArray[1] = 20
	intArray[2] = 30
	intArray[3] = 40
	intArray[4] = 50
	for _, value := range intArray {
		fmt.Println(value)
	}
	for i := 0; i < len(intArray); i++ {
		fmt.Println(intArray[i])
	}
	fmt.Println(intArray[0:3])
	/* two dimensional array */
	var matrix [2][2]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			matrix[i][j] = i + j
		}
	}
	fmt.Println(matrix)
}
