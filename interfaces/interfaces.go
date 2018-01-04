package main

import (
	"fmt"
	"math"
	"sort"
)

/* Example 1 */

type square struct {
	side float64
}

/* implementation of area() from shape interface */
func (s square) area() float64 {
	return s.side * s.side
}

type circle struct {
	radius float64
}

/* implementation of area() from shape interface */
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func printArea(s shape) {
	fmt.Println(s.area())
}

/* Example 2 */

type person struct {
	name string
	age  int
}

type people []person

/* below 3 methods implicitly implement methods from sort.Interface */

func (p people) Len() int {
	return len(p)
}

func (p people) Less(i, j int) bool {
	return p[i].age > p[j].age
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

/* Example 3 - embedded interfaces */

type reader interface {
	read()
}

type writer interface {
	write()
}

/* Example 4 - empty interfaces */

type any interface{}

func printValues(input any) {
	fmt.Println(input)
}

/* reader and writer interfaces are embedded in readWriter interface */
type readWriter interface {
	reader
	writer
}

func main() {
	/* Example 1 */
	sq := square{10.0}
	printArea(sq)
	ci := circle{10.0}
	printArea(ci)
	/* Example 2 */
	students := people{person{"Joe", 31}, person{"Ajay", 21}, person{"Jade", 20}}
	sort.Sort(students)
	fmt.Println(students)
	/* Example 4 */
	printValues(sq)
	printValues(ci)
	printValues(students)
}
