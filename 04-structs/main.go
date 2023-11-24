package main

import "fmt"

// Value Types: int, float, string, bool, structs
// vs
// Reference Types: slices, maps, channels, pointers, functions

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// jigyo := new(person)
	// jigyo.firstName = "Jigyo"
	// jigyo.lastName = "Park"
	// fmt.Println(jigyo)

	// var youri person
	// fmt.Println(youri) // default value(zero value): empty string here
	// youri.firstName = "Youri"
	// youri.lastName = "Choi"
	// fmt.Println(youri)

	sunggon := person{
		firstName: "Sunggon",
		lastName:  "Park",
		contactInfo: contactInfo{
			email:   "f2hard3@gmail.com",
			zipCode: 3560017,
		},
	}

	// sunggon.updateName("gon", "Park")
	// sunggon.print()

	// &variable: operator turning a value to a pointer
	// sunggonPointer := &sunggon
	// sunggonPointer.updateName("gon", "Bak")

	sunggon.updateName("gon", "Bak")
	sunggon.print()
}

// *type: pointer to a person
func (pointerToPerson *person) updateName(newFirstName string, newLastName string) {
	// *variable: operator dereferencing the pointer
	(*pointerToPerson).firstName = newFirstName
	(*pointerToPerson).lastName = newLastName
}

// pass by value: copy the data inside the struct
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
