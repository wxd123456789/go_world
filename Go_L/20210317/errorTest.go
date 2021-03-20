package main

import "fmt"

func main() {
	test()
}

func catchError() {
	if err := recover(); err != nil {
		fmt.Println(err.(string))
	}
}

func test() {
	defer catchError()
	panic("panic error!")
	fmt.Println("end")
}
