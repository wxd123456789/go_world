package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	pase_student()
	testWG()
	testInterface()
}


type People interface {
	Speak(string) string
}
type Stduent struct{}

func (stu *Stduent) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "You are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func testInterface() {
	var p People
	p = &Stduent{}
	think := "bitch"
	fmt.Println(p.Speak(think))
}

/*go执行的随机性和闭包
解答：谁也不知道执行后打印的顺序是什么样的，所以只能说是随机数字。
但是A:均为输出10，B:从0~9输出(顺序不定)。第一个go func中i是外部for的一个变量，地址不变化。遍历完成后，最终i=10。故go func执行时，i的值始终是10。
第二个go func中i是函数参数，与外部for中的i完全是两个变量。尾部(i)将发生值拷贝，go func内部指向值拷贝地址。*/
func testWG() {
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(20)
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println("A: ", i)
			wg.Done()
		}()
	}
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("B: ", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

type student struct {
	Name string
	Age  int
}

//foreach
func pase_student() {
	m := make(map[string]*student)
	stus := []student{
		{
			Name: "zhou",
			Age:  24,
		}, {
			Name: "li",
			Age:  23,
		}, {
			Name: "wang",
			Age:  22,
		},
	}
	//for _, stu := range stus {
	//	m[stu.Name] = &stu
	//}
	//这样的写法初学者经常会遇到的，很危险！与Java的foreach一样，
	//都是使用副本的方式。所以m[stu.Name]=&stu实际上一致指向同一个指针，
	//最终该指针的值为遍历的最后一个struct的值拷贝。

	for i := 0; i < len(stus); i++ {
		m[stus[i].Name] = &stus[i]
	}
	for _, v := range m {
		fmt.Println(v.Name)
	}
}
