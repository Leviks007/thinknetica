package lesson9

import (
	"fmt"
	"io"
)

//Задание 1 начало
type Employee struct {
	name string
	age  int
}

func NewEmployee(name string, age int) *Employee {
	var e Employee
	e.name = name
	e.age = age
	return &e
}

func (e *Employee) GetAge() int {
	return e.age
}

type Customer struct {
	name string
	age  int
}

func NewCustomer(name string, age int) *Customer {
	var e Customer
	e.name = name
	e.age = age
	return &e
}

func (e *Customer) GetAge() int {
	return e.age
}

type AgeGet interface {
	GetAge() int
}

func MaxAge(user ...AgeGet) int {
	maxAge := 0
	for _, u := range user {
		userAge := u.GetAge()
		if maxAge < userAge {
			maxAge = userAge
		}
	}
	return maxAge
}

//Задание 1 конец

//Задание 2 начало
type Employee_2 struct {
	name string
	Age  int
}

func NewEmployee_2(name string, age int) *Employee_2 {
	var e Employee_2
	e.name = name
	e.Age = age
	return &e
}

type Customer_2 struct {
	name string
	Age  int
}

func NewCustomer_2(name string, age int) *Customer_2 {
	var e Customer_2
	e.name = name
	e.Age = age
	return &e
}

func MaxAge_2(users ...interface{}) interface{} {
	maxAge := 0
	var maxAgeElem interface{}

	for _, elem := range users {
		switch v := elem.(type) {
		case *Employee_2:
			if v.Age > maxAge {
				maxAge = v.Age
				maxAgeElem = elem
			}
		case *Customer_2:
			if v.Age > maxAge {
				maxAge = v.Age
				maxAgeElem = elem
			}
		}
	}
	return maxAgeElem
}

//Задание 2 конец

//Задание 3
func WriteStrings(w io.Writer, args ...interface{}) {
	for _, arg := range args {
		str := fmt.Sprintf("%v", arg)
		io.WriteString(w, str)
	}

}
