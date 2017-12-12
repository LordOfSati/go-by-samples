package main

import "fmt"

/* variable declaration */

var number1 int
var string1 string
var float1 float32
var bool1 bool

/* declaring multiple variables in single statement */

var b1, b2, b3 bool

/* variable declaration with type inference */

var number4 = 100
var string4 = "Hello from Go!!"

/* constant declaration */

const number3 = 100
const float3 = 999.999

/* const block declaration */

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {

	/* assigning values to variables */

	number1 = 10
	string1 = "Hello"
	float1 = 10.99
	bool1 = true

	/* short hand declaration and instantiation */

	number2 := 20
	string2 := "Hello World"
	bool2, float2 := true, 99.99

	fmt.Println(number2)
	fmt.Println(string2)
	fmt.Println(bool2)
	fmt.Println(float2)

	fmt.Println("kb: ", kb)
	fmt.Println("mb: ", mb)
	fmt.Println("gb: ", gb)
}
