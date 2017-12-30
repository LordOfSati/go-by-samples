package main

import "fmt"

func main() {
	/* slices are reference types */
	var slice1 []int
	slice2 := []int{100, 200, 300, 400, 500}
	var slice3 = make([]int, 5, 10)
	if slice1 == nil {
		fmt.Println("slice1 is nil") /* prints */
	}
	fmt.Println(slice1)      /* prints [] */
	fmt.Println(slice2)      /* prints [100 200 300 400 500] */
	fmt.Println(slice3)      /* prints [0 0 0 0 0] */
	fmt.Println(len(slice3)) /* prints 5 */
	fmt.Println(cap(slice3)) /* prints 10 */
	slice1 = append(slice1, 10)
	slice1 = append(slice1, 20)
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
	for _, value := range slice1 {
		fmt.Println(value)
	}
	fmt.Println(slice1[:])  /* prints [10 20 100 200 300 400 500] */
	fmt.Println(slice1[2:]) /* prints [100 200 300 400 500] */
	fmt.Println(slice1[:4]) /* prints [10 20 100 200] */

	/* 2 dimensional slice */

	slice4 := [][]int{}
	fmt.Println(slice4) /* prints [] */
	slice4 = append(slice4, slice1)
	slice4 = append(slice4, slice2)
	fmt.Println(slice4) /* prints [[10 20 100 200 300 400 500] [100 200 300 400 500]] */

	copy(slice3, slice1)
	fmt.Println(slice3) /* prints [10 20 100 200 300] */
}
