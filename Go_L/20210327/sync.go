package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var a sync.Mutex

func testRuntime() {
	fmt.Println(runtime.NumCPU())
}
func main() {
	//testPool()
	//standard()
	// usePool()
	testRuntime()
}

type C1 struct {
	B1 [10000000]int
}

func usePool() {
	pool := sync.Pool{New:
	func() interface{} {
		return new(C1)
	}}
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		c := pool.Get().(*C1)
		fmt.Println(&c)
		c.B1[0] = 1
		pool.Put(c) //需要加上
	}
	fmt.Println("Used time : ", time.Since(startTime))
}

func standard() {
	startTime := time.Now()
	for i := 0; i < 10000; i++ {
		var c C1
		c.B1[0] = 1
	}
	fmt.Println("Used time : ", time.Since(startTime))
}

type Student struct {
	name string
}

func (s *Student) String() string {
	return s.name
}

func testPool() {
	studentPool := sync.Pool{
		New: func() interface{} {
			return &Student{"abc"}
		}}

	for i := 0; i < 100000; i++ {
		stud := studentPool.Get().(*Student)
		fmt.Printf("%p %v\n", stud, stud)
	}
}

func testNew() {
	a := new(int)
	b := 10
	a = &b
	*a = 11
	fmt.Println(b)
}
