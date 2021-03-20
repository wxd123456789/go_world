package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

func main() {
	testSlice()
	testMap()
	testS()
	jsonTest()
	mapSorted()
	//
	ifTest()
	referTest()
}

func testSlice() {
	aR := [10]int{1, 2, 3, 4, 4, 5}
	aSlice := make([]int, 2)
	aSlice[0] = 1
	aSlice = append(aSlice, 1, 2, 3)
	bSlice := []int{1, 23, 3}
	cSlice := aR[0:2:4]
	fmt.Println(aSlice, bSlice, cSlice)
}

func testMap() {
	myMap := make(map[string]int, 3)
	myMap["a"] = 1
	myMap["b"] = 2

	myMap2 := map[string][]int{
		"a": {1, 2},
		"b": {3, 4},
	}
	fmt.Println(myMap, myMap2)

	k := "wxd"
	v, ok := myMap2[k]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Printf("%s not exists\n", k)
	}
	delete(myMap2, "a")
}

type Student2 struct {
	name string
	age  int
}

type Student struct {
	ID       int `json:"id"`
	Gender   string
	Name     string
	privatek string
}

type Class struct {
	Title    string
	Students []*Student
}

func jsonTest() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)

}

func testS() {
	stus := []Student2{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}
	m := make(map[string]Student2)
	for _, p := range stus {
		//fmt.Println(p.name, p)
		m[p.name] = p
	}
	for _, v := range m {
		fmt.Println(v)
	}
}

func mapSorted() {
	map1 := make(map[int]string, 5)
	map1[1] = "www.topgoer.com"
	map1[2] = "rpc.topgoer.com"
	map1[5] = "ceshi"
	map1[3] = "xiaohong"
	map1[4] = "xiaohuang"
	slist := []int{}
	for k, _ := range map1 {
		slist = append(slist, k)
	}
	sort.Ints(slist)
	for _, v := range slist {
		fmt.Println(map1[v])
	}

}

func ifTest(){
	var k = 0
	switch k {
	case 0:
		println("fallthrough")
		fallthrough
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
}

func referTest(){
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}