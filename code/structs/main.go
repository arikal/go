package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode string
}

func (p person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p *person) updateName(firstName string) {
	(*p).firstName = firstName
}

func main() {
	mark := person{
		firstName: "Mark",
		lastName:  "Shercliff",
		contactInfo: contactInfo{
			email:   "mark@example.com",
			zipCode: "AB123 4CD",
		},
	}

	mark.updateName("Mike")

	fmt.Printf("%+v", mark)
	fmt.Println(mark.fullName())
}
