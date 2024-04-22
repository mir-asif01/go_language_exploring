package main

import "fmt"

type Human struct {
	name   string
	age    int
	salary int
}

func createHumanObj(name string, age int, salary int) *Human {
	h := Human{
		name:   name,
		age:    age,
		salary: salary,
	}
	return &h
}

type Personal_Info struct {
	f_name string
	l_name string
	age    int
}
type Contact struct {
	email string
	phone string
}
type Addreess struct {
	house_no int
	road     string
	area     string
}

type Person struct {
	Personal_Info Personal_Info
	Contact       Contact
	Addreess      Addreess
}

func main() {
	// var asif Human
	// asif.name = "Asif"
	// asif.age = 21
	// asif.salary = 10000

	// asif := createHumanObj("asif", 21, 10000)

	// fmt.Println("created by function", asif)

	var person1 Person
	person1.Personal_Info = Personal_Info{
		f_name: "Asif",
		l_name: "Ahsan",
		age:    21,
	}
	person1.Contact = Contact{
		email: "asif@gnail.com",
		phone: "013556",
	}
	person1.Addreess = Addreess{
		house_no: 154,
		road:     "dhaka road",
		area:     "mullanpur",
	}
	fmt.Println(person1)
}
