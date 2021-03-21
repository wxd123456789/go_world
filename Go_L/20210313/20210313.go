package main

import (
	"errors"
	"fmt"
)

// enum
const (
	Unknown = 0
	F       = 1
	M       = 2
)

func main() {
	fmt.Println(123)
	fmt.Println(Unknown)
	// iota 行索引
	const (
		a = iota //0
		b        //1
		c        //2
		d = "ha" //独立值，iota += 1
		e        //"ha"   iota += 1
		f = 100  //iota +=1
		g        //100  iota +=1
		h = iota //7,恢复计数
		i        //8
	)
	//
	switch {
	case false:
		fmt.Println("1、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("2、case 条件语句为 true")
		fallthrough
	case false:
		fmt.Println("3、case 条件语句为 false")
		fallthrough
	case true:
		fmt.Println("4、case 条件语句为 true")
	case false:
		fmt.Println("5、case 条件语句为 false")
		fallthrough
	default:
		fmt.Println("6、默认 case")
	}
	/*2、case 条件语句为 true
	3、case 条件语句为 false
	4、case 条件语句为 true*/

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}

	_slice := []string{"1",}
	for i, v := range _slice {
		fmt.Println(i, v)
	}

	//
	aArray := [3]int{1, 2, 3}
	bArray := [...]int{1, 2, 4}
	fmt.Println(aArray, bArray)
	var values [][]int
	row1 := []int{1, 2, 3}
	row2 := []int{4, 5, 6}
	values = append(values, row1)
	values = append(values, row2)
	_mArray := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(_mArray)
	books := Books{title: "123", author: "wxd", subject: "12", bookId: 11}
	fmt.Println(books)

	//slice
	sliceA := []int{1, 2}
	p(sliceA)
	sliceB := make([]int, 2, 10)
	p(sliceB) //[[0 0]]
	printSlice(sliceB)
	sliceB = append(sliceB, 3, 4, 5)
	printSlice(sliceB)
	sliceC := make([]int, len(sliceB), 2*cap(sliceB))
	copy(sliceC, sliceB)
	//
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	capital, ok := countryCapitalMap["American"]
	if ok {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}
	delete(countryCapitalMap, "123")
	//
	var dog Animal
	dog = new(Dog)
	dog.eat()
	//
	//_, err := Sqrt(-1)
	//if err != nil {
	//	return
	//}
	// chan
	s := []int{7, 2, 8, -9, 4, 0}
	mychan := make(chan int)
	go sumTotal(s[:len(s)/2], mychan)
	go sumTotal(s[len(s)/2:], mychan)
	x, y := <-mychan, <-mychan
	fmt.Println(x + y)
	fmt.Println(111)
	//
	fmt.Println(`aaa"123"`)
	fmt.Println('1')
}

func p(a ...interface{}) {
	fmt.Println(a)
}
func printSlice(x []int) {
	// len(x) cap(x)
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("math: square root of negative number")
	}
	// 实现
	return 0, nil
}

func sumTotal(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

type Animal interface {
	eat()
	sale()
}
type Dog struct {
	name string
	Name string
}

func (dog Dog) sale() {
	panic("implement me")
}

func (dog Dog) eat() {
	p("eat")
}

type AClz struct {
	key int
	val int
}

func (e AClz) f1() {
	c := e.key
	p(c)
}

func (e AClz) f2() {
}
