package main

import "fmt"

type contactInfo struct {
	email string
	phone string
}
type person struct {
	firstName  string
	secondName string
	contact    contactInfo
}

func main() {
	alex := person{
		firstName:  "Alex",
		secondName: "Smith",
		contact: contactInfo{
			email: "alex@example.com",
			phone: "123-456-7890",
		},
	}
	fmt.Println(alex)
	alex.updateName("John")
	fmt.Println(alex)
	alex.updateLastName("Doe")
	fmt.Println(alex)
}

func (p person) updateName(newName string) {
	p.firstName = newName
}

func (p *person) updateLastName(newName string) {
	// (*p).secondName = newName
	p.secondName = newName
}

func testMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
