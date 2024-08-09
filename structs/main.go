package main

import "fmt"

//Here we are creating a new type of contactInfo which is a struct
type contactInfo struct {
	email string
	zipCode int
}
//Here we are creating a new type of person which is a struct
type person struct {
	firstName string
	lastName  string
	Contact contactInfo
}

func main() {

	//Here we are creating a new person struct
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		Contact: contactInfo {
			email: "email",
			zipCode: 94000,
		},
	}
	//Here we assign the address of the variable jim to the variable jimPointer
	// jimPointer := &jim
	jim.updateName("Jimmy")
	jim.print()
}

//Here we are using a receiver function for the person struct
func (p person) print() {
	fmt.Printf("%+v", p)
}

//Here we are using a pointer to the person struct
func (pointerToPerson *person) updateName(newFirstName string) {
	//Here we are using the * operator to get the value from the address
	(*pointerToPerson).firstName = newFirstName
}