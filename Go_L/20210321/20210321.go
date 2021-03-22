package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	T1 = iota
	T2
	T3
)

type User struct {
	Name string
	Age  int
}

type MyInt int

func testReflect() {
	u := User{"xxx", 19,}
	var i MyInt = 12
	t := reflect.TypeOf(u)
	fmt.Println(t.Kind() == reflect.Struct)
	t = reflect.TypeOf(i)
	fmt.Println(t.Kind() == reflect.Int) // 输出true
}

func testAppend() {
	s0 := []int{0}
	s1 := append(s0, 2)
	s2 := append(s0, s1...)
	fmt.Println(s2)
}

func testCast() {
	a := 1
	//类型转换语法：Type(expression)
	var b = float32(a)
	fmt.Println(b)
	// 类型断言 expression.(Type)  expression 必须是接口类型
	var c interface{}
	c = a
	e := c.(int)
	fmt.Println(e)
}

func main() {
	testCast()
	var a chan int
	fmt.Println(a == nil)
	//testS()
	//fmt.Println(StringSliceEqualBCE([]string{"a1", "b"}, []string{"a", "b"}))
	//fmt.Println(T2)
	//testSet()
	////
	//testF()
	//testChan()
	//testBuffChan()
	//
	//testRPrint()
	//testReflect()
}

type Stu struct {
	Name string `json:"stu_name"`
	ID   string `json:"stu_id"`
	Age  int    `json:"-"`
}

func StringSliceEqualBCE(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

//两个协成交替打印1-100的奇数偶数
var POOL int = 100
var wg sync.WaitGroup

func g1(p chan int) {
	for i := 1; i <= POOL; i++ {
		//同步的交替实现
		p <- i
		if i%2 == 1 {
			fmt.Println("g1 ", i)
		}
	}
	wg.Done()
}
func g2(p chan int) {
	for i := 1; i <= POOL; i++ {
		<-p
		if i%2 == 0 {
			fmt.Println("g2 ", i)
		}
	}
	wg.Done()
}

func testRPrint() {
	wg.Add(2)
	ch := make(chan int)
	go g1(ch)
	go g2(ch)
	wg.Wait()
}

func testChan() {
	st := time.Now()
	ch := make(chan bool)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true // 无缓冲，发送方阻塞直到接收方接收到数据。
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds())
	time.Sleep(time.Second * 5)
}

func testBuffChan() {
	st := time.Now()
	ch := make(chan bool, 2)
	go func() {
		time.Sleep(time.Second * 2)
		<-ch
	}()
	ch <- true
	ch <- true                                                // 缓冲区为 2，发送方不阻塞，继续往下执行
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 0.0 s
	ch <- true                                                // 缓冲区使用完，发送方阻塞，2s 后接收方接收到数据，释放一个插槽，继续往下执行
	fmt.Printf("cost %.1f s\n", time.Now().Sub(st).Seconds()) // cost 2.0 s
	time.Sleep(time.Second * 5)
}

func testJson() {
	buf, _ := json.Marshal(Stu{"Tom", "t001", 18})
	fmt.Printf("%s\n", buf)
}

func testS() {
	var str strings.Builder
	for i := 0; i < 1000; i++ {
		str.WriteString("a")
	}
	fmt.Printf(str.String())
}

type Set map[string]struct{}

func testSet() {
	set := make(Set)
	keys := []string{"A", "A", "b", "A"}
	for _, v := range keys {
		set[v] = struct{}{}
	}
	fmt.Println(len(set))
	if _, ok := set["A"]; ok {
		fmt.Printf("Key A exist")
	}
}

type T string

func (t *T) hello() {
	fmt.Printf("hello")
}

func (t T) hello2() {
	fmt.Printf("hello2")
}
func testF() {
	fmt.Print("\n")
	var t1 T = "123"
	t1.hello()
	t2 := &t1
	t2.hello2()
}
