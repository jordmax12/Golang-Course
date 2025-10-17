package main

import "fmt"

type Person struct {
	Age int
	AdultAge int
}

func main() {
	person := &Person{
		Age: 32,
		AdultAge: 0,
	}

	fmt.Println(person)
	fmt.Println(person.Age)

	getAdultYears(person)
	fmt.Println(person.AdultAge)

	intAge := getAdultYearsInt(&person.Age)
	fmt.Println(intAge)

}

func getAdultYears(person *Person) {
	person.AdultAge = person.Age - 18
}

func getAdultYearsInt(age *int) int {
	return *age - 18
}
