package main

import "fmt"

type Person struct {
	StreetAddress, PhoneNumber, City, PostalCode string
	Salary                                       float64
	Position                                     string
}
type PersonBuilder struct {
	person *Person //composition
}
type PersonAddressBuilder struct {
	PersonBuilder
}
type PersonJobInfoBuilder struct {
	PersonBuilder
}

func (pb *PersonBuilder) Works() *PersonJobInfoBuilder {
	return &PersonJobInfoBuilder{*pb}
}
func (pjb *PersonJobInfoBuilder) AsA(position string) *PersonJobInfoBuilder {
	pjb.person.Position = position
	return pjb
}
func (pjb *PersonJobInfoBuilder) Earning(salary float64) *PersonJobInfoBuilder {
	pjb.person.Salary = salary
	return pjb
}

func (pb *PersonBuilder) Lives() *PersonAddressBuilder {
	return &PersonAddressBuilder{*pb}
}
func (pjb *PersonAddressBuilder) At(street string) *PersonAddressBuilder {
	pjb.person.StreetAddress = street
	return pjb
}
func (pjb *PersonAddressBuilder) In(city string) *PersonAddressBuilder {
	pjb.person.City = city
	return pjb
}
func (pjb *PersonAddressBuilder) WithPostCode(postalCode string) *PersonAddressBuilder {
	pjb.person.PostalCode = postalCode
	return pjb
}

func NewPersonBuilder() *PersonBuilder {
	return &PersonBuilder{&Person{}}
}
func (pb *PersonBuilder) Build() *Person {
	return pb.person
}

func main() {
	pb := NewPersonBuilder()

	pb.Works().
		AsA("Programmer").
		Earning(12.500).
		Lives().
		In("Tehran").
		At("X").
		WithPostCode("154SW")

	person := pb.Build()
	fmt.Println(person)
}
