package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func (p person) getFullName() string {
	return p.firstName + " " + p.lastName
}

type employee struct {
	person     /* embedded type */
	department string
}

func (e employee) getFullName() string {
	return e.person.getFullName()
}

func (e *employee) getDepartment() string {
	return "Employee belongs to " + e.department + " department"
}

func main() {
	person1 := person{firstName: "John", lastName: "Doe", age: 20}
	fmt.Println(person1) /* prints {John Doe 20} */

	employee1 := employee{person: person1, department: "IT"}
	fmt.Println(employee1.person.firstName) /* prints John */
	fmt.Println(employee1.department)       /* prints IT */

	employee2 := employee{person: person{firstName: "Jane", lastName: "Smith", age: 30}, department: "FINANCE"}
	fmt.Println(employee2) /* prints {{Jane Smith 30} FINANCE} */

	fmt.Println(employee2.getFullName())   /* prints Jane Smith */
	fmt.Println(employee2.getDepartment()) /* prints Employee belongs to FINANCE department */

	newEmployee := &employee2
	fmt.Println(newEmployee)                  /* prints &{{Jane Smith 30} FINANCE} */
	fmt.Println(newEmployee.person.firstName) /* prints Jane */
}
