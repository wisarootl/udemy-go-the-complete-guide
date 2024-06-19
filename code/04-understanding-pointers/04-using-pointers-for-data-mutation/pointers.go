package main

import "fmt"

func main() {
	age := 32 // regular variable

	// var agePointer *int
	// agePointer = &age
	agePointer := &age // merge variable declaration into one line

	fmt.Println("Age:", *agePointer)

	editAgeToAdultYears(agePointer)
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age = *age - 18
}
