package main

import "fmt"

func main() {
	/* maps are reference types */
	map1 := make(map[int]string)
	map2 := map[string]int{"John": 1, "James": 2}
	map1[1] = "John"
	map1[2] = "James"
	fmt.Println(map1) /* prints map[1:John 2:James] */
	fmt.Println(map2) /* prints map[John:1 James:2] */
	for key, value := range map1 {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
	delete(map1, 1)
	delete(map2, "John")
	fmt.Println(map1) /* prints map[2:James] */
	fmt.Println(map2) /* prints map[James:2] */
	fmt.Println(len(map1), len(map2))
	if value, exists := map1[2]; exists {
		fmt.Println("Value is", value) /* prints James */
	}
}
