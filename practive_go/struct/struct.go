package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var sahil person
	sahil.firstName = "Sahil"
	sahil.lastName = "Verma"
	sahil.age = 23

	fmt.Println(sahil)
 //                              --------- another way of writing---------
	vijay := person{
		firstName: "vijay",
		lastName: "verma",
		age: 25,
	}

	fmt.Println(vijay)

}