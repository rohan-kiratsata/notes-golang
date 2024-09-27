package main

import "fmt"

func main() {

	rohan := User{"Rohan", "rohan@gmail.com", true, 20}
	fmt.Println("rohan:", rohan)
	fmt.Printf("rohan details are: %+v\n", rohan)
	fmt.Printf("name is %v and email is %v", rohan.Name, rohan.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
