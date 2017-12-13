package main

import "fmt"

func main() {
	/* for loop 1 */
	for i := 5000; i <= 5100; i++ {
		fmt.Println(string(i))
	}
	/* for loop 2 */
	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}
	/* for loop 3 */
	k := 1
	for {
		fmt.Println(k)
		k++
		if k > 5 {
			break
		}
	}
	/* printing even numbers using continue */
	for num := 1; num <= 25; num++ {
		if num%2 != 0 {
			continue
		}
		fmt.Println(num)
	}
}
