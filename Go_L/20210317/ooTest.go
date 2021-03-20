package main

import "fmt"

type People interface {
	Speak(string) string
}

type Student22 struct{}

func (stu *Student22) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}

func main() {
	var peo People = &Student22{}
	think := "bitch"
	fmt.Println(peo.Speak(think))
	//
	testTypeAssert()
}

func testTypeAssert(){
	var x interface{}
	s := "123"
	x = s
	v, ok := x.(string)
	if ok{
		fmt.Printf("T: %T, V: %s\n", v, v)
	}
	justifyType(123)
}

func justifyType(x interface{}) {
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}
