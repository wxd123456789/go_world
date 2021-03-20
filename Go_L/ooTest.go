package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p Person) f() {
	fmt.Printf("in person")
}

func (p Person) String() string {
	return "name: ** age: **"
}

type Student struct {
	Person
	id int
}

func (s Student) f() {
	fmt.Println("in student")
}

type Human interface {
	f()
}

// interface {} 任意的类型，类似c中的void *
type Element interface{}
type myList []Element

func main() {
	person := Person{"parent", 18}
	student1 := Student{person, 1111}
	student1.f()
	fmt.Println(student1)
	var human Human
	human = person
	human.f()
	_, ok := human.(Person)
	if ok {
		fmt.Println("ok")
	}
	//
	list := make(myList, 3)
	list[0] = 12
	list[1] = "abc"
	list[2] = person
	for _, v := range list {
		switch v.(type) {
		case int:
			fmt.Println("int")
		case string:
			fmt.Println("string")
		case Person:
			fmt.Println("person")
		default:
		}

	}

}
