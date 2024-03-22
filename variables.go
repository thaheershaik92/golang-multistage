package main

import "fmt"

func main() {
	// var name string
	// var name = "Wahyu Setiawan"

	name := "Wahyu Setiawan"
	fmt.Println(name)

	name = "Wahyu Usman"
	fmt.Println(name)

	var (
		firstName  = "Wahyu"
		middleName = "Setiawan"
		lastName   = "Usman"
	)

	fmt.Println(firstName, middleName, lastName)
}
